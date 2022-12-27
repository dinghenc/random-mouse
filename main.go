package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

const MoveStep = 30

var MoveDirection [4]robotgo.Point = [4]robotgo.Point{
	{X: -MoveStep, Y: -MoveStep},
	{X: -MoveStep, Y: MoveStep},
	{X: MoveStep, Y: -MoveStep},
	{X: MoveStep, Y: MoveStep},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var freshTime int
	flag.IntVar(&freshTime, "fresh", 1, "mouse fresh time, must be greater zero")
	flag.Parse()

	currentPosition := GetMousePos()
	for {
		time.Sleep(time.Duration(freshTime) * time.Second)
		currentPosition = NextPosition(currentPosition)
		MoveMouse(currentPosition)
	}
}

func GetMousePos() robotgo.Point {
	x, y := robotgo.GetMousePos()
	return robotgo.Point{X: x, Y: y}
}

func MoveMouse(point robotgo.Point) {
	robotgo.Move(point.X, point.Y)
}

func NextPosition(point robotgo.Point) robotgo.Point {
	nextIndex := rand.Intn(len(MoveDirection))
	nextPoint := robotgo.Point{
		X: point.X + MoveDirection[nextIndex].X,
		Y: point.Y + MoveDirection[nextIndex].Y,
	}
	return reNextPosition(nextPoint)
}

func reNextPosition(point robotgo.Point) robotgo.Point {
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
