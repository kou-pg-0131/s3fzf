package cmd

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dustin/go-humanize"
	"github.com/kou-pg-0131/s3fzf/src/infrastructures"
)

// Command .
type Command struct {
	s3Client infrastructures.IS3Client
	fzf      infrastructures.IFZF
}

// New .
func New() *Command {
	s3API := s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

	return &Command{
		s3Client: infrastructures.NewS3Client(s3API),
		fzf:      infrastructures.NewFZF(),
	}
}

// Do ...
func (c *Command) Do() error {
	b, err := c.findBucket()
	if err != nil {
		return err
	}

	o, err := c.findObject(*b.Name)
	if err != nil {
		return err
	}

	resp, err := c.s3Client.GetObject(*b.Name, *o.Key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	fmt.Print(buf.String())
	return nil
}

func (c *Command) findBucket() (*s3.Bucket, error) {
	bs := []*s3.Bucket{}

	chidx := make(chan int, 1)
	chlserr := make(chan error, 1)
	chfderr := make(chan error, 1)

	go func() {
		resp, err := c.s3Client.ListBuckets()
		if err != nil {
			chlserr <- err
		}

		bs = resp.Buckets
	}()

	go func() {
		i, err := c.fzf.Find(&bs, func(i int) string {
			return *bs[i].Name
		}, func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("%s\n\nCreationDate: %s", *bs[i].Name, *bs[i].CreationDate)
		})
		if err != nil {
			chfderr <- err
		}

		chidx <- i
	}()

	select {
	case err := <-chlserr:
		c.fzf.Close()
		return nil, err
	case err := <-chfderr:
		return nil, err
	case idx := <-chidx:
		return bs[idx], nil
	}
}

func (c *Command) findObject(bucket string) (*s3.Object, error) {
	os := []*s3.Object{}

	chlserr := make(chan error, 1)
	chfderr := make(chan error, 1)
	chidx := make(chan int, 1)

	go func() {
		resp, err := c.s3Client.GetBucketLocation(bucket)
		if err != nil {
			chlserr <- err
		}

		c.s3Client.SetAPI(s3.New(session.New(), aws.NewConfig().WithRegion(*resp.LocationConstraint)))

		tkn := (*string)(nil)
		for {
			resp, err := c.s3Client.ListObjects(bucket, tkn)
			if err != nil {
				chlserr <- err
			}

			os = append(os, resp.Contents...)
			if len(os) == 0 {
				chlserr <- fmt.Errorf("`s3://%s` is empty", bucket)
			}

			tkn = resp.NextContinuationToken
			if tkn == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		i, err := c.fzf.Find(&os, func(i int) string {
			return *os[i].Key
		}, func(i, w, h int) string {
			if i == -1 {
				return ""
			}

			return fmt.Sprintf("%s\n\nSize: %s\nLastModified: %s", *os[i].Key, humanize.Bytes(uint64(*os[i].Size)), *os[i].LastModified)
		})
		if err != nil {
			chfderr <- err
		}
		chidx <- i
	}()

	select {
	case err := <-chlserr:
		c.fzf.Close()
		return nil, err
	case err := <-chfderr:
		return nil, err
	case idx := <-chidx:
		return os[idx], nil
	}
}
