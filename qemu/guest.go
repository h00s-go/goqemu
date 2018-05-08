package qemu

import (
	"bytes"
	"errors"
	"os/exec"
)

// Guest represent one guest data
type Guest struct {
	Qemu struct {
		Monitor struct {
			Address string `json:"address"`
			Port    string `json:"port"`
		} `json:"monitor" binding:"required"`
		Binary string `json:"binary" binding:"required"`
	} `json:"qemu" binding:"required"`
	Password string                 `json:"password"`
	Params   map[string]interface{} `json:"params" binding:"required"`
	QMP      *QMP
}

// ParseParams generates qemu command line from Params map
func (g *Guest) ParseParams() (string, error) {
	var c bytes.Buffer
	c.WriteString(g.Qemu.Binary)

	for param, value := range g.Params {
		switch value.(type) {
		case string:
			c.WriteString(" -" + param + " " + value.(string))
		case []interface{}:
			for _, v := range value.([]interface{}) {
				c.WriteString(" -" + param + " " + v.(string))
			}
		default:
			return "", errors.New("Unable to parse: " + param)
		}
	}

	c.WriteString(" -qmp tcp:" + g.Qemu.Monitor.Address + ":" + g.Qemu.Monitor.Port + ",server,nowait")
	c.WriteString(" -daemonize")
	return c.String(), nil
}

// IsRunning checks if Guest is running
func (g *Guest) IsRunning() bool {
	return g.QMP.IsRunning()
}

// Start starts guest using Command
func (g *Guest) Start() (string, error) {
	startCommand, err := g.ParseParams()
	if g.IsRunning() {
		return "", errors.New("Guest is already running")
	}
	if err != nil {
		return "", errors.New("Unable to parse guest params:" + err.Error())
	}
	output, err := exec.Command("bash", "-c", startCommand).CombinedOutput()
	if err != nil {
		return string(output), errors.New("There was an error starting guest")
	}
	return string(output), nil
}

// Reset does guest system reset
func (g *Guest) Reset() (string, error) {
	if !g.IsRunning() {
		return "", errors.New("Guest is not running")
	}
	return g.QMP.SendCommand("system_reset")
}
