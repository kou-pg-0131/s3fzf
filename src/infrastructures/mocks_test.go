package infrastructures

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"github.com/stretchr/testify/mock"
)

/*
 * S3API
 */

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

/*
 * fzf
 */

type mockFZF struct {
	mock.Mock
}

func (m *mockFZF) Find(list interface{}, itemFunc func(int) string, opts ...fzf.Option) (int, error) {
	args := m.Called(list, itemFunc, opts)
	return args.Int(0), args.Error(1)
}

/*
 * termbox
 */

type mockTermbox struct {
	mock.Mock
}

func (m *mockTermbox) Close() {
	m.Called()
}

func (m *mockTermbox) Sync() error {
	args := m.Called()
	return args.Error(0)
}

/*
 * os
 */

type mockOS struct {
	mock.Mock
}

func (m *mockOS) Create(path string) (*os.File, error) {
	args := m.Called(path)
	return args.Get(0).(*os.File), args.Error(1)
}

/*
 * io
 */

type mockIO struct {
	mock.Mock
}

func (m *mockIO) Copy(dist io.Writer, src io.Reader) (int64, error) {
	args := m.Called(dist, src)
	return args.Get(0).(int64), args.Error(1)
}
