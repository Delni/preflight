package preflight

import (
	"log"
)

type OSInterpreter struct {
	Interpreter                string
	InterpreterArgs            string
	InterpreterInteractiveArgs string
	Command                    string
	CommandArgs                string
}

func getInterpreterCommand(os string) OSInterpreter {
	switch os {
	case "windows":
		return OSInterpreter{
			Interpreter:                "powershell.exe",
			InterpreterArgs:            "",
			InterpreterInteractiveArgs: "",
			Command:                    "command",
			CommandArgs:                "",
		}
	case "darwin":
		return OSInterpreter{
			Interpreter:                "bash",
			InterpreterArgs:            "-c",
			InterpreterInteractiveArgs: "-ic",
			Command:                    "command",
			CommandArgs:                "-v",
		}
	case "linux":
		return OSInterpreter{
			Interpreter:                "bash",
			InterpreterArgs:            "-c",
			InterpreterInteractiveArgs: "-ic",
			Command:                    "command",
			CommandArgs:                "-v",
		}
	default:
		log.Fatalf("OS %s is not currently supported", os)
		return OSInterpreter{}
	}
}
