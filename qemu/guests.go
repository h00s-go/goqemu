package qemu

import (
	"encoding/json"
	"errors"
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
		guest.QMP = NewQMP(guest.Qemu.Monitor.Address, guest.Qemu.Monitor.Port)
	}
	return g, err
}

// Start guest with specified name
func (g Guests) Start(name string) (string, error) {
	if guest, ok := g[name]; ok {
		return guest.Start()
	}
	return "", errors.New("Guest does not exist")
}

// Reset guest with specified name
func (g Guests) Reset(name string) (string, error) {
	if guest, ok := g[name]; ok {
		guest.Reset()
		return "", nil
	}
	return "", errors.New("Guest does not exist")
}
