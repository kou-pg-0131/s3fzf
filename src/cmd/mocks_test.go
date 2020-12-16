package cmd

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/mock"
)

/*
 * ScannerAPI
 */

type mockScannerAPI struct {
	mock.Mock
}

func (m *mockScannerAPI) Scan() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *mockScannerAPI) Text() string {
	args := m.Called()
	return args.String(0)
}

func (m *mockScannerAPI) Err() error {
	args := m.Called()
	return args.Error(0)
}

/*
 * Scanner
 */

type mockScanner struct {
	mock.Mock
}

func (m *mockScanner) Scan() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

/*
 * Confirmer
 */

type mockConfirmer struct {
	mock.Mock
}

func (m *mockConfirmer) Confirm(msg string) (bool, error) {
	args := m.Called(msg)
	return args.Bool(0), args.Error(1)
}

/*
 * FileWriter
 */

type mockFileWriter struct {
	mock.Mock
}

func (m *mockFileWriter) Write(out string, r io.Reader) error {
	args := m.Called(out, r)
	return args.Error(0)
}

/*
 * io.ReadCloser
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
 * S3Controller
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

func (m *mockS3Controller) DeleteObject(bucket, key string) error {
	args := m.Called(bucket, key)
	return args.Error(0)
}

/*
 * FileSystem
 */

type mockFileSystem struct {
	mock.Mock
}

func (m *mockFileSystem) Copy(dest io.Writer, src io.Reader) error {
	args := m.Called(dest, src)
	return args.Error(0)
}

func (m *mockFileSystem) Create(path string) (io.WriteCloser, error) {
	args := m.Called(path)
	return args.Get(0).(io.WriteCloser), args.Error(1)
}

/*
 * io.WriteCloser
 */

type mockIOWriteCloser struct {
	mock.Mock
}

func (m *mockIOWriteCloser) Write(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *mockIOWriteCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}
