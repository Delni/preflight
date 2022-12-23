package preflight

import (
	"io"
	"io/ioutil"
	"net/http"
	domain "preflight/src/domain"

	"gopkg.in/yaml.v3"
)

func ReadFile(path string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	return buf, err
}

func ReadHttpFile(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func ReadChecklist(checklist []byte) ([]domain.SystemCheck, error) {
	data := []domain.SystemCheck{}
	err := yaml.Unmarshal(checklist, &data)
	if err != nil {
		return []domain.SystemCheck{}, err
	}

	return data, nil
}
