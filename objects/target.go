package objects

import (
	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

type Target struct {
	Sprite  graphics.Sprite
	Texture graphics.Texture
}

func NewTarget(texture graphics.Texture) Object {
	p := &Target{
		Texture: texture,
	}

	p.Sprite = graphics.NewSprite()
	p.Sprite.SetTexture(p.Texture)
	p.Sprite.SetTextureRect(1*TEXTURE_TILE_SIZE, 3*TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE)
	p.Sprite.GetTransform().Size.X = float32(config.OBJECT_TILE_SIZE)
	p.Sprite.GetTransform().Size.Y = float32(config.OBJECT_TILE_SIZE)

	return p
}

func (t *Target) Tick(input core.Input, timer core.Timer, objects []Object) {

}

func (t *Target) Draw(renderer *sdl.Renderer) {
	t.Sprite.Draw(renderer)
}

func (t *Target) Intersect(Object) {

}

func (t *Target) GetTransform() *core.Transform {
	return t.Sprite.GetTransform()
}

func (t *Target) Destroy() {

}
