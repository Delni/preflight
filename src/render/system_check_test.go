package render

import (
	"strings"
	"testing"

	domain "preflight/src/domain"

	"github.com/charmbracelet/bubbles/spinner"
)

func fakeSystemCheck() domain.SystemCheck {
	return domain.SystemCheck{
		Name:        "SYSTEM_CHECK",
		Description: "DESCRIPTION",
		Optional:    false,
		Checkpoints: []domain.Checkpoint{
			{Name: "CHECKPOINT", Command: "CMD", Documentation: "DOC", UseInteractive: true},
		},
		Check: false,
	}
}

func fakeSpinner() spinner.Model {
	return spinner.New()
}

func assertMultipleContains(t *testing.T, ans string, want []string) {
	for _, s := range want {
		if !strings.Contains(ans, s) {
			t.Errorf("%s should contain %s", ans, s)
		}
	}
}

func TestRenderActive(t *testing.T) {
	system := fakeSystemCheck()
	spinner := fakeSpinner()

	ans := RenderSystemCheck(system, true, spinner)

	want := "| SYSTEM_CHECK\n"
	if ans != want {
		t.Errorf("got %s, want %s", ans, want)
	}
}

func TestRenderInactive(t *testing.T) {
	system := fakeSystemCheck()
	spinner := fakeSpinner()

	ans := RenderSystemCheck(system, false, spinner)

	want := "- SYSTEM_CHECK\n"
	if ans != want {
		t.Errorf("got %s, want %s", ans, want)
	}
}

func TestRenderResultUnchecked(t *testing.T) {
	system := fakeSystemCheck()

	ans := RenderResultFor(system)

	want := []string{"✕", "SYSTEM_CHECK", "DESCRIPTION", "CHECKPOINT", "DOC"}
	assertMultipleContains(t, ans, want)
}

func TestRenderResultUncheckedWarning(t *testing.T) {
	system := fakeSystemCheck()
	system.Optional = true

	ans := RenderResultFor(system)

	want := []string{"!", "SYSTEM_CHECK", "DESCRIPTION", "CHECKPOINT", "DOC"}
	assertMultipleContains(t, ans, want)
}

func TestRenderResultChecked(t *testing.T) {
	system := fakeSystemCheck()
	system.Check = true

	ans := RenderResultFor(system)

	want := []string{"✓", "SYSTEM_CHECK"}
	assertMultipleContains(t, ans, want)
}
