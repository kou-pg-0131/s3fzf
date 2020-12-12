package infrastructures

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3APIFactory .
type S3APIFactory struct {
	profile string
}

// NewS3APIFactory .
func NewS3APIFactory(profile string) *S3APIFactory {
	return &S3APIFactory{profile: profile}
}

// Create .
func (f *S3APIFactory) Create(region string) s3iface.S3API {
	return s3.New(session.New(), aws.NewConfig().WithRegion(region).WithCredentials(credentials.NewSharedCredentials("", f.profile)))
}
