package main

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

func main() {
	s3c := infrastructures.NewS3Client(
		s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1")),
	)
	fzf := infrastructures.NewFZF()

	bs := []*s3.Bucket{}

	go func() {
		resp, err := s3c.ListBuckets()
		if err != nil {
			panic(err)
		}

		bs = resp
	}()

	i, err := fzf.Find(&bs, func(i int) string {
		return *bs[i].Name
	}, func(i, w, h int) string {
		if i == -1 {
			return ""
		}
		return fmt.Sprintf("%s\n\nCreationDate: %s", *bs[i].Name, *bs[i].CreationDate)
	})
	if err != nil {
		panic(err)
	}

	os := []*s3.Object{}

	go func() {
		reg, err := s3c.GetRegion(*bs[i].Name)
		if err != nil {
			panic(err)
		}

		s3c.SetAPI(s3.New(session.New(), aws.NewConfig().WithRegion(reg)))

		tkn := (*string)(nil)
		for {
			resp, ntkn, err := s3c.ListObjects(*bs[i].Name, tkn)
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

	j, err := fzf.Find(&os, func(j int) string {
		return *os[j].Key
	}, func(j, w, h int) string {
		if j == -1 {
			return ""
		}

		return fmt.Sprintf("%s\n\nSize: %s\nLastModified: %s", *os[j].Key, humanize.Bytes(uint64(*os[j].Size)), *os[j].LastModified)
	})
	if err != nil {
		panic(err)
	}

	f, err := s3c.GetObject(*bs[i].Name, *os[j].Key)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)

	fmt.Print(buf.String())
}
