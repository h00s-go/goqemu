package main

import (
	"log"

	"github.com/h00s/goqemu/config"
)

func main() {
	_, err := config.Load("config.json")
	if err != nil {
		log.Fatal(err)
	}
}
