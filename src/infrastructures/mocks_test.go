package infrastructures

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/mock"
)

type mockS3API struct {
	mock.Mock
	s3iface.S3API
}

func (m *mockS3API) ListBuckets(i *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	args := m.Called(i)
	return args.Get(0).(*s3.ListBucketsOutput), args.Error(1)
}

func (m *mockS3API) GetBucketLocation(i *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	args := m.Called(i)
	return args.Get(0).(*s3.GetBucketLocationOutput), args.Error(1)
}

func (m *mockS3API) ListObjectsV2(i *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	args := m.Called(i)
	return args.Get(0).(*s3.ListObjectsV2Output), args.Error(1)
}

func (m *mockS3API) GetObject(i *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	args := m.Called(i)
	return args.Get(0).(*s3.GetObjectOutput), args.Error(1)
}
