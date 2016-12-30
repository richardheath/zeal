package command

import (
	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/pack"
)

func InitPackCommand() cli.Command {
	return cli.Command{
		Path:  []string{"pack"},
		Usage: "pack {configPath}",
		FlagTypes: []cli.FlagType{
			cli.FlagType{
				Key:        "configPath",
				Shorthand:  "c",
				Default:    "zeal.json",
				Prefix:     "--",
				Validators: []cli.FlagValidator{},
			},
		},
		Action: func(flags cli.ProcessedFlags) error {
			pack.Pack(flags.Known["--configPath"], "c:\\temp", flags.Unknown)
			return nil
		},
		Commands: []cli.Command{
			cli.Command{
				Path: []string{"{{--configPath}}"},
				Action: func(flags cli.ProcessedFlags) error {
					pack.Pack(flags.Known["--configPath"], "c:\\temp", flags.Unknown)
					return nil
				},
			},
		},
	}
}
