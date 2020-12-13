package gateways

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
	DeleteObject(bucket, key string) (*s3.DeleteObjectOutput, error)
}
