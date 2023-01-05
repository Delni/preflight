package command

import (
	"strings"
	"testing"
)

func TestCommandFlags(t *testing.T) {
	cmd := PreflightCommand()

	for _, flag := range []string{"remote", "checklists"} {
		given := cmd.Flags().Lookup(flag)
		if given == nil {
			prettyFatal(t, flag, given)
		}
	}

}

func TestLongDescription(t *testing.T) {
	desc := makeLongDescription()

	if !strings.Contains(desc, "🛫") {
		prettyFatal(t, "🛫", desc)
	}
}

func prettyFatal(t *testing.T, expected interface{}, given interface{}) {
	t.Fatalf("\nExpected:\n\t%v\n\nGot\n\t%v", expected, given)
}
