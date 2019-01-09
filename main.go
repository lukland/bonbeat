package main

import (
	"os"

	"github.com/lucaslandry/bonbeat/cmd"

	_ "github.com/lucaslandry/bonbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
