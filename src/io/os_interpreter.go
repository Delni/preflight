package io

import (
	"fmt"
)

type OSInterpreter struct {
	Interpreter                string
	InterpreterArgs            string
	InterpreterInteractiveArgs string
	Command                    string
	CommandArgs                string
}

func GetInterpreterCommand(os string) (OSInterpreter, error) {
	switch os {
	case "windows":
		return OSInterpreter{
			Interpreter:                "powershell.exe",
			InterpreterArgs:            "",
			InterpreterInteractiveArgs: "",
			Command:                    "command",
			CommandArgs:                "",
		}, nil
	case "darwin":
		return OSInterpreter{
			Interpreter:                "bash",
			InterpreterArgs:            "-c",
			InterpreterInteractiveArgs: "-ic",
			Command:                    "command",
			CommandArgs:                "-v",
		}, nil
	case "linux":
		return OSInterpreter{
			Interpreter:                "bash",
			InterpreterArgs:            "-c",
			InterpreterInteractiveArgs: "-ic",
			Command:                    "command",
			CommandArgs:                "-v",
		}, nil
	default:
		return OSInterpreter{}, fmt.Errorf("OS %s is not currently supported", os)
	}
}
