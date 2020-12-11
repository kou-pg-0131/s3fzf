package cmd

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
 * New()
 */

func Test_New_ReturnCommand(t *testing.T) {
	t.Skip()
}

/*
 * Command.Do()
 */

func TestCommand_Do_ReturnNilWhenSucceeded(t *testing.T) {
	w := new(bytes.Buffer)
	f := ioutil.NopCloser(strings.NewReader("FILE"))

	ms3c := new(mockS3Controller)
	ms3c.On("FindBucket").Return(&s3.Bucket{Name: aws.String("BUCKET")}, nil)
	ms3c.On("FindObject", "BUCKET").Return(&s3.Object{Key: aws.String("KEY")}, nil)
	ms3c.On("GetObject", "BUCKET", "KEY").Return(f, nil)

	cmd := &Command{fileWriter: w, s3Controller: ms3c}

	err := cmd.Do()

	assert.Nil(t, err)
	assert.Equal(t, "FILE", w.String())
	ms3c.AssertNumberOfCalls(t, "FindBucket", 1)
	ms3c.AssertNumberOfCalls(t, "FindObject", 1)
	ms3c.AssertNumberOfCalls(t, "GetObject", 1)
}

func TestCommand_Do_ReturnErrorWhenFindBucketFailed(t *testing.T) {
	ms3c := new(mockS3Controller)
	ms3c.On("FindBucket").Return((*s3.Bucket)(nil), errors.New("SOMETHING_WRONG"))

	cmd := &Command{s3Controller: ms3c}

	err := cmd.Do()

	assert.EqualError(t, err, "SOMETHING_WRONG")
	ms3c.AssertNumberOfCalls(t, "FindBucket", 1)
}

func TestCommand_Do_ReturnErrorWhenFindObjectFailed(t *testing.T) {
	ms3c := new(mockS3Controller)
	ms3c.On("FindBucket").Return(&s3.Bucket{Name: aws.String("BUCKET")}, nil)
	ms3c.On("FindObject", "BUCKET").Return((*s3.Object)(nil), errors.New("SOMETHING_WRONG"))

	cmd := &Command{s3Controller: ms3c}

	err := cmd.Do()

	assert.EqualError(t, err, "SOMETHING_WRONG")
	ms3c.AssertNumberOfCalls(t, "FindBucket", 1)
	ms3c.AssertNumberOfCalls(t, "FindObject", 1)
}

func TestCommand_Do_ReturnErrorWhenGetObjectFailed(t *testing.T) {
	ms3c := new(mockS3Controller)
	ms3c.On("FindBucket").Return(&s3.Bucket{Name: aws.String("BUCKET")}, nil)
	ms3c.On("FindObject", "BUCKET").Return(&s3.Object{Key: aws.String("KEY")}, nil)
	ms3c.On("GetObject", "BUCKET", "KEY").Return((*mockIOReadCloser)(nil), errors.New("SOMETHING_WRONG"))

	cmd := &Command{s3Controller: ms3c}

	err := cmd.Do()

	assert.EqualError(t, err, "SOMETHING_WRONG")
	ms3c.AssertNumberOfCalls(t, "FindBucket", 1)
	ms3c.AssertNumberOfCalls(t, "FindObject", 1)
	ms3c.AssertNumberOfCalls(t, "GetObject", 1)
}

func TestCommand_Do_ReturnErrorWhenCopyFailed(t *testing.T) {
	w := new(bytes.Buffer)
	mf := new(mockIOReadCloser)
	mf.On("Read", mock.Anything).Return(0, errors.New("SOMETHING_WRONG"))
	mf.On("Close").Return(nil)

	ms3c := new(mockS3Controller)
	ms3c.On("FindBucket").Return(&s3.Bucket{Name: aws.String("BUCKET")}, nil)
	ms3c.On("FindObject", "BUCKET").Return(&s3.Object{Key: aws.String("KEY")}, nil)
	ms3c.On("GetObject", "BUCKET", "KEY").Return(mf, nil)

	cmd := &Command{fileWriter: w, s3Controller: ms3c}

	err := cmd.Do()

	assert.EqualError(t, err, "SOMETHING_WRONG")
	ms3c.AssertNumberOfCalls(t, "FindBucket", 1)
	ms3c.AssertNumberOfCalls(t, "FindObject", 1)
	mf.AssertNumberOfCalls(t, "Read", 1)
	mf.AssertNumberOfCalls(t, "Close", 1)
}
