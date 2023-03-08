package main

import (
	"flag"
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
	var exitDuration time.Duration
	flag.DurationVar(&freshDuration, "fresh", time.Minute, "mouse fresh time")
	flag.DurationVar(&exitDuration, "duration", 2*time.Hour, "time duration of random-mouse exit")
	flag.Parse()

	now := time.Now()
	exitTime := now.Add(exitDuration)
	log.Printf("now is %s\n", now.Format(TimeFormat))
	log.Printf("%s will exit when time is %s, it's fresh time is %ds\n",
		ProgramName, exitTime.Format(TimeFormat), freshDuration/time.Second)

	exitTimer := time.NewTimer(exitDuration)
	ticker := time.NewTicker(freshDuration)
	for {
		select {
		case <-ticker.C:
			currentPosition := GetMousePos()
			nextPosition := GetNextPosition(currentPosition)
			MoveMouse(nextPosition)
		case <-exitTimer.C:
			exitTimer.Stop()
			ticker.Stop()
			log.Printf("now is %s, time is up, %s exit...\n", time.Now().Format(TimeFormat), ProgramName)
			time.Sleep(5 * time.Second)
			LockScreen()
			os.Exit(0)
		}
	}
}

func LockScreen() {
	robotgo.KeyTap("q", "lcmd", "lctrl")
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
