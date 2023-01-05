package programs

import (
	"preflight/src/systemcheck"
	"testing"
)

var testPresets = map[string]systemcheck.SystemCheck{
	"c": {Name: "A"},
	"b": {Name: "B"},
	"a": {Name: "C"},
}

func TestUsePresetsFilterOutUnknownPresets(t *testing.T) {
	// Act
	systemcheck := UsePresets([]string{"a", "d"}, testPresets)
	// Assert
	if len(systemcheck) != 1 {
		t.Fatalf("expected 1 system check, got %v", systemcheck)
	}
}

func TestUsePresetsReturnKnownPresets(t *testing.T) {
	// Act
	systemcheck := UsePresets([]string{"a"}, testPresets)
	// Assert
	if len(systemcheck) == 0 {
		t.Fatalf("No preset found")
	}
	if systemcheck[0].Name != testPresets["a"].Name {
		t.Fatalf("got %v, expected %v", systemcheck[0], testPresets["a"])
	}
}

func TestAvailablePresetsAreAlphabeticallyOrdered(t *testing.T) {

	given := AvailablePresets(testPresets)
	expect := "Available presets are: a, b, c"
	if given != expect {
		t.Fatalf("got %s, expected %s", given, expect)
	}

}
