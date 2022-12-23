package preflight

import (
	"strings"
	"testing"
)

func assertMultipleContains(t *testing.T, ans string, want []string) {
	for _, s := range want {
		if !strings.Contains(ans, s) {
			t.Errorf("%s should contain %s", ans, s)
		}
	}
}
func TestViewOnGoing(t *testing.T) {
	p := makeTestModel()

	ans := p.View()

	want := []string{"|", "SYSTEM_CHECK", "0%"}
	assertMultipleContains(t, ans, want)
}

func TestViewDone(t *testing.T) {
	p := makeTestModel()
	p.done = true

	ans := p.View()

	want := []string{"✕", "SYSTEM_CHECK", "DESCRIPTION", "CHECKPOINT", "DOC", "No go, no go! Check above for more details. 🛬"}
	assertMultipleContains(t, ans, want)
}
