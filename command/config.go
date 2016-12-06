package command

import (
	"github.com/richardheath/cli"
)

// InitConfigCommand Initialize config command.
func InitConfigCommand() cli.Command {
	return cli.Command{
		Path: []string{"config"},
		FlagTypes: []cli.FlagType{
			cli.FlagType{
				Key:        "--key",
				Shorthand:  "k",
				Default:    "",
				Prefix:     "--",
				Validators: []cli.FlagValidator{},
			},
		},
		Commands: []cli.Command{
			cli.Command{
				Path: []string{"set {{--key}} {{--value}}"},
				FlagTypes: []cli.FlagType{
					cli.FlagType{
						Key:        "--value",
						Shorthand:  "v",
						Default:    "",
						Prefix:     "--",
						Validators: []cli.FlagValidator{},
					},
				},
				Action: func(flags cli.ProcessedFlags) error {
					return nil
				},
			},
			cli.Command{
				Path: []string{"get {{--key}}"},
				Action: func(flags cli.ProcessedFlags) error {
					return nil
				},
			},
			cli.Command{
				Path: []string{"delete {{--key}}"},
				Action: func(flags cli.ProcessedFlags) error {
					return nil
				},
			},
			cli.Command{
				Path: []string{"list"},
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
