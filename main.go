package main

import (
	"fmt"
	"os"

	"bitmap/internal/app"
)

func main() {
	err := app.ParseArgsAndRunCommands(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
