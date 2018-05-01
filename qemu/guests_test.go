package qemu

import (
	"fmt"
	"testing"
)

func TestLoadingAndParsing(t *testing.T) {
	g, err := Load("../guests.json")
	if err != nil {
		t.Error("Error while loading guests", err)
	}
	for _, guest := range g {
		fmt.Println(guest.ParseParams())
	}
}
