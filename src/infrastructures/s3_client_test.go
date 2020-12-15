package infrastructures

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

/*
 * NewS3Client()
 */

func Test_NewS3Client_ReturnS3Client(t *testing.T) {
	mapi := new(mockS3API)
	s3c := NewS3Client(mapi)

	assert.Equal(t, mapi, s3c.s3api)
}

/*
 * S3Client.ListBuckets()
 */

func TestS3Client_ListBuckets_ReturnOutputWhenSucceeded(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("ListBuckets", &s3.ListBucketsInput{}).Return(&s3.ListBucketsOutput{
		Buckets: []*s3.Bucket{{Name: aws.String("BUCKET")}},
	}, nil)

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.ListBuckets()

	assert.Equal(t, &s3.ListBucketsOutput{
		Buckets: []*s3.Bucket{{Name: aws.String("BUCKET")}},
	}, out)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "ListBuckets", 1)
}

func TestS3Client_ListBuckets_ReturnErrorWhenListBucketsFailed(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("ListBuckets", &s3.ListBucketsInput{}).Return((*s3.ListBucketsOutput)(nil), errors.New("SOMETHING_WRONG"))

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.ListBuckets()

	assert.Nil(t, out)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "ListBuckets", 1)
}

/*
 * S3Client.GetBucketLocation()
 */

func TestS3Client_GetBucketLocation_ReturnOutputWhenSucceded(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("GetBucketLocation", &s3.GetBucketLocationInput{Bucket: aws.String("BUCKET")}).Return(&s3.GetBucketLocationOutput{LocationConstraint: aws.String("REGION")}, nil)

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.GetBucketLocation("BUCKET")

	assert.Equal(t, &s3.GetBucketLocationOutput{LocationConstraint: aws.String("REGION")}, out)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "GetBucketLocation", 1)
}

func TestS3Client_GetBucketLocation_ReturnOutputIncludingUsEast1WhenLocationConstraintIsNil(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("GetBucketLocation", &s3.GetBucketLocationInput{Bucket: aws.String("BUCKET")}).Return(&s3.GetBucketLocationOutput{LocationConstraint: nil}, nil)

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.GetBucketLocation("BUCKET")

	assert.Equal(t, &s3.GetBucketLocationOutput{LocationConstraint: aws.String("us-east-1")}, out)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "GetBucketLocation", 1)
}

func TestS3Client_GetBucketLocation_ReturnErrorWhenGetBucketLocationFailed(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("GetBucketLocation", &s3.GetBucketLocationInput{Bucket: aws.String("BUCKET")}).Return((*s3.GetBucketLocationOutput)(nil), errors.New("SOMETHING_WRONG"))

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.GetBucketLocation("BUCKET")

	assert.Nil(t, out)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "GetBucketLocation", 1)
}

/*
 * S3Client.ListObjects()
 */

func TestS3Client_ListObjects_ReturnOutputWhenSucceeded(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("ListObjectsV2", &s3.ListObjectsV2Input{
		Bucket: aws.String("BUCKET"), ContinuationToken: aws.String("TOKEN"),
	}).Return(&s3.ListObjectsV2Output{Contents: []*s3.Object{&s3.Object{Key: aws.String("KEY")}}}, nil)

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.ListObjects("BUCKET", aws.String("TOKEN"))

	assert.Equal(t, &s3.ListObjectsV2Output{Contents: []*s3.Object{&s3.Object{Key: aws.String("KEY")}}}, out)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "ListObjectsV2", 1)
}

func TestS3Client_ListObjects_ReturnErrorWhenListObjectsV2Failed(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("ListObjectsV2", &s3.ListObjectsV2Input{
		Bucket: aws.String("BUCKET"), ContinuationToken: aws.String("TOKEN"),
	}).Return((*s3.ListObjectsV2Output)(nil), errors.New("SOMETHING_WRONG"))

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.ListObjects("BUCKET", aws.String("TOKEN"))

	assert.Nil(t, out)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "ListObjectsV2", 1)
}

/*
 * S3Client.GetObject()
 */

func TestS3Client_GetObject_ReturnOutputWhenSucceeded(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("GetObject", &s3.GetObjectInput{Bucket: aws.String("BUCKET"), Key: aws.String("KEY")}).Return(&s3.GetObjectOutput{ETag: aws.String("ETAG")}, nil)

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.GetObject("BUCKET", "KEY")

	assert.Equal(t, &s3.GetObjectOutput{ETag: aws.String("ETAG")}, out)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "GetObject", 1)
}

func TestS3Client_GetObject_ReturnErrorWhenGetObjectFailed(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("GetObject", &s3.GetObjectInput{Bucket: aws.String("BUCKET"), Key: aws.String("KEY")}).Return((*s3.GetObjectOutput)(nil), errors.New("SOMETHING_WRONG"))

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.GetObject("BUCKET", "KEY")

	assert.Nil(t, out)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "GetObject", 1)
}

/*
 * S3Client.DeleteObject()
 */

func TestS3Client_DeleteObject_ReturnOutputWhenSucceeded(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("DeleteObject", &s3.DeleteObjectInput{Bucket: aws.String("BUCKET"), Key: aws.String("KEY")}).Return(&s3.DeleteObjectOutput{DeleteMarker: aws.Bool(true)}, nil)

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.DeleteObject("BUCKET", "KEY")

	assert.Equal(t, &s3.DeleteObjectOutput{DeleteMarker: aws.Bool(true)}, out)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "DeleteObject", 1)
}

func TestS3Client_DeleteObject_ReturnErrorWhenDeleteObjectFailed(t *testing.T) {
	mapi := new(mockS3API)
	mapi.On("DeleteObject", &s3.DeleteObjectInput{Bucket: aws.String("BUCKET"), Key: aws.String("KEY")}).Return((*s3.DeleteObjectOutput)(nil), errors.New("SOMETHING_WRONG"))

	s3c := &S3Client{s3api: mapi}

	out, err := s3c.DeleteObject("BUCKET", "KEY")

	assert.Nil(t, out)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "DeleteObject", 1)
}
