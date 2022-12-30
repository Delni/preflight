package io

import (
	"fmt"
	"testing"
)

func TestGetInterpreterCommand(t *testing.T) {
	var tests = []struct {
		os   string
		want OSInterpreter
		err  error
	}{
		{"windows", OSInterpreter{
			Interpreter:                "powershell.exe",
			InterpreterArgs:            "",
			InterpreterInteractiveArgs: "",
			Command:                    "command",
			CommandArgs:                "",
		}, nil},
		{"darwin", OSInterpreter{
			Interpreter:                "bash",
			InterpreterArgs:            "-c",
			InterpreterInteractiveArgs: "-ic",
			Command:                    "command",
			CommandArgs:                "-v",
		}, nil},
		{"linux", OSInterpreter{
			Interpreter:                "bash",
			InterpreterArgs:            "-c",
			InterpreterInteractiveArgs: "-ic",
			Command:                    "command",
			CommandArgs:                "-v",
		}, nil},
		{"other_os", OSInterpreter{}, fmt.Errorf("OS %s is not currently supported", "other_os")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("GetIntrepreterCommand for OS %s", tt.os)
		t.Run(testname, func(t *testing.T) {
			ans, err := GetInterpreterCommand(tt.os)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("got %+v, want %+v", err, tt.err)
				}
			} else if ans != tt.want {
				t.Errorf("got %+v, want %+v", ans, tt.want)
			}
		})
	}
}
