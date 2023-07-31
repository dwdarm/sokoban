package objects

import (
	"github.com/dwdarm/sokoban/libs/core"
)

type Object interface {
	Tick(input core.Input, timer core.Timer, objects []Object)
	Draw(game core.Game)
	Destroy()
	Intersect(obj Object)
	GetTransform() *core.Transform
}
