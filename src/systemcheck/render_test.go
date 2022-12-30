package systemcheck

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/spinner"
)

func fakeSystemCheck() SystemCheck {
	return SystemCheck{
		Name:        "SYSTEM_CHECK",
		Description: "DESCRIPTION",
		Optional:    false,
		Checkpoints: []Checkpoint{
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

	ans := system.RenderSystemCheck(true, spinner)

	want := "| SYSTEM_CHECK\n"
	if ans != want {
		t.Errorf("got %s, want %s", ans, want)
	}
}

func TestRenderInactive(t *testing.T) {
	system := fakeSystemCheck()
	spinner := fakeSpinner()

	ans := system.RenderSystemCheck(false, spinner)

	want := "- SYSTEM_CHECK\n"
	if ans != want {
		t.Errorf("got %s, want %s", ans, want)
	}
}

func TestRenderResultUnchecked(t *testing.T) {
	system := fakeSystemCheck()

	ans := system.RenderResult()

	want := []string{"✕", "SYSTEM_CHECK", "DESCRIPTION", "CHECKPOINT", "DOC"}
	assertMultipleContains(t, ans, want)
}

func TestRenderResultUncheckedWarning(t *testing.T) {
	system := fakeSystemCheck()
	system.Optional = true

	ans := system.RenderResult()

	want := []string{"!", "SYSTEM_CHECK", "DESCRIPTION", "CHECKPOINT", "DOC"}
	assertMultipleContains(t, ans, want)
}

func TestRenderResultChecked(t *testing.T) {
	system := fakeSystemCheck()
	system.Check = true

	ans := system.RenderResult()

	want := []string{"✓", "SYSTEM_CHECK"}
	assertMultipleContains(t, ans, want)
}
