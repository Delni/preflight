package preflight

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadChecklistFile(path string) ([]SystemCheck, error) {
	buf, err := ioutil.ReadFile(path)

	if err != nil {
		return []SystemCheck{}, fmt.Errorf("checklist \"%s\" not found", path)
	}

	data := []SystemCheck{}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return []SystemCheck{}, fmt.Errorf("cannot parse file: %v", err)
	}

	return data, nil
}
