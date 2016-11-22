package command

import (
	"github.com/richardheath/cli"
)

func InitInstallCommand() cli.Command {
	return cli.Command{
		Path: []string{"install"},
		FlagTypes: []cli.FlagType{
			cli.FlagType{
				Key:        "package",
				Shorthand:  "p",
				Default:    "",
				Prefix:     "--",
				Validators: []cli.FlagValidator{},
			},
			cli.FlagType{
				Key:        "repo",
				Shorthand:  "r",
				Default:    "",
				Prefix:     "--",
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
