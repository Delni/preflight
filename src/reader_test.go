package preflight

import (
	"io/ioutil"
	"os"
	"reflect"
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

	ans := ReadChecklistFile("tmp.yaml")

	if !reflect.DeepEqual(ans, data) {
		t.Errorf("got %+v, want %+v", 1, 2)
	}

	errDel := os.Remove("tmp.yaml")
	if errDel != nil {
		t.Error("Error while deleting YAML file")
	}
}
