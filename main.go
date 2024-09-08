package main

import (
	"os"

	"bitmap/internal/app"
	"bitmap/pkg"
)

func main() {
	err := app.ParseArgsAndRunCommands(os.Args)
	if err != nil {
		pkg.PrintUsage()
		os.Exit(1)
	}
}
