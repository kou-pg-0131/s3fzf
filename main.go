package main

import (
	"fmt"
	"os"

	"github.com/kou-pg-0131/s3fzf/src/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	var (
		profile string
		bucket  string
		output  string
	)

	app := cli.NewApp()
	app.Name = "s3fzf"
	app.Usage = "usage"          // TODO
	app.UsageText = "usage text" // TODO
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "b",
			Usage:       "bucket usage", // TODO
			Aliases:     []string{"bucket"},
			Destination: &bucket,
		},
		&cli.StringFlag{
			Name:        "p",
			Usage:       "profile usage", // TODO
			Aliases:     []string{"profile"},
			Destination: &profile,
		},
		&cli.StringFlag{
			Name:        "o",
			Usage:       "output usage", // TODO
			Aliases:     []string{"output"},
			Destination: &output,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		return cmd.New(profile).Do(bucket, output)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
