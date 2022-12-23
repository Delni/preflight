package preflight

import (
	"strings"
	"testing"
)

func TestRenderConclusionFail(t *testing.T) {
	p := makeTestModel()

	ans := p.RenderConclusion()

	want := "No go, no go! Check above for more details. 🛬"
	if !strings.Contains(ans, want) {
		t.Errorf("got %s, want %s", ans, want)
	}
}

func TestRenderConclusionWarning(t *testing.T) {
	p := makeTestModel()
	p.checks[0].Optional = true

	ans := p.RenderConclusion()

	want := "You're good to go, but check above, some checks were unsuccessful 🎫"
	if !strings.Contains(ans, want) {
		t.Errorf("got %s, want %s", ans, want)
	}
}

func TestRenderConclusionSuccess(t *testing.T) {
	p := makeTestModel()
	p.checks[0].Check = true

	ans := p.RenderConclusion()

	want := "Done! You're good to go 🛫"
	if !strings.Contains(ans, want) {
		t.Errorf("got %s, want %s", ans, want)
	}
}
