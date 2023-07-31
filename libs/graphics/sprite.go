package graphics

import (
	"github.com/dwdarm/sokoban/libs/core"
)

type Sprite interface {
	SetTexture(texture Texture)
	SetTextureRect(x int32, y int32, w int32, h int32)

	GetTransform() *core.Transform

	GetTexture() Texture
	GetTextureRect() *core.Rect

	Draw(game core.Game)
}

type SpriteImp struct {
	texture     Texture
	textureRect core.Rect
	core.Transform
}

func NewSprite() Sprite {
	s := &SpriteImp{}
	s.texture = nil
	s.textureRect.X = 0
	s.textureRect.Y = 0
	s.textureRect.W = 0
	s.textureRect.H = 0
	s.Scale.X = 1.0
	s.Scale.Y = 1.0

	return s
}

func (s *SpriteImp) SetTexture(texture Texture) {
	s.texture = texture
	s.textureRect.X = 0
	s.textureRect.Y = 0
	s.textureRect.W = 0
	s.textureRect.H = 0
	s.Size.X = float32(texture.GetWidth())
	s.Size.Y = float32(texture.GetHeight())
}

func (s *SpriteImp) SetTextureRect(x int32, y int32, w int32, h int32) {
	s.textureRect.X = float32(x)
	s.textureRect.Y = float32(y)
	s.textureRect.W = float32(w)
	s.textureRect.H = float32(h)
	s.Size.X = float32(w)
	s.Size.Y = float32(h)
}

func (s *SpriteImp) GetTransform() *core.Transform {
	return &s.Transform
}

func (s *SpriteImp) GetTexture() Texture {
	return s.texture
}

func (s *SpriteImp) GetTextureRect() *core.Rect {
	return &s.textureRect
}

func (s *SpriteImp) Draw(game core.Game) {
	Draw(game, s)
}
