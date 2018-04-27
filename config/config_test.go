package config

import (
	"testing"
)

func TestConfiguration(t *testing.T) {
	_, err := Load("../config.json")
	if err != nil {
		t.Error("Unable to load configuration")
	}
}
