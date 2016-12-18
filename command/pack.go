package command

import (
	"fmt"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/log"
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
		Commands: []cli.Command{
			cli.Command{
				Path: []string{"{{--configPath}}"},
				Action: func(flags cli.ProcessedFlags) error {
					log.Info("pack")
					return nil
				},
			},
		},
		Action: func(flags cli.ProcessedFlags) error {
			fmt.Println("pack default")
			fmt.Println(flags.Known)
			fmt.Println(flags.Unknown)
			return nil
		},
	}
}
