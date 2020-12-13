package infrastructures

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3Client .
type S3Client struct {
	s3api s3iface.S3API
}

// NewS3Client .
func NewS3Client(s3api s3iface.S3API) *S3Client {
	return &S3Client{s3api: s3api}
}

// SetAPI .
func (c *S3Client) SetAPI(s3api s3iface.S3API) {
	c.s3api = s3api
}

// ListBuckets .
func (c *S3Client) ListBuckets() (*s3.ListBucketsOutput, error) {
	resp, err := c.s3api.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBucketLocation .
func (c *S3Client) GetBucketLocation(bucket string) (*s3.GetBucketLocationOutput, error) {
	resp, err := c.s3api.GetBucketLocation(&s3.GetBucketLocationInput{
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
	resp, err := c.s3api.ListObjectsV2(&s3.ListObjectsV2Input{
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
	resp, err := c.s3api.GetObject(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteObject .
func (c *S3Client) DeleteObject(bucket, key string) (*s3.DeleteObjectOutput, error) {
	resp, err := c.s3api.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
