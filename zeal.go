package main

import (
	"os"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/command"
)

func main() {
	var baseCommand = cli.Command{
		Path: "",
		Commands: []cli.Command{
			command.InitPackCommand(),
		},
		FlagGroups: []cli.FlagGroup{
			cli.FlagGroup{
				Prefix:          "--",
				ShorthandPrefix: "-",
				Group:           "settings",
			},
		},
		FlagTypes: nil,
		Action:    nil,
	}

	args := os.Args[1:]
	cli.Run("zeal", "0.1.0", args, baseCommand)
}
