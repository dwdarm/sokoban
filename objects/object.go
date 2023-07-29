package objects

import (
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/veandco/go-sdl2/sdl"
)

type Object interface {
	Tick(input core.Input, timer core.Timer, objects []Object)
	Draw(renderer *sdl.Renderer)
	Destroy()
	Intersect(obj Object)
	GetTransform() *core.Transform
}
