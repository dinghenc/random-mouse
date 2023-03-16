package main

import (
	"log"

	"github.com/dinghenc/random-mouse/config"
	"github.com/dinghenc/random-mouse/runner/mouse"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("new config failed: %v", err)
	}
	r := mouse.NewMouseRunner(cfg)
	r.Run()
}
