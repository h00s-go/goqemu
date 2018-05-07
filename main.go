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
	c, err := config.Load(os.Getenv("HOME") + "/.goqemu/config.json")
	if err != nil {
		log.Fatal(err)
	}

	l, err := logger.New(c.Log.Filename)
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
		switch os.Args[1] {
		case "start":
			guest, err := g.GetGuest(os.Args[2])
			if err == nil {
				output, err := guest.Start()
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Print(output)
			} else {
				fmt.Println(err.Error())
			}
		case "reset":
			guest, err := g.GetGuest(os.Args[2])
			if err == nil {
				guest.Reset()
			} else {
				fmt.Println(err.Error())
			}
		}
	} else {
		fmt.Println("No commands specified. Exiting.")
	}
}
