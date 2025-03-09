package main

import (
	"fmt"
	"os"

	"github.com/mstgnz/cli-task-manager/commands"
)

func main() {
	app, err := commands.NewApp()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing application: %v\n", err)
		os.Exit(1)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
