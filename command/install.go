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
				Action: func(flags cli.ProcessedFlags) error {
					return nil
				},
			},
		},
		Action: func(flags cli.ProcessedFlags) error {
			// Error out when package is not provided.
			return nil
		},
	}
}
