package controllers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kou-pg-0131/s3fzf/src/interfaces/gateways"
)

// IS3Controller .
type IS3Controller interface {
	FindBucket() (*s3.Bucket, error)
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
