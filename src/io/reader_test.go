package io

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"preflight/src/systemcheck"

	"gopkg.in/yaml.v3"
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

func fakeYamlSystemCheckBytes() []byte {
	data := []systemcheck.SystemCheck{testSystemCheck}
	dataBytes, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}
	return dataBytes
}

func TestReadFile(t *testing.T) {
	data := fakeYamlSystemCheckBytes()

	errWrite := ioutil.WriteFile("tmp.yaml", data, 0644)
	if errWrite != nil {
		t.Error("Error while writing YAML file")
	}

	fileContent, _ := ReadFile("tmp.yaml")

	if string(fileContent) != string(data) {
		t.Errorf("got %s, want %s", fileContent, data)
	}

	errDel := os.Remove("tmp.yaml")
	if errDel != nil {
		t.Error("Error while deleting YAML file")
	}
}

func TestReadChecklist(t *testing.T) {
	systemCheck, err := ReadChecklist(fakeYamlSystemCheckBytes())
	if err != nil {
		t.Errorf("got error %s when reading check list: ", err.Error())
	}

	want := []systemcheck.SystemCheck{testSystemCheck}
	if !reflect.DeepEqual(systemCheck, want) {
		t.Errorf("got %+v, want %+v", systemCheck, want)
	}

}

func TestReadChecklistMalformatedFile(t *testing.T) {
	_, err := ReadChecklist([]byte("fake_yml"))

	want := "cannot unmarshal !!str `fake_yml` into []systemcheck.SystemCheck"
	if !strings.Contains(err.Error(), want) {
		t.Errorf("got %+v, want %+v", err, want)
	}
}
