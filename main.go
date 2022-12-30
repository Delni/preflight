package main

import (
	"fmt"
	"os"
	"preflight/src/command"
)

func main() {
	if err := command.PreflightCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
