package main

import (
	"encoding/json"
	"io/ioutil"
)

type DockerConfiguration struct {
	Network string
}

func getDockerConfiguration(path string) (*DockerConfiguration, error) {
	jsonBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dockerConfiguration DockerConfiguration
	err = json.Unmarshal(jsonBytes, &dockerConfiguration)

	return &dockerConfiguration, err
}
