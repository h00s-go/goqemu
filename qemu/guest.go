package qemu

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

// Guest represent one guest data
type Guest struct {
	Qemu     string `json:"qemu" binding:"required"`
	Password string `json:"password" binding:"required"`
	QMP      struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"monitor" binding:"required"`
	Params  map[string]interface{} `json:"params"`
	Command string
}

// ParseParams generates qemu command line from Params map
func (g *Guest) ParseParams() error {
	var c bytes.Buffer
	c.WriteString(g.Qemu)

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

	c.WriteString(" -qmp tcp:" + g.QMP.Address + ":" + g.QMP.Port + ",server,nowait")
	c.WriteString(" -daemonize")
	g.Command = c.String()
	return nil
}

// Start starts guest using Command
func (g *Guest) Start() {
	out, err := exec.Command("bash", "-c", g.Command).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}
