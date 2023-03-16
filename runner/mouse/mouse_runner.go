package mouse

import (
	"log"
	"time"

	"github.com/dinghenc/random-mouse/config"
	"github.com/dinghenc/random-mouse/runner"
	"github.com/dinghenc/random-mouse/utils"
)

type Runner struct {
	config *config.Config
}

func NewMouseRunner(config *config.Config) runner.Runner {
	return &Runner{config: config}
}

func (m *Runner) Run() {
	log.Printf("now is %s\n", m.config.Now.Format(config.TimeFormat))
	log.Printf("%s will exit when time is %s, duration is %s, it's fresh time is %s\n",
		utils.ProgramName, m.config.ExitTime.Format(config.TimeFormat),
		m.config.ExitDuration.String(), m.config.FreshDuration.String())

	exitTimer := time.NewTimer(m.config.ExitDuration)
	ticker := time.NewTicker(m.config.FreshDuration)
	for {
		select {
		case <-ticker.C:
			utils.RandomMoveMouse()
		case <-exitTimer.C:
			exitTimer.Stop()
			ticker.Stop()
			log.Printf("now is %s, time is up, %s exit...\n", time.Now().Format(config.TimeFormat), utils.ProgramName)
			time.Sleep(5 * time.Second)
			utils.LockScreen()
			return
		}
	}
}
