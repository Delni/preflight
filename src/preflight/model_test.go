package preflight

import (
	"preflight/src/systemcheck"
	"reflect"
	"testing"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
)

var testSystemCheck = systemcheck.SystemCheck{
	Name:        "SYSTEM_CHECK",
	Description: "DESCRIPTION",
	Optional:    false,
	Checkpoints: []systemcheck.Checkpoint{
		{Name: "CHECKPOINT", Command: "CMD", Documentation: "DOC", UseInteractive: true},
	},
	Check: false,
}

func makeTestModel() PreflightModel {
	return PreflightModel{
		checks:                []systemcheck.SystemCheck{testSystemCheck},
		spinner:               spinner.New(),
		progress:              progress.New(),
		activeIndex:           0,
		activeCheckpointIndex: 0,
		done:                  false,
	}
}

func TestGetActive(t *testing.T) {
	p := makeTestModel()

	ans := p.getActive()

	want := &p.checks[0]
	if ans != want {
		t.Errorf("got %+v, want %+v", ans, want)
	}
}

func TestGetActiveCheckpoint(t *testing.T) {
	p := makeTestModel()

	ans := p.getActiveCheckpoint()

	want := p.checks[0].Checkpoints[0]
	if ans != want {
		t.Errorf("got %+v, want %+v", ans, want)
	}
}

func TestRunCheckpointFailed(t *testing.T) {
	p := makeTestModel()

	ansFnc := p.runCheckpoint()

	ans := ansFnc()
	want := systemCheckMsg{check: false}
	if ans != want {
		t.Errorf("got %+v, want %+v", ans, want)
	}
}

func TestRunCheckpointSuccess(t *testing.T) {
	p := makeTestModel()
	p.checks[0].Checkpoints[0].Command = "echo"

	ansFnc := p.runCheckpoint()

	ans := ansFnc()
	want := systemCheckMsg{check: true}
	if ans != want {
		t.Errorf("got %+v, want %+v", ans, want)
	}
}

func TestInitPreflightModel(t *testing.T) {
	s := testSystemCheck

	ans := InitPreflightModel([]systemcheck.SystemCheck{s})
	check := ans.checks[0]

	if !reflect.DeepEqual(check, s) {
		t.Errorf("got %+v, want %+v", check, s)
	}
}
