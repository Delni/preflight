package preflight

import "testing"

func TestUpdateInternalStateFinished(t *testing.T) {
	p := makeTestModel()

	ans, _ := p.UpdateInternalState(systemCheckMsg{check: true})

	if ans.done != true || ans.activeIndex != 1 {
		t.Errorf("got %+v", ans)
	}
}
