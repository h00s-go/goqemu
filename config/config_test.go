package config

import (
	"testing"
)

func TestConfiguration(t *testing.T) {
	_, err := Load("../config_test.json")
	if err != nil {
		t.Error("Unable to load configuration")
	}
}
