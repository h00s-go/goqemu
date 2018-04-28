package qemu

import (
	"encoding/json"
	"io/ioutil"
)

// Guests are all guests defined in json
type Guests map[string]*Guest

// Load loads guests configuration from path
func Load(path string) (Guests, error) {
	var g Guests
	guestsJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return g, err
	}
	err = json.Unmarshal(guestsJSON, &g)
	if err != nil {
		return g, err
	}
	for _, guest := range g {
		err = guest.ParseParams()
		if err != nil {
			return g, err
		}
	}
	return g, err
}
