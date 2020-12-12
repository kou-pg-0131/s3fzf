package cmd

import (
	"github.com/kou-pg-0131/s3fzf/src/interfaces/controllers"
)

// ICommand .
type ICommand interface {
	Do(bucket, output string) error
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

// Do ...
func (c *Command) Do(bucket, output string) error {
	if bucket == "" {
		b, err := c.s3Controller.FindBucket()
		if err != nil {
			return err
		}
		bucket = *b.Name
	}

	o, err := c.s3Controller.FindObject(bucket)
	if err != nil {
		return err
	}

	f, err := c.s3Controller.GetObject(bucket, *o.Key)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := c.fileWriter.Write(output, f); err != nil {
		return err
	}

	return nil
}
