package infrastructures

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// IS3Client .
type IS3Client interface {
	SetAPI(s3iface.S3API)
	ListBuckets() ([]*s3.Bucket, error)
	GetRegion(bucket string) (string, error)
}

// S3Client .
type S3Client struct {
	s3API s3iface.S3API
}

// NewS3Client .
func NewS3Client(s3API s3iface.S3API) IS3Client {
	return &S3Client{s3API: s3API}
}

// SetAPI .
func (c *S3Client) SetAPI(s3API s3iface.S3API) {
	c.s3API = s3API
}

// ListBuckets .
func (c *S3Client) ListBuckets() ([]*s3.Bucket, error) {
	resp, err := c.s3API.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return resp.Buckets, nil
}

// GetRegion .
func (c *S3Client) GetRegion(bucket string) (string, error) {
	resp, err := c.s3API.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: &bucket,
	})
	if err != nil {
		return "", err
	}

	if resp.LocationConstraint == nil {
		return "us-east-1", nil
	}

	return *resp.LocationConstraint, nil
}
