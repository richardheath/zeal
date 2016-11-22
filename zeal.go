package main

import (
	"fmt"
	"os"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/command"
)

func main() {
	app := cli.App{
		Name:    "zeal",
		Version: "0.1.0",
		FlagPrefixes: []cli.FlagPrefix{
			cli.FlagPrefix{
				Key:         "--",
				Shorthand:   "-",
				Description: "options",
			},
			cli.FlagPrefix{
				Key:         "#",
				Shorthand:   "",
				Description: "settings",
			},
		},
		FlagTypes: nil,
		Commands: []cli.Command{
			command.InitPackCommand(),
		},
	}

	err := app.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
