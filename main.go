package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(*bs[i])
}
