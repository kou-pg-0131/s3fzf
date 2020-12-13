package cmd

import (
	"github.com/kou-pg-0131/s3fzf/src/interfaces/controllers"
)

// ICommand .
type ICommand interface {
	Copy(bucket, output string) error
}

// Command .
type Command struct {
	fileWriter   IFileWriter
	s3Controller controllers.IS3Controller
}

// CommandConfig .
type CommandConfig struct {
	FileWriter   IFileWriter
	S3Controller controllers.IS3Controller
}

// New .
func New(cnf *CommandConfig) ICommand {
	return &Command{
		fileWriter:   cnf.FileWriter,
		s3Controller: cnf.S3Controller,
	}
}
