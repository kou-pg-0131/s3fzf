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
	app.UsageText = "s3fzf <command> [options]"
	app.HideHelpCommand = true

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "show help.",
	}

	defaultFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "bucket",
			Usage:       "name of the bucket containing the objects.",
			Aliases:     []string{"b"},
			Destination: &bucket,
		},
		&cli.StringFlag{
			Name:        "profile",
			Usage:       "use a specific profile from your credential file.",
			Aliases:     []string{"p"},
			Destination: &profile,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:      "cp",
			Usage:     "Copy S3 object to local.",
			UsageText: "s3fzf cp [options]",
			Flags: append(defaultFlags, []cli.Flag{
				&cli.StringFlag{
					Name:        "output",
					Usage:       "file path of the output destination. if '-' is specified, output to stdout.",
					Aliases:     []string{"o"},
					Destination: &output,
					Required:    true,
				},
			}...),
			Action: func(ctx *cli.Context) error {
				return cmd.NewFactory().Create(profile).Copy(bucket, output)
			},
		},
		{
			Name:      "rm",
			Usage:     "Delete an S3 object.",
			UsageText: "s3fzf rm [options]",
			Flags: append(defaultFlags, []cli.Flag{
				&cli.BoolFlag{
					Name:        "no-confirm",
					Usage:       "skip the confirmation before deleting.",
					Destination: &noconf,
				},
			}...),
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
