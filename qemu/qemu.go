package qemu

import (
	"encoding/json"
	"io/ioutil"
)

// Guest represent one guest data
type Guest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Monitor  struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"monitor"`
	Params map[string]interface{} `json:"params"`
}

// Load loads guests configuration from path
func Load(path string) ([]Guest, error) {
	var g []Guest
	guestsJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return g, err
	}
	err = json.Unmarshal(guestsJSON, &g)
	return g, err
}
