package config

import (
	"os"
	"testing"
)

func TestConfiguration(t *testing.T) {
	_, err := Load(os.Getenv("HOME") + "/.goqemu/config.toml")
	if err != nil {
		t.Error("Unable to load configuration")
	}
}
