package mouse

import (
	"log"
	"time"

	"github.com/dinghenc/random-mouse/checker/mouse"
	"github.com/dinghenc/random-mouse/config"
	"github.com/dinghenc/random-mouse/runner"
	"github.com/dinghenc/random-mouse/utils"
)

type Runner struct {
	config *config.Config
}

func NewRunner(config *config.Config) runner.Runner {
	return &Runner{config: config}
}

func (m *Runner) Run() {
	log.Printf("now is %s\n", m.config.Now.Format(config.TimeFormat))
	log.Printf("%s will exit when time is %s, duration is %s, it's fresh time is %s, checker status is %t\n",
		utils.ProgramName, m.config.ExitTime.Format(config.TimeFormat),
		m.config.ExitDuration.String(), m.config.FreshDuration.String(), m.config.Check)

	exitTimer := time.NewTimer(m.config.ExitDuration)
	ticker := time.NewTicker(m.config.FreshDuration)
	moveChecker := mouse.NewMoveChecker()
	for {
		select {
		case <-ticker.C:
			utils.RandomMoveMouse()
		case <-exitTimer.C:
			exitTimer.Stop()
			ticker.Stop()
			m.Exit("time is up")
			return
		case <-moveChecker.Changed():
			if m.config.Check && moveChecker.Check() {
				m.Exit("checker hit")
				return
			}
		}
	}
}

func (m *Runner) Exit(reason string) {
	log.Printf("now is %s, reason: %s, %s exit...\n",
		time.Now().Format(config.TimeFormat), reason, utils.ProgramName)
	time.Sleep(5 * time.Second)
	utils.LockScreen()
}
