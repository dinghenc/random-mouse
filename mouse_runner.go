package main

import (
	"log"
	"time"
)

type MouseRunner struct {
	config *Config
}

func NewMouseRunner(config *Config) Runner {
	return &MouseRunner{config: config}
}

func (m *MouseRunner) Run() {
	log.Printf("now is %s\n", m.config.now.Format(TimeFormat))
	log.Printf("%s will exit when time is %s, duration is %s, it's fresh time is %s\n",
		ProgramName, m.config.exitTime.Format(TimeFormat),
		m.config.exitDuration.String(), m.config.freshDuration.String())

	exitTimer := time.NewTimer(m.config.exitDuration)
	ticker := time.NewTicker(m.config.freshDuration)
	for {
		select {
		case <-ticker.C:
			RandomMoveMouse()
		case <-exitTimer.C:
			exitTimer.Stop()
			ticker.Stop()
			log.Printf("now is %s, time is up, %s exit...\n", time.Now().Format(TimeFormat), ProgramName)
			time.Sleep(5 * time.Second)
			LockScreen()
			return
		}
	}
}
