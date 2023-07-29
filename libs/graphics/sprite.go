package graphics

import (
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite interface {
	SetTexture(texture Texture)
	SetTextureRect(x int32, y int32, w int32, h int32)

	GetTransform() *core.Transform

	Draw(renderer *sdl.Renderer)
}

type SpriteImp struct {
	texture     Texture
	textureRect *sdl.Rect
	core.Transform
}

func NewSprite() Sprite {
	s := &SpriteImp{}
	s.texture = nil
	s.textureRect = &sdl.Rect{0, 0, 0, 0}
	s.Scale.X = 1.0
	s.Scale.Y = 1.0

	return s
}

func (s *SpriteImp) SetTexture(texture Texture) {
	s.texture = texture
	s.textureRect = &sdl.Rect{0, 0, 0, 0}
	s.Size.X = float32(texture.GetWidth())
	s.Size.Y = float32(texture.GetHeight())
}

func (s *SpriteImp) SetTextureRect(x int32, y int32, w int32, h int32) {
	s.textureRect = &sdl.Rect{x, y, w, h}
	s.Size.X = float32(w)
	s.Size.Y = float32(h)
}

func (s *SpriteImp) GetTransform() *core.Transform {
	return &s.Transform
}

func (s *SpriteImp) Draw(renderer *sdl.Renderer) {
	Draw(renderer, s.texture, s.textureRect, &s.Transform)
}
