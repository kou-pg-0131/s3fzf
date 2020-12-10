package main

import (
	"github.com/kou-pg-0131/s3ls/src/cmd"
)

func main() {
	if err := cmd.New().Do(); err != nil {
		panic(err)
	}
}
