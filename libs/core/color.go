package core

import "github.com/veandco/go-sdl2/sdl"

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (c *Color) ToSDLColor() *sdl.Color {
	return &sdl.Color{
		R: c.R,
		G: c.G,
		B: c.B,
		A: c.A,
	}
}
