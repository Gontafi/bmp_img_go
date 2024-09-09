package main

import (
	"bitmap/internal/app"
	"os"
)

func main() {
	err := app.ParseArgsAndRunCommands(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
