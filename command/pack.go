package command

import (
	"github.com/richardheath/cli"
)

func InitPackCommand() cli.Command {
    return cli.Command{
		Path:     "pack",
		Commands: []cli.Command{},
		FlagGroups: nil,
		FlagTypes: nil,
		Action:    nil,
	}
}