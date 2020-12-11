package controllers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kou-pg-0131/s3fzf/src/infrastructures"
)

// S3ControllerFactory .
type S3ControllerFactory struct{}

// NewS3ControllerFactory .
func NewS3ControllerFactory() *S3ControllerFactory {
	return new(S3ControllerFactory)
}

// Create .
func (f *S3ControllerFactory) Create() IS3Controller {
	return NewS3Controller(&S3ControllerConfig{
		S3Client: infrastructures.NewS3Client(s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))),
		FZF:      infrastructures.NewFZF(),
	})
}
