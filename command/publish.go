package command

import (
	"github.com/richardheath/cli"
)

func InitPublishCommand() cli.Command {
	return cli.Command{
		Path: []string{"publish"},
		FlagTypes: []cli.FlagType{
			cli.FlagType{
				Key:        "configPath",
				Shorthand:  "c",
				Default:    "zeal.json",
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
				Path: []string{"{{--configPath}}"},
				Action: func(flags cli.ProcessedFlags) error {
					return nil
				},
			},
		},
		Action: func(flags cli.ProcessedFlags) error {
			return nil
		},
	}
}
