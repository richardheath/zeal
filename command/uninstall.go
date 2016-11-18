package command

import (
	"github.com/richardheath/cli"
)

func InitUninstallCommand() cli.Command {
	return cli.Command{
		Path: []string{"uninstall"},
		FlagTypes: []cli.FlagType{
			cli.FlagType{
				Key:        "package",
				Shorthand:  "p",
				Default:    "",
				Group:      "options",
				Validators: []cli.FlagValidator{},
			},
			cli.FlagType{
				Key:        "force",
				Shorthand:  "f",
				Default:    "false",
				Group:      "options",
				Validators: []cli.FlagValidator{},
			},
			cli.FlagType{
				Key:        "all",
				Shorthand:  "a",
				Default:    "false",
				Group:      "options",
				Validators: []cli.FlagValidator{},
			},
		},
		Commands: []cli.Command{
			cli.Command{
				Path: []string{"{{--package}}"},
				Action: func(app cli.App, knownFlags map[string]string, unknownFlags map[string]string) error {
					return nil
				},
			},
		},
		Action: func(app cli.App, knownFlags map[string]string, unknownFlags map[string]string) error {
			// Error out when package is not provided.
			return nil
		},
	}
}
