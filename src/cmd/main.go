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

	f, err := c.s3Client.GetObject(*b.Name, *o.Key)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)

	fmt.Print(buf.String())
	return nil
}

func (c *Command) findBucket() (*s3.Bucket, error) {
	bs := []*s3.Bucket{}

	go func() {
		resp, err := c.s3Client.ListBuckets()
		if err != nil {
			panic(err)
		}

		bs = resp
	}()

	i, err := c.fzf.Find(&bs, func(i int) string {
		return *bs[i].Name
	}, func(i, w, h int) string {
		if i == -1 {
			return ""
		}
		return fmt.Sprintf("%s\n\nCreationDate: %s", *bs[i].Name, *bs[i].CreationDate)
	})
	if err != nil {
		return nil, err
	}

	return bs[i], nil
}

func (c *Command) findObject(bucket string) (*s3.Object, error) {
	os := []*s3.Object{}

	go func() {
		reg, err := c.s3Client.GetRegion(bucket)
		if err != nil {
			panic(err)
		}

		c.s3Client.SetAPI(s3.New(session.New(), aws.NewConfig().WithRegion(reg)))

		tkn := (*string)(nil)
		for {
			resp, ntkn, err := c.s3Client.ListObjects(bucket, tkn)
			if err != nil {
				panic(err)
			}

			os = append(os, resp...)

			tkn = ntkn
			if tkn == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()

	i, err := c.fzf.Find(&os, func(i int) string {
		return *os[i].Key
	}, func(i, w, h int) string {
		if i == -1 {
			return ""
		}

		return fmt.Sprintf("%s\n\nSize: %s\nLastModified: %s", *os[i].Key, humanize.Bytes(uint64(*os[i].Size)), *os[i].LastModified)
	})
	if err != nil {
		return nil, err
	}

	return os[i], nil
}
