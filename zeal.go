package main

import (
	"fmt"
	"os"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/command"
	"github.com/richardheath/zeal/log"
)

func main() {
	err := log.Initialize()
	if err != nil {
		fmt.Println("Failed to initialize logger:")
		fmt.Println(err)
		os.Exit(1)
	}

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

	err = app.Run(os.Args[1:])

	if err != nil {
		fmt.Println("Command failed:")
		fmt.Println(err)
		os.Exit(1)
	}
}
