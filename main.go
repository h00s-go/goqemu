package main

import (
	"log"

	"github.com/h00s/goqemu/config"
	"github.com/h00s/goqemu/logger"
	"github.com/h00s/goqemu/qemu"
)

func main() {
	c, err := config.Load("config.json")
	if err != nil {
		log.Fatal(err)
	}

	l, err := logger.New(c.Log.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	_, err = qemu.Load("guests.json")
	if err != nil {
		l.Fatal(err.Error())
	}
}
