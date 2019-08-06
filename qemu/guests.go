package qemu

import (
	"errors"

	"github.com/BurntSushi/toml"
)

// Guests are all guests defined in json
type Guests map[string]*Guest

// Load loads guests configuration from path
func Load(path string) (Guests, error) {
	var g Guests
	_, err := toml.DecodeFile(path, &g)
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
		return guest.Reset()
	}
	return "", errors.New("Guest does not exist")
}

// Shutdown guest with specified name
func (g Guests) Shutdown(name string) (string, error) {
	if guest, ok := g[name]; ok {
		return guest.Shutdown()
	}
	return "", errors.New("Guest does not exist")
}

// PowerOff guest with specified name
func (g Guests) PowerOff(name string) (string, error) {
	if guest, ok := g[name]; ok {
		return guest.PowerOff()
	}
	return "", errors.New("Guest does not exist")
}
