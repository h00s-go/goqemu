package main

import (
	"fmt"
	"log"
	"os"

	"github.com/h00s/goqemu/config"
	"github.com/h00s/goqemu/logger"
	"github.com/h00s/goqemu/qemu"
)

func main() {
	c, err := config.Load(os.Getenv("HOME") + "/.goqemu/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	l, err := logger.New(c.Log)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	g, err := qemu.Load(os.Getenv("HOME") + "/.goqemu/guests.json")
	if err != nil {
		fmt.Println(err)
		l.Fatal(err.Error())
	}

	if len(os.Args) > 1 {
		var output string
		switch os.Args[1] {
		case "start":
			output, err = g.Start(os.Args[2])
		case "reset":
			output, err = g.Reset(os.Args[2])
		case "shutdown":
			output, err = g.Shutdown(os.Args[2])
		case "poweroff":
			output, err = g.PowerOff(os.Args[2])
		}
		if err != nil {
			l.Error(err.Error())
			fmt.Println(err)
		}
		if output != "" {
			fmt.Println(output)
		}
	} else {
		fmt.Println("No commands specified. Exiting.")
	}
}
