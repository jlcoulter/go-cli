package main

import (
	"os"

	"github.com/jlcoulter/go-cli-template/cmd"
)

var version = "dev"

func main() {
	cmd.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}