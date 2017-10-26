package main

import (
	"os"

	"github.com/jaredbancroft/ugebeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
