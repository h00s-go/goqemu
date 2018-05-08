package qemu

import (
	"os"
)

// Start guest with specified name
func Start(g Guests, name string) (string, error) {
	guest, err := g.GetGuest(os.Args[2])
	if err == nil {
		output, err := guest.Start()
		return output, err
	}
	return "", err
}

// Reset guest with specified name
func Reset(g Guests, name string) (string, error) {
	guest, err := g.GetGuest(os.Args[2])
	if err == nil {
		guest.Reset()
	}
	return "", err
}
