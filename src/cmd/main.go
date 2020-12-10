package cmd

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dustin/go-humanize"
	"github.com/kou-pg-0131/s3ls/src/infrastructures"
)

// Command .
type Command struct {
	s3Client infrastructures.IS3Client
	fzf      infrastructures.IFZF
}

// New .
func New() *Command {
	s3API := s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

	return &Command{
		s3Client: infrastructures.NewS3Client(s3API),
		fzf:      infrastructures.NewFZF(),
	}
}

// Do ...
func (c *Command) Do() error {
	bs := []*s3.Bucket{}

	go func() {
		resp, err := c.s3Client.ListBuckets()
		if err != nil {
			panic(err)
		}

		bs = resp
	}()

	i, err := c.fzf.Find(&bs, func(i int) string {
		return *bs[i].Name
	}, func(i, w, h int) string {
		if i == -1 {
			return ""
		}
		return fmt.Sprintf("%s\n\nCreationDate: %s", *bs[i].Name, *bs[i].CreationDate)
	})
	if err != nil {
		return err
	}

	os := []*s3.Object{}

	go func() {
		reg, err := c.s3Client.GetRegion(*bs[i].Name)
		if err != nil {
			panic(err)
		}

		c.s3Client.SetAPI(s3.New(session.New(), aws.NewConfig().WithRegion(reg)))

		tkn := (*string)(nil)
		for {
			resp, ntkn, err := c.s3Client.ListObjects(*bs[i].Name, tkn)
			if err != nil {
				panic(err)
			}

			os = append(os, resp...)

			tkn = ntkn
			if tkn == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()

	j, err := c.fzf.Find(&os, func(j int) string {
		return *os[j].Key
	}, func(j, w, h int) string {
		if j == -1 {
			return ""
		}

		return fmt.Sprintf("%s\n\nSize: %s\nLastModified: %s", *os[j].Key, humanize.Bytes(uint64(*os[j].Size)), *os[j].LastModified)
	})
	if err != nil {
		return err
	}

	f, err := c.s3Client.GetObject(*bs[i].Name, *os[j].Key)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)

	fmt.Print(buf.String())
	return nil
}
