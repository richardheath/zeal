package main

import (
	"github.com/richardheath/cli"
	"github.com/richardheath/zeal/config"
)

func initCLI(app *cli.App) {
	globalConfig := config.GetZealConfig()
	app.Flag("option", "repo", globalConfig.DefaultRepo, nil)
	app.Flag("option", "registry", globalConfig.Registry, nil)

	initBuildCommand(app)
	initTestCommand(app)
	initPackageCommand(app)
	initPublishCommand(app)
	initInstallCommand(app)
	initUninstallCommand(app)
	initRunCommand(app)
	initStopCommand(app)
}

func initBuildCommand(app *cli.App) {
	cmd := app.Command("build", nil)
	cmd.Flag("option", "zealPath", "./zeal", nil)
	cmd.Command("{{option:zealPath}}", nil)
}

func initTestCommand(app *cli.App) {
	cmd := app.Command("test", nil)
	cmd.Flag("option", "zealPath", "./zeal", nil)
	cmd.Command("{{option:zealPath}}", nil)
}

func initPackageCommand(app *cli.App) {
	cmd := app.Command("package", nil)
	cmd.Flag("option", "zealPath", "./zeal", nil)
	cmd.Command("{{option:zealPath}}", nil)
}

func initPublishCommand(app *cli.App) {
	cmd := app.Command("publish", nil)
	cmd.Flag("option", "zealPath", "./zeal", nil)
	cmd.Command("{{option:zealPath}}", nil)
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