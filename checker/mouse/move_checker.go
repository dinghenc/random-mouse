package mouse

import (
	"time"

	"github.com/dinghenc/random-mouse/checker"
	"github.com/dinghenc/random-mouse/utils"
	"github.com/go-vgo/robotgo"
)

const (
	defaultMoveDuration = time.Second
	defaultMoveDistance = 100
)

type MoveChecker struct {
	moveDuration   time.Duration
	moveDistance   int
	ch             chan struct{}
	changed        bool
	lastMousePoint robotgo.Point
}

type MoveOption func(m *MoveChecker)

func WithTickerDuration(duration time.Duration) MoveOption {
	return func(m *MoveChecker) {
		m.moveDuration = duration
	}
}

func WithMoveDistance(distance int) MoveOption {
	return func(m *MoveChecker) {
		m.moveDistance = distance
	}
}

func NewMoveChecker(opts ...MoveOption) checker.Checker {
	m := &MoveChecker{
		moveDuration: defaultMoveDuration,
		moveDistance: defaultMoveDistance,
		ch:           make(chan struct{}),
	}
	for _, o := range opts {
		o(m)
	}

	go func() {
		ticker := time.NewTicker(m.moveDuration)
		for {
			select {
			case <-ticker.C:
				m.doCheck()
			}
		}
	}()

	return m
}

func (m *MoveChecker) Check() bool {
	return m.changed
}

func (m *MoveChecker) Changed() <-chan struct{} {
	return m.ch
}

func (m *MoveChecker) doCheck() {
	curMousePoint := utils.GetMousePos()
	defer func() {
		m.lastMousePoint = curMousePoint
	}()
	if utils.EmptyMousePos(m.lastMousePoint) {
		return
	}

	if utils.DistanceBetweenPos(m.lastMousePoint, curMousePoint) > float64(m.moveDistance) {
		m.changed = true
		m.ch <- struct{}{}
	}
}
