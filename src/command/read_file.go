package command

import (
	"fmt"
	"os"
	"preflight/src/io"
	"preflight/src/programs"
	"preflight/src/systemcheck"
)

func ReadFile(filePath string) []systemcheck.SystemCheck {
	var (
		dataBytes []byte
		err       error
	)
	if Remote {
		dataBytes, err = programs.LoadHttpFileFrom(filePath)
	} else {
		dataBytes, err = io.ReadFile(filePath)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	systemChecks, err := io.ReadChecklist(dataBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return systemChecks
}
