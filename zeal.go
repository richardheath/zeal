package main

import (
	"fmt"
	"os"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/config"
	"github.com/richardheath/zeal/log"
)

func main() {
	err := config.Initialize()
	if err != nil {
		exitWithError("Failed to initialize config:", err)
	}

	err = log.Initialize()
	if err != nil {
		exitWithError("Failed to initialize logger:", err)
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
			Description: "variables",
		},
	}

	app.Commands = generateCliCommands()
	err = app.Run(os.Args[1:])
	
	if err != nil {
		exitWithError("Command failed:", err)
	}
}

func exitWithError(message string, err error) {
	fmt.Println(message)
	fmt.Println(err)
	os.Exit(1)
}
