package checker

type Checker interface {
	Check() bool
	Changed() <-chan struct{}
}
