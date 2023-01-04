package programs

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestSentenceIsFetchingOnRunningFetch(t *testing.T) {
	testModel := loadOverHttpModel{
		url: "test_url",
	}

	given := testModel.getSentence("P", false)
	expect := "P Fetching file from test_url"

	if !strings.Contains(given, expect) {
		t.Errorf("got %s, expected %s", given, expect)
	}
}

func TestSentenceIsFetchedOnEndedFetch(t *testing.T) {
	testModel := loadOverHttpModel{
		url: "test_url",
	}

	given := testModel.getSentence("P", true)
	expect := "P Fetched file from test_url"

	if !strings.Contains(given, expect) {
		t.Errorf("got %s, expected %s", given, expect)
	}
}

func TestInitiateProgramShouldStoreUrlAndBody(t *testing.T) {
	model := initialModel("test_url")

	if model.url != "test_url" {
		t.Errorf("got %v, expected %v", "test_url", model.url)
	}

	if model.quitting != false {
		t.Errorf("got %v, expected %v", false, model.quitting)
	}

	if model.body != nil {
		t.Errorf("got %v, expected %v", nil, model.body)
	}
}

func TestShouldQuitOnCtrlC(t *testing.T) {
	model := initialModel("test_url")

	given, _ := model.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	quitting := given.(loadOverHttpModel).quitting
	if quitting != true {
		t.Errorf("got %v, expected %v", quitting, true)
	}
}

func TestShouldQuitOnResponse(t *testing.T) {
	model := initialModel("test_url")

	given, _ := model.Update(responseMsg{})
	quitting := given.(loadOverHttpModel).quitting
	if quitting != true {
		t.Errorf("got %v, expected %v", quitting, true)
	}
}

func TestShouldQuitOnError(t *testing.T) {
	model := initialModel("test_url")

	given, _ := model.Update(errMsg{})
	quitting := given.(loadOverHttpModel).quitting
	if quitting != true {
		t.Errorf("got %v, expected %v", quitting, true)
	}
}

func TestShouldQuitOnQuitting(t *testing.T) {
	model := initialModel("test_url")
	model.quitting = true

	_, cmd := model.Update(nil)
	if cmd == nil {
		t.Errorf("got %v, expected %v", cmd, nil)
	}
}

func TestShouldSimplyTickOtherwise(t *testing.T) {
	model := initialModel("test_url")

	_, cmd := model.Update(nil)
	if cmd != nil {
		t.Errorf("got %v, expected tea.Update", cmd)
	}
}

func TestViewShouldPrintSentence(t *testing.T) {
	model := initialModel("test_url")

	given := model.View()
	expect := "Fetching file from test_url"

	if !strings.Contains(given, expect) {
		t.Errorf("got %s, expected %s", given, expect)
	}
}

func TestViewShouldBeEmptyWhenBodyIsLoaded(t *testing.T) {
	model := initialModel("test_url")
	model.body = []byte{1}

	given := model.View()
	expect := ""

	if given != expect {
		t.Errorf("got %s, expected %s", given, expect)
	}
}
