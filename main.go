package main

import (
	"os"

	"bitmap/internal/app"
)

func main() {
	err := app.ParseArgsAndRunCommands(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
