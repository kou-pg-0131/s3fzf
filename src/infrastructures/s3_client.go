package infrastructures

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// IS3Client .
type IS3Client interface {
	ListBuckets() ([]*s3.Bucket, error)
}

// S3Client .
type S3Client struct {
	s3API s3iface.S3API
}

// NewS3Client .
func NewS3Client(s3API s3iface.S3API) IS3Client {
	return &S3Client{s3API: s3API}
}

// ListBuckets .
func (c *S3Client) ListBuckets() ([]*s3.Bucket, error) {
	resp, err := c.s3API.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return resp.Buckets, nil
}
