package core

import "github.com/veandco/go-sdl2/sdl"

type Rect struct {
	X float32
	Y float32
	W float32
	H float32
}

func (r *Rect) ToSDLRect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(r.X),
		Y: int32(r.Y),
		W: int32(r.W),
		H: int32(r.H),
	}
}

func (r *Rect) HasIntersection(other *Rect) bool {
	return r.ToSDLRect().HasIntersection(other.ToSDLRect())
}
