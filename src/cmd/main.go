package cmd

import (
	"io"

	"github.com/kou-pg-0131/s3fzf/src/interfaces/controllers"
)

// Command .
type Command struct {
	fileWriter   io.Writer
	s3Controller controllers.IS3Controller
}

// New .
func New(profile string, out io.Writer) *Command {
	return &Command{
		fileWriter:   out,
		s3Controller: controllers.NewS3ControllerFactory().Create(profile),
	}
}

// Do ...
func (c *Command) Do(bucket string) error {
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

	if _, err := io.Copy(c.fileWriter, f); err != nil {
		return err
	}

	return nil
}
