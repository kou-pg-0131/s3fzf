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
		noconf  bool
	)

	app := cli.NewApp()
	app.Name = "s3fzf"
	app.Usage = "Fuzzy Finder for AWS S3."
	app.UsageText = "s3fzf [global options]"
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "bucket",
			Usage:       "The name of the bucket containing the objects",
			Aliases:     []string{"b"},
			Destination: &bucket,
		},
		&cli.StringFlag{
			Name:        "profile",
			Usage:       "Use a specific profile from your credential file",
			Aliases:     []string{"p"},
			Destination: &profile,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:      "cp",
			Usage:     "cp usage",      // TODO
			UsageText: "cp usage text", // TODO
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "output",
					Usage:       "File path of the output destination",
					Aliases:     []string{"o"},
					Destination: &output,
					Required:    true,
				},
			},
			Action: func(ctx *cli.Context) error {
				return cmd.NewFactory().Create(profile).Copy(bucket, output)
			},
		},
		{
			Name:      "rm",
			Usage:     "rm usage",      // TODO
			UsageText: "rm usage text", // TODO
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "no-confirm",
					Usage:       "no-confirm usage",
					Destination: &noconf,
				},
			},
			Action: func(ctx *cli.Context) error {
				return cmd.NewFactory().Create(profile).Remove(bucket, noconf)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
