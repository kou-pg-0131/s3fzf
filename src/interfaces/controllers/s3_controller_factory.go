package controllers

import (
	"github.com/kou-pg-0131/s3fzf/src/infrastructures"
)

// S3ControllerFactory .
type S3ControllerFactory struct{}

// NewS3ControllerFactory .
func NewS3ControllerFactory() *S3ControllerFactory {
	return new(S3ControllerFactory)
}

// Create .
func (f *S3ControllerFactory) Create(profile string) IS3Controller {
	apif := infrastructures.NewS3APIFactory(profile)

	return NewS3Controller(&S3ControllerConfig{
		S3APIFactory: apif,
		S3Client:     infrastructures.NewS3Client(apif.Create("us-east-1")),
		FZF:          infrastructures.NewFZF(),
	})
}
