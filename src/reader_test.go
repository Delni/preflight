package preflight

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestReadChecklistFile(t *testing.T) {
	data := []SystemCheck{
		{Name: "System", Description: "Descr", Optional: true, Checkpoints: []Checkpoint{
			{Name: "Checkpoint", Command: "CMD", Documentation: "Doc", UseInteractive: true},
		}},
	}
	dataStr, err := yaml.Marshal(&data)
	if err != nil {
		t.Error("Error while serializing YAML")
	}
	errWrite := ioutil.WriteFile("tmp.yaml", dataStr, 0644)
	if errWrite != nil {
		t.Error("Error while writing YAML file")
	}

	ans, _ := ReadChecklistFile("tmp.yaml")

	if !reflect.DeepEqual(ans, data) {
		t.Errorf("got %+v, want %+v", 1, 2)
	}

	errDel := os.Remove("tmp.yaml")
	if errDel != nil {
		t.Error("Error while deleting YAML file")
	}
}

func TestReadChecklistFileMissingFile(t *testing.T) {
	_, err := ReadChecklistFile("missing_file.yaml")

	want := `checklist "missing_file.yaml" not found`
	if err.Error() != want {
		t.Errorf("got %+v, want %+v", err, want)
	}
}

func TestReadChecklistFileMalformatedFile(t *testing.T) {
	errWrite := ioutil.WriteFile("tmp.yaml", []byte("fake_yml"), 0644)
	if errWrite != nil {
		t.Error("Error while writing YAML file")
	}

	_, err := ReadChecklistFile("tmp.yaml")

	want := "cannot parse file: yaml: unmarshal errors"
	if !strings.Contains(err.Error(), want) {
		t.Errorf("got %+v, want %+v", err, want)
	}

	errDel := os.Remove("tmp.yaml")
	if errDel != nil {
		t.Error("Error while deleting YAML file")
	}
}
