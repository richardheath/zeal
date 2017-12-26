package main

import (
	"fmt"

	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/config"
)

func initCLI(app *cli.App) {
	globalConfig := config.GetZealConfig()
	app.Flag("option", "repo", globalConfig.DefaultRepo, nil)
	app.Flag("option", "registry", globalConfig.Registry, nil)
	app.Flag("option", "path p", "./zeal", nil)

	initInitCommand(app)
	initBuildCommand(app)
	initTestCommand(app)
	initPackCommand(app)
	initPublishCommand(app)
	initInstallCommand(app)
	initUninstallCommand(app)
	initRunCommand(app)
	initStopCommand(app)
}

func initInitCommand(app *cli.App) {
	app.Command("init", nil)
}

func initBuildCommand(app *cli.App) {
	app.Command("build", buildCommand)
}

func initTestCommand(app *cli.App) {
	app.Command("test", nil)
}

func initPackCommand(app *cli.App) {
	app.Command("pack", packCommand)
}

func initPublishCommand(app *cli.App) {
	app.Command("publish", nil)
}

func initInstallCommand(app *cli.App) {
	cmd := app.Command("install {{option:package}}", nil)
	cmd.Flag("option", "package p", "", nil)
	cmd.Flag("option", "group g", "default", nil)
}

func initUninstallCommand(app *cli.App) {
	cmd := app.Command("uninstall {{option:package}}", nil)
	cmd.Flag("option", "package p", "", nil)
	cmd.Flag("option", "group g", "default", nil)
}

func initRunCommand(app *cli.App) {
	cmd := app.Command("run {{option:package}}", nil)
	cmd.Flag("option", "package p", "", nil)
	cmd.Flag("option", "group g", "default", nil)
}

func initStopCommand(app *cli.App) {
	cmd := app.Command("stop {{option:package}}", nil)
	cmd.Flag("option", "package p", "", nil)
	cmd.Flag("option", "group g", "default", nil)
}

func echoHello(ctx *cli.Context) error {
	fmt.Println("build!!!")
	return nil
}