package preflight

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/progress"
)

func fakePreflightModel() PreflightModel {
	return PreflightModel{
		checks:                []SystemCheck{fakeSystemCheck()},
		spinner:               fakeSpinner(),
		progress:              fakeProgress(),
		activeIndex:           0,
		activeCheckpointIndex: 0,
		done:                  false,
	}
}

func fakeProgress() progress.Model {
	return progress.New()
}

func TestGetActive(t *testing.T) {
	p := fakePreflightModel()

	ans := p.getActive()

	want := &p.checks[0]
	if ans != want {
		t.Errorf("got %+v, want %+v", ans, want)
	}
}

func TestGetActiveCheckpoint(t *testing.T) {
	p := fakePreflightModel()

	ans := p.getActiveCheckpoint()

	want := p.checks[0].Checkpoints[0]
	if ans != want {
		t.Errorf("got %+v, want %+v", ans, want)
	}
}

func TestRunCheckpointFailed(t *testing.T) {
	p := fakePreflightModel()

	ansFnc := p.runCheckpoint()

	ans := ansFnc()
	want := systemCheckMsg{check: false}
	if ans != want {
		t.Errorf("got %+v, want %+v", ans, want)
	}
}

func TestRunCheckpointSuccess(t *testing.T) {
	p := fakePreflightModel()
	p.checks[0].Checkpoints[0].Command = "echo"

	ansFnc := p.runCheckpoint()

	ans := ansFnc()
	want := systemCheckMsg{check: true}
	if ans != want {
		t.Errorf("got %+v, want %+v", ans, want)
	}
}

func TestUpdateInternalStateFinished(t *testing.T) {
	p := fakePreflightModel()

	ans, _ := p.UpdateInternalState(systemCheckMsg{check: true})

	if ans.done != true || ans.activeIndex != 1 {
		t.Errorf("got %+v", ans)
	}
}

func TestRenderConclusionFail(t *testing.T) {
	p := fakePreflightModel()

	ans := p.RenderConclusion()

	want := "No go, no go! Check above for more details. 🛬"
	if !strings.Contains(ans, want) {
		t.Errorf("got %s, want %s", ans, want)
	}
}

func TestRenderConclusionWarning(t *testing.T) {
	p := fakePreflightModel()
	p.checks[0].Optional = true

	ans := p.RenderConclusion()

	want := "You're good to go, but check above, some checks were unsuccessful 🎫"
	if !strings.Contains(ans, want) {
		t.Errorf("got %s, want %s", ans, want)
	}
}

func TestRenderConclusionSuccess(t *testing.T) {
	p := fakePreflightModel()
	p.checks[0].Check = true

	ans := p.RenderConclusion()

	want := "Done! You're good to go 🛫"
	if !strings.Contains(ans, want) {
		t.Errorf("got %s, want %s", ans, want)
	}
}
