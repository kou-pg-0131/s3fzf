package main

import (
	"fmt"
	"os"

	"github.com/kou-pg-0131/s3fzf/src/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "s3fzf"
	app.Usage = "usage"          // TODO
	app.UsageText = "usage text" // TODO
	app.HideHelpCommand = true

	app.Action = func(ctx *cli.Context) error {
		return cmd.New(os.Stdout).Do()
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
