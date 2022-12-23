package preflight

import (
	"io/ioutil"
	"log"
	"os"
	domain "preflight/src/domain"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func fakeYamlSystemCheckBytes() []byte {
	data := []domain.SystemCheck{fakeSystemCheck()}
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

	want := []domain.SystemCheck{fakeSystemCheck()}
	if !reflect.DeepEqual(systemCheck, want) {
		t.Errorf("got %+v, want %+v", systemCheck, want)
	}

}

func TestReadChecklistMalformatedFile(t *testing.T) {
	_, err := ReadChecklist([]byte("fake_yml"))

	want := "cannot unmarshal !!str `fake_yml` into []preflight.SystemCheck"
	if !strings.Contains(err.Error(), want) {
		t.Errorf("got %+v, want %+v", err, want)
	}
}
