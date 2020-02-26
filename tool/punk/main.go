package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "punk"
	app.Usage = "punk工具集"
	app.Version = Version
	app.Commands = []*cli.Command{
		{
			Name:            "new",
			Aliases:         []string{"n"},
			Usage:           "创建新项目",
			Action:          runNew,
			SkipFlagParsing: true,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "punk build",
			Action:  buildAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "punk run",
			Action:  runAction,
		},
		{
			Name:            "tool",
			Aliases:         []string{"t"},
			Usage:           "punk tool",
			Action:          toolAction,
			SkipFlagParsing: true,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "punk version",
			Action: func(c *cli.Context) error {
				fmt.Println(getVersion())
				return nil
			},
		},
		{
			Name:   "self-upgrade",
			Usage:  "punk self-upgrade",
			Action: upgradeAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func runNew(ctx *cli.Context) error {
	return installAndRun("genproject", ctx.Args().Slice())
}
