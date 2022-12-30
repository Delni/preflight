package preflight

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
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

func TestUpdateWithSystemCheckMsgShouldUpdateInternalState(t *testing.T) {
	p := makeTestModel()

	raw_ans, _ := p.Update(systemCheckMsg{check: true})
	ans := raw_ans.(PreflightModel)
	if ans.done != true || ans.activeIndex != 1 {
		t.Errorf("got %+v", ans)
	}
}

func TestUpdateWithFrameMsgShouldUpdateProgress(t *testing.T) {
	p := makeTestModel()

	raw_ans, _ := p.Update(progress.FrameMsg{})
	ans := raw_ans.(PreflightModel)
	if &ans.progress == &p.progress {
		t.Errorf("got %+v", ans)
	}
}

func TestUpdateWithWindowSizeMsgShouldUpdateProgressWidth(t *testing.T) {
	p := makeTestModel()

	raw_ans, _ := p.Update(tea.WindowSizeMsg{Width: 100})
	ans := raw_ans.(PreflightModel)
	if ans.progress.Width != 100 {
		t.Errorf("got %+v", ans)
	}
}

func TestShouldQuitOnCtrlC(t *testing.T) {
	p := makeTestModel()

	_, cmd := p.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Errorf("got %+v", cmd)
	}
}
