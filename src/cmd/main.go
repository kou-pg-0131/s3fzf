package cmd

import (
	"fmt"
	"io"

	"github.com/kou-pg-0131/s3fzf/src/interfaces/controllers"
)

// Command .
type Command struct {
	fileWriter   io.Writer
	s3Controller controllers.IS3Controller
}

// New .
func New(out io.Writer) *Command {
	return &Command{
		fileWriter:   out,
		s3Controller: controllers.NewS3ControllerFactory().Create(),
	}
}

// Do ...
func (c *Command) Do() error {
	b, err := c.s3Controller.FindBucket()
	if err != nil {
		return err
	}

	o, err := c.s3Controller.FindObject(*b.Name)
	if err != nil {
		return err
	}

	bs, err := c.s3Controller.GetObject(*b.Name, *o.Key)
	if err != nil {
		return err
	}

	fmt.Fprint(c.fileWriter, string(bs))
	return nil
}
