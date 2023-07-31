package objects

import (
	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
)

type Wall struct {
	Sprite  graphics.Sprite
	Texture graphics.Texture
}

func NewWall(texture graphics.Texture) Object {
	p := &Wall{
		Texture: texture,
	}

	p.Sprite = graphics.NewSprite()
	p.Sprite.SetTexture(p.Texture)
	p.Sprite.SetTextureRect(7*TEXTURE_TILE_SIZE, 7*TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE)
	p.Sprite.GetTransform().Size.X = float32(config.OBJECT_TILE_SIZE)
	p.Sprite.GetTransform().Size.Y = float32(config.OBJECT_TILE_SIZE)

	return p
}

func (w *Wall) Tick(input core.Input, timer core.Timer, objects []Object) {

}

func (w *Wall) Draw(game core.Game) {
	w.Sprite.Draw(game)
}

func (w *Wall) Intersect(Object) {

}

func (w *Wall) GetTransform() *core.Transform {
	return w.Sprite.GetTransform()
}

func (w *Wall) Destroy() {

}
