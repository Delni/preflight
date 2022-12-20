package preflight

import (
	"fmt"
	"testing"
)

func TestGetInterpreterCommand(t *testing.T) {
	var tests = []struct {
		os   string
		want OSInterpreter
	}{
		{"windows", OSInterpreter{
			Interpreter:                "powershell.exe",
			InterpreterArgs:            "",
			InterpreterInteractiveArgs: "",
			Command:                    "command",
			CommandArgs:                "",
		}},
		{"darwin", OSInterpreter{
			Interpreter:                "bash",
			InterpreterArgs:            "-c",
			InterpreterInteractiveArgs: "-ic",
			Command:                    "command",
			CommandArgs:                "-v",
		}},
		{"linux", OSInterpreter{
			Interpreter:                "bash",
			InterpreterArgs:            "-c",
			InterpreterInteractiveArgs: "-ic",
			Command:                    "command",
			CommandArgs:                "-v",
		}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("GetIntrepreterCommand for OS %s", tt.os)
		t.Run(testname, func(t *testing.T) {
			ans := GetInterpreterCommand(tt.os)
			if ans != tt.want {
				t.Errorf("got %+v, want %+v", ans, tt.want)
			}
		})
	}
}
