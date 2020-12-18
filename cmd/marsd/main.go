package main

import (
	"os"

	"github.com/ilgooz/mars/cmd/marsd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
