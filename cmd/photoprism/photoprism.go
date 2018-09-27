package main

import (
	"github.com/photoprism/photoprism/internal/commands"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "PhotoPrism"
	app.Usage = "Digital Photo Archive"
	app.Version = "0.0.0"
	app.Copyright = "Copyright (c) 2018 Michael Mayer <michael@liquidbytes.net> and contributors"
	app.EnableBashCompletion = true
	app.Flags = commands.GlobalFlags

	app.Commands = []cli.Command{
		commands.ConfigCommand,
		commands.StartCommand,
		commands.MigrateCommand,
		commands.ImportCommand,
		commands.IndexCommand,
		commands.ConvertCommand,
		commands.ThumbnailsCommand,
		commands.ExportCommand,
	}

	app.Run(os.Args)
}
