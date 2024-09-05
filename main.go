package main

import (
	"bitmap/internal/app"
	"fmt"
	"os"
)

func main() {
	err := app.ParseArgsAndRunCommands(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
