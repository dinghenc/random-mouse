package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	MoveStep   = 1
	TimeFormat = "2006-01-02 15:04:05"
)

var (
	MoveDirection [4]robotgo.Point = [4]robotgo.Point{
		{X: -MoveStep, Y: 0},
		{X: 0, Y: -MoveStep},
		{X: MoveStep, Y: 0},
		{X: 0, Y: MoveStep},
	}
	ProgramName = filepath.Base(os.Args[0])
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var freshDuration time.Duration
	var inExitDuration time.Duration
	var exitTimeStr string
	flag.DurationVar(&freshDuration, "fresh", time.Minute, "mouse fresh time")
	flag.DurationVar(&inExitDuration, "duration", 2*time.Hour, "time duration of random-mouse exit")
	flag.StringVar(&exitTimeStr, "time", "", "absolute time of random-mouse exit, eg: (2006-01-02 15:04:05)")
	flag.Parse()

	now, exitTime, exitDuration, err := CalcExitTime(inExitDuration, exitTimeStr)
	if err != nil {
		log.Fatalf("calc exit time failed: %v", err)
	}

	log.Printf("now is %s\n", now.Format(TimeFormat))
	log.Printf("%s will exit when time is %s, duration is %0.1fs, it's fresh time is %0.1fs\n",
		ProgramName, exitTime.Format(TimeFormat), exitDuration.Seconds(), freshDuration.Seconds())

	exitTimer := time.NewTimer(exitDuration)
	ticker := time.NewTicker(freshDuration)
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

func CalcExitTime(inExitDuration time.Duration, exitTimeStr string) (
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

func LockScreen() {
	robotgo.KeyTap("q", "lcmd", "lctrl")
}

func RandomMoveMouse() {
	currentPosition := GetMousePos()
	nextPosition := GetNextPosition(currentPosition)
	MoveMouse(nextPosition)
}

func GetMousePos() robotgo.Point {
	x, y := robotgo.GetMousePos()
	return robotgo.Point{X: x, Y: y}
}

func MoveMouse(point robotgo.Point) {
	robotgo.Move(point.X, point.Y)
}

func GetNextPosition(point robotgo.Point) robotgo.Point {
	nextIndex := rand.Intn(len(MoveDirection))
	nextPoint := robotgo.Point{
		X: point.X + MoveDirection[nextIndex].X,
		Y: point.Y + MoveDirection[nextIndex].Y,
	}
	return handleCrossScreen(nextPoint)
}

func handleCrossScreen(point robotgo.Point) robotgo.Point {
	width, height := robotgo.GetScreenSize()
	return robotgo.Point{
		X: adjustPos(width, point.X),
		Y: adjustPos(height, point.Y),
	}
}

func adjustPos(maxValue, pos int) int {
	if pos < 0 {
		return 0
	} else if pos >= maxValue {
		return maxValue - 1
	}
	return pos
}
