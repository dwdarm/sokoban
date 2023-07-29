package core

import "github.com/veandco/go-sdl2/sdl"

type Timer interface {
	Tick()
	DeltaTime() float64
	Reset()
}

type TimerImp struct {
	now       uint64
	last      uint64
	deltaTime float64
}

func NewTimer() Timer {
	return &TimerImp{
		now:       0,
		last:      0,
		deltaTime: 0,
	}
}

func (t *TimerImp) Tick() {
	t.last = t.now
	t.now = sdl.GetPerformanceCounter()
	t.deltaTime = float64((t.now - t.last)) / float64(sdl.GetPerformanceFrequency())
}

func (t *TimerImp) Reset() {
	t.last = 0
	t.now = 0
	t.deltaTime = 0.0
}

func (t *TimerImp) DeltaTime() float64 {
	return t.deltaTime
}
