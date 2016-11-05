package command

import (
	"github.com/richardheath/cli"
)

func InitPackCommand() cli.Command {
	return cli.Command{
		Path:      []string{"pack"},
		Commands:  []cli.Command{},
		FlagTypes: nil,
		Action:    nil,
	}
}
