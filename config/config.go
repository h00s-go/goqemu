package config

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration struct have all fields from configuration JSON file
type Configuration struct {
	Log Log `json:"log"`
}

// Log defines logging configuration (log filename)
type Log struct {
	Filename string `json:"filename"`
}

// Load loads configuration from path
func Load(path string) (Configuration, error) {
	var c Configuration
	configJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(configJSON, &c)
	return c, err
}
