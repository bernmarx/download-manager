package main

import (
	"os"

	"github.com/urfave/cli/v2"

	cl "github.com/bernmarx/download-manager/internal/app/cli"
	"github.com/bernmarx/download-manager/internal/infrastructure/logger"
)

func main() {
	app := &cli.App{
		Name:   "dm",
		Usage:  "Download files",
		Action: cl.Process,
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Logger().Error(err.Error())
	}
}
