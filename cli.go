package main

import (
	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/processor"
)

func generateCliCommands() []cli.Command {
	return []cli.Command{
		initCommonCommandsHandler(),
	}
}

func initCommonCommandsHandler() cli.Command {
	return cli.Command{
		Path:  []string{"{{--command}}"},
		Usage: "{command} {package}",
		FlagTypes: []cli.FlagType{
			cli.FlagType{
				Key:        "command",
				Shorthand:  "c",
				Default:    "",
				Prefix:     "--",
				Validators: []cli.FlagValidator{},
			},
			cli.FlagType{
				Key:        "package",
				Shorthand:  "p",
				Default:    "./zeal",
				Prefix:     "--",
				Validators: []cli.FlagValidator{},
			},
		},
		Action: func(flags cli.ProcessedFlags) error {
			cmd := flags.Known["--command"]
			pkg := flags.Known["--package"]
			return processor.ExecuteCommand(cmd, pkg)
		},
		Commands: []cli.Command{
			cli.Command{
				Path: []string{"{{--package}}"},
				Action: func(flags cli.ProcessedFlags) error {
					cmd := flags.Known["--command"]
					pkg := flags.Known["--package"]
					return processor.ExecuteCommand(cmd, pkg)
				},
			},
		},
	}
}