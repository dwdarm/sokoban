package objects

import (
	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
)

type Floor struct {
	Sprite  graphics.Sprite
	Texture graphics.Texture
}

func NewFloor(texture graphics.Texture) Object {
	p := &Floor{
		Texture: texture,
	}

	p.Sprite = graphics.NewSprite()
	p.Sprite.SetTexture(p.Texture)
	p.Sprite.SetTextureRect(11*TEXTURE_TILE_SIZE, 6*TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE)
	p.Sprite.GetTransform().Size.X = float32(config.OBJECT_TILE_SIZE)
	p.Sprite.GetTransform().Size.Y = float32(config.OBJECT_TILE_SIZE)

	return p
}

func (w *Floor) Tick(input core.Input, timer core.Timer, objects []Object) {

}

func (w *Floor) Draw(game core.Game) {
	w.Sprite.Draw(game)
}

func (w *Floor) Intersect(Object) {

}

func (w *Floor) GetTransform() *core.Transform {
	return w.Sprite.GetTransform()
}

func (w *Floor) Destroy() {

}
