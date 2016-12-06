package main

import (
	"fmt"
	"os"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/command"
)

func main() {
	app := cli.NewApp("zeal", "0.1.0")

	app.FlagPrefixes = []cli.FlagPrefix{
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
	}

	app.Commands = []cli.Command{
		command.InitPackCommand(),
		command.InitConfigCommand(),
	}

	err := app.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
