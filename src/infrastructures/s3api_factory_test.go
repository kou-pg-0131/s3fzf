package infrastructures

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

/*
 * NewS3APIFactory()
 */

func Test_NewS3APIFactory_ReturnS3APIFactory(t *testing.T) {
	f := NewS3APIFactory("PROFILE")

	assert.Equal(t, &S3APIFactory{profile: "PROFILE"}, f)
}

/*
 * S3APIFactory.Create()
 */

func TestS3APIFactory_Create_ReturnS3API(t *testing.T) {
	f := &S3APIFactory{profile: "PROFILE"}
	api := f.Create("REGION")

	assert.IsType(t, &s3.S3{}, api)
	assert.NotNil(t, api)
}
