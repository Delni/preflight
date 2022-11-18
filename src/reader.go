package preflight

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadChecklistFile(path string) []SystemCheck {
	buf, err := ioutil.ReadFile(fmt.Sprintf("./checklists/%s.yaml", path))

	if err != nil {
		fmt.Printf("Checklist \"%s\" not found\n", path)
		os.Exit(1)
	}

	data := []SystemCheck{}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		log.Fatalf("Cannot parse file: %v", err)
	}

	return data
}
