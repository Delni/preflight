package io

import (
	"io"
	"io/ioutil"
	"net/http"

	"preflight/src/systemcheck"

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

func ReadChecklist(checklist []byte) ([]systemcheck.SystemCheck, error) {
	data := []systemcheck.SystemCheck{}
	err := yaml.Unmarshal(checklist, &data)
	if err != nil {
		return []systemcheck.SystemCheck{}, err
	}

	return data, nil
}
