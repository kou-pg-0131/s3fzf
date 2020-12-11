package cmd

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kou-pg-0131/s3fzf/src/infrastructures"
	"github.com/kou-pg-0131/s3fzf/src/interfaces/controllers"
	"github.com/kou-pg-0131/s3fzf/src/interfaces/gateways"
)

// Command .
type Command struct {
	s3Controller controllers.IS3Controller
	s3Client     gateways.IS3Client
	fzf          gateways.IFZF
}

// New .
func New() *Command {
	s3api := s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

	return &Command{
		s3Controller: controllers.NewS3ControllerFactory().Create(),
		s3Client:     infrastructures.NewS3Client(s3api),
		fzf:          infrastructures.NewFZF(),
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

	resp, err := c.s3Client.GetObject(*b.Name, *o.Key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	fmt.Print(buf.String())
	return nil
}
