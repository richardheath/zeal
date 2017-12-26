package main

import (
	"fmt"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/config"
)

func buildCommand(ctx *cli.Context) error {
	configPath, _ := ctx.Flags.GetValue("option", "path")
	cfg, err := config.LoadDir(configPath)
	if err != nil {
		return err
	}

	fmt.Println("build")
	fmt.Println(configPath)
	fmt.Println(cfg.Filter("build", "", ""))
	return nil
}