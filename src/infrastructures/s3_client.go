package infrastructures

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// IS3Client .
type IS3Client interface {
	SetAPI(s3iface.S3API)
	ListBuckets() (*s3.ListBucketsOutput, error)
	GetBucketLocation(bucket string) (*s3.GetBucketLocationOutput, error)
	ListObjects(bucket string, tkn *string) (*s3.ListObjectsV2Output, error)
	GetObject(bucket, key string) (*s3.GetObjectOutput, error)
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
func (c *S3Client) ListBuckets() (*s3.ListBucketsOutput, error) {
	resp, err := c.s3API.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBucketLocation .
func (c *S3Client) GetBucketLocation(bucket string) (*s3.GetBucketLocationOutput, error) {
	resp, err := c.s3API.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: &bucket,
	})
	if err != nil {
		return nil, err
	}

	if resp.LocationConstraint == nil {
		resp.SetLocationConstraint("us-east-1")
	}

	return resp, nil
}

// ListObjects .
func (c *S3Client) ListObjects(bucket string, tkn *string) (*s3.ListObjectsV2Output, error) {
	resp, err := c.s3API.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:            &bucket,
		ContinuationToken: tkn,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetObject .
func (c *S3Client) GetObject(bucket, key string) (*s3.GetObjectOutput, error) {
	resp, err := c.s3API.GetObject(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
