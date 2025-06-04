package main

import "log"

func main() {
	cmd := NewCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
