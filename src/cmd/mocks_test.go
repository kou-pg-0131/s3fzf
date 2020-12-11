package cmd

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/mock"
)

/*
 * mockIOReadCloser
 */

type mockIOReadCloser struct {
	mock.Mock
}

func (m *mockIOReadCloser) Read(p []byte) (int, error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *mockIOReadCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

/*
 * mockS3Controller
 */

type mockS3Controller struct {
	mock.Mock
}

func (m *mockS3Controller) FindBucket() (*s3.Bucket, error) {
	args := m.Called()
	return args.Get(0).(*s3.Bucket), args.Error(1)
}

func (m *mockS3Controller) FindObject(bucket string) (*s3.Object, error) {
	args := m.Called(bucket)
	return args.Get(0).(*s3.Object), args.Error(1)
}

func (m *mockS3Controller) GetObject(bucket, key string) (io.ReadCloser, error) {
	args := m.Called(bucket, key)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}
