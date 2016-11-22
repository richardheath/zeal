package command

import (
	"fmt"

	"github.com/richardheath/cli"
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
				Action: func(app cli.App, knownFlags map[string]string, unknownFlags map[string]string) error {
					fmt.Println("pack --configPath")
					fmt.Println(knownFlags)
					fmt.Println(unknownFlags)
					return nil
				},
			},
		},
		Action: func(app cli.App, knownFlags map[string]string, unknownFlags map[string]string) error {
			fmt.Println("pack default")
			fmt.Println(knownFlags)
			fmt.Println(unknownFlags)
			return nil
		},
	}
}
