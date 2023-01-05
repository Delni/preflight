package preflight

import (
	"preflight/src/systemcheck"
	"testing"
)

func TestUpdateInternalStateFinished(t *testing.T) {
	p := makeTestModel()

	ans, _ := p.UpdateInternalState(systemCheckMsg{check: true})

	if ans.done != true || ans.activeIndex != 1 {
		t.Errorf("got %+v", ans)
	}
}

func TestUpdateInternalStateContinue(t *testing.T) {
	p := makeTestModel()
	p.checks = []systemcheck.SystemCheck{
		testSystemCheck,
		testSystemCheck,
	}

	ans, _ := p.UpdateInternalState(systemCheckMsg{check: true})

	if ans.done != false {
		t.Errorf("got %+v", ans.done)
	}

	if ans.activeIndex != 1 {
		t.Errorf("got %+v", ans.activeIndex)
	}

	if ans.progress.Percent() != .5 {
		t.Errorf("got %+v", ans.progress.Percent())
	}
}
