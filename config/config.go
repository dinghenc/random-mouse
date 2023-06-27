package config

import (
	"flag"
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

var (
	freshDuration  time.Duration
	inExitDuration time.Duration
	exitTimeStr    string
	check          bool
)

type Config struct {
	Now           time.Time
	ExitTime      time.Time
	ExitDuration  time.Duration
	FreshDuration time.Duration
	Check         bool
}

func init() {
	flag.DurationVar(&freshDuration, "fresh", time.Minute, "mouse fresh time")
	flag.DurationVar(&inExitDuration, "duration", 2*time.Hour, "time duration of random-mouse exit")
	flag.StringVar(&exitTimeStr, "time", "", "absolute time of random-mouse exit, eg: (2006-01-02 15:04:05)")
	flag.BoolVar(&check, "check", false, "do check mouse move")
	flag.Parse()
}

func NewConfig() (*Config, error) {
	now, exitTime, exitDuration, err := calcExitTime(inExitDuration, exitTimeStr)
	if err != nil {
		return nil, fmt.Errorf("calc exit time failed: %w", err)
	}
	return &Config{
		Now:           now,
		ExitTime:      exitTime,
		ExitDuration:  exitDuration,
		FreshDuration: freshDuration,
		Check:         check,
	}, nil
}

func calcExitTime(inExitDuration time.Duration, exitTimeStr string) (
	now time.Time, exitTime time.Time, exitDuration time.Duration, err error) {
	now = time.Now()
	if exitTimeStr == "" {
		return now, now.Add(inExitDuration), inExitDuration, nil
	}
	exitTime, err = time.ParseInLocation(TimeFormat, exitTimeStr, time.Local)
	if err != nil {
		return now, exitTime, inExitDuration, fmt.Errorf("parse exit time failed: %w", err)
	} else if exitTime.Before(now) {
		return now, exitTime, inExitDuration, fmt.Errorf("exit time is illegal, before now")
	}
	return now, exitTime, exitTime.Sub(now), nil
}
