package qemu

import (
	"bytes"
	"errors"
)

// Guest represent one guest data
type Guest struct {
	Password string `json:"password"`
	Monitor  struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"monitor"`
	Params  map[string]interface{} `json:"params"`
	Command string
}

// ParseParams generates qemu command line from Params map
func (g *Guest) ParseParams() error {
	var c bytes.Buffer

	for param, value := range g.Params {
		switch value.(type) {
		case string:
			c.WriteString(" -" + param + " " + value.(string))
		case []interface{}:
			for _, v := range value.([]interface{}) {
				c.WriteString(" -" + param + " " + v.(string))
			}
		default:
			return errors.New("Unable to parse: " + param)
		}
	}

	g.Command = c.String()
	return nil
}
