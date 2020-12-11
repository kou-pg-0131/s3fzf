package controllers

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dustin/go-humanize"
	"github.com/kou-pg-0131/s3fzf/src/interfaces/gateways"
)

// IS3Controller .
type IS3Controller interface {
	FindBucket() (*s3.Bucket, error)
	FindObject(bucket string) (*s3.Object, error)
	GetObject(bucket, key string) ([]byte, error)
}

// S3Controller .
type S3Controller struct {
	s3Client gateways.IS3Client
	fzf      gateways.IFZF
}

// S3ControllerConfig .
type S3ControllerConfig struct {
	S3Client gateways.IS3Client
	FZF      gateways.IFZF
}

// NewS3Controller .
func NewS3Controller(cnf *S3ControllerConfig) *S3Controller {
	return &S3Controller{
		s3Client: cnf.S3Client,
		fzf:      cnf.FZF,
	}
}

// FindBucket .
func (c *S3Controller) FindBucket() (*s3.Bucket, error) {
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

// FindObject .
func (c *S3Controller) FindObject(bucket string) (*s3.Object, error) {
	os := []*s3.Object{}

	chlserr := make(chan error, 1)

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

	chfderr := make(chan error, 1)
	chidx := make(chan int, 1)

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

// GetObject .
func (c *S3Controller) GetObject(bucket, key string) ([]byte, error) {
	resp, err := c.s3Client.GetObject(bucket, key)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.Bytes(), nil
}
