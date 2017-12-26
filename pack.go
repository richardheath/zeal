package main

import (
	"fmt"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/config"
)

func packCommand(ctx *cli.Context) error {
	configPath, _ := ctx.Flags.GetValue("option", "path")
	fmt.Println("pack")

	// load/init local state on target folder

	// load zeal configs on target folder
	cfg, err := config.LoadDir(configPath)
	if err != nil {
		return err
	}

	packageCfg := cfg.Filter("package", "", "")
	cfg.SetItems(packageCfg)

	// evaluate zeal config
	err = cfg.Evaluate(nil)
	if err != nil {
		return err
	}
	// register local handlers for pack
	// - settings, contents, options

	// get pack specific configs
	/*

		for _, item := range packageCfg {
			fmt.Println(item.Name)
			err = item.Interpolate(nil)
			if err != nil {
				return err
			}

			fmt.Println(item)
		}
	*/
	// run pack configs

	return nil
}
