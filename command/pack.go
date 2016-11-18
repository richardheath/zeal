package command

import (
	"github.com/richardheath/cli"
)

func InitPackCommand() cli.Command {
	return cli.Command{
		Path: []string{"pack"},
		FlagTypes: []cli.FlagType{
			cli.FlagType{
				Key:        "configPath",
				Shorthand:  "c",
				Default:    "zeal.json",
				Group:      "options",
				Validators: []cli.FlagValidator{},
			},
		},
		Commands: []cli.Command{
			cli.Command{
				Path: []string{"{{--configPath}}"},
				Action: func(app cli.App, knownFlags map[string]string, unknownFlags map[string]string) error {
					return nil
				},
			},
		},
		Action: func(app cli.App, knownFlags map[string]string, unknownFlags map[string]string) error {
			return nil
		},
	}
}
