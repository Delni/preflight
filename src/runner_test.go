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
