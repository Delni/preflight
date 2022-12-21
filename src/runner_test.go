package preflight

import (
	"reflect"
	"testing"
)

func TestPreflighModel(t *testing.T) {
	s := fakeSystemCheck()

	ans := PreflighModel([]SystemCheck{s})
	check := ans.checks[0]

	if !reflect.DeepEqual(check, s) {
		t.Errorf("got %+v, want %+v", check, s)
	}
}

func TestViewOnGoing(t *testing.T) {
	p := fakePreflightModel()

	ans := p.View()

	want := []string{"|", "SYSTEM_CHECK", "0%"}
	assertMultipleContains(t, ans, want)
}

func TestViewDone(t *testing.T) {
	p := fakePreflightModel()
	p.done = true

	ans := p.View()

	want := []string{"✕", "SYSTEM_CHECK", "DESCRIPTION", "CHECKPOINT", "DOC", "No go, no go! Check above for more details. 🛬"}
	assertMultipleContains(t, ans, want)
}
