package controllers

import (
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dustin/go-humanize"
	"github.com/kou-pg-0131/s3fzf/src/interfaces/gateways"
)

// IS3Controller .
type IS3Controller interface {
	FindBucket() (*s3.Bucket, error)
	FindObject(bucket string) (*s3.Object, error)
	GetObject(bucket, key string) (io.ReadCloser, error)
	DeleteObject(bucket, key string) error
}

// S3Controller .
type S3Controller struct {
	s3apiFactory gateways.IS3APIFactory
	s3Client     gateways.IS3Client
	fzf          gateways.IFZF
}

// S3ControllerConfig .
type S3ControllerConfig struct {
	S3APIFactory gateways.IS3APIFactory
	S3Client     gateways.IS3Client
	FZF          gateways.IFZF
}

// NewS3Controller .
func NewS3Controller(cnf *S3ControllerConfig) *S3Controller {
	return &S3Controller{
		s3apiFactory: cnf.S3APIFactory,
		s3Client:     cnf.S3Client,
		fzf:          cnf.FZF,
	}
}

// FindBucket .
func (c *S3Controller) FindBucket() (*s3.Bucket, error) {
	bs := []*s3.Bucket{}

	chidx := make(chan int, 1)
	chfderr := make(chan error, 1)

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
			return
		}

		chidx <- i
	}()

	chlserr := make(chan error, 1)

	go func() {
		resp, err := c.s3Client.ListBuckets()
		if err != nil {
			chlserr <- err
			return
		}

		bs = resp.Buckets
	}()

	select {
	case err := <-chlserr:
		c.fzf.Sync()
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

	chfderr := make(chan error, 1)
	chidx := make(chan int, 1)

	go func() {
		i, err := c.fzf.Find(&os, func(i int) string {
			return *os[i].Key
		}, func(i, w, h int) string {
			if i == -1 {
				return ""
			}

			return fmt.Sprintf("%s\n\nBucket: %s\nKey: %s\nSize: %s\nLastModified: %s", filepath.Base(*os[i].Key), bucket, *os[i].Key, humanize.Bytes(uint64(*os[i].Size)), *os[i].LastModified)
		})
		if err != nil {
			chfderr <- err
			return
		}
		chidx <- i
	}()

	chlserr := make(chan error, 1)

	go func() {
		resp, err := c.s3Client.GetBucketLocation(bucket)
		if err != nil {
			chlserr <- err
			return
		}

		c.s3Client.SetAPI(c.s3apiFactory.Create(*resp.LocationConstraint))

		tkn := (*string)(nil)
		for {
			resp, err := c.s3Client.ListObjects(bucket, tkn)
			if err != nil {
				chlserr <- err
				return
			}

			os = append(os, resp.Contents...)
			if len(os) == 0 {
				chlserr <- fmt.Errorf("`s3://%s` is empty", bucket)
				return
			}

			tkn = resp.NextContinuationToken
			if tkn == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case err := <-chlserr:
		c.fzf.Sync()
		c.fzf.Close()
		return nil, err
	case err := <-chfderr:
		return nil, err
	case idx := <-chidx:
		return os[idx], nil
	}
}

// GetObject .
func (c *S3Controller) GetObject(bucket, key string) (io.ReadCloser, error) {
	resp, err := c.s3Client.GetObject(bucket, key)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// DeleteObject .
func (c *S3Controller) DeleteObject(bucket, key string) error {
	_, err := c.s3Client.DeleteObject(bucket, key)
	return err
}
