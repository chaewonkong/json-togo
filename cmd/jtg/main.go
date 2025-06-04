package main

import (
	"log"

	"github.com/chaewonkong/json-togo/cmd"
)

func main() {
	command := cmd.New()
	if err := command.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
