package main

import (
	"log"
)

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("new config failed: %v", err)
	}
	r := NewMouseRunner(config)
	r.Run()
}
