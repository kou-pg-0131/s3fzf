package cmd

import (
	"github.com/kou-pg-0131/s3fzf/src/interfaces/controllers"
)

// Factory .
type Factory struct{}

// NewFactory .
func NewFactory() *Factory {
	return new(Factory)
}

// Create .
func (f *Factory) Create(profile string) ICommand {
	return New(&CommandConfig{
		FileWriter:   NewFileWriter(),
		S3Controller: controllers.NewS3ControllerFactory().Create(profile),
	})
}
