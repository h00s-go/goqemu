package config

import "github.com/BurntSushi/toml"

// Configuration struct have all fields from configuration JSON file
type Configuration struct {
	Log string
}

// Load loads configuration from path
func Load(path string) (*Configuration, error) {
	c := new(Configuration)

	if _, err := toml.DecodeFile(path, c); err != nil {
		return c, err
	}

	return c, nil
}
