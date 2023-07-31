package objects

import (
	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
)

type Box struct {
	Sprite  graphics.Sprite
	Texture graphics.Texture
}

func NewBox(texture graphics.Texture) Object {
	p := &Box{
		Texture: texture,
	}

	p.Sprite = graphics.NewSprite()
	p.Sprite.SetTexture(p.Texture)
	p.Sprite.SetTextureRect(6*TEXTURE_TILE_SIZE, 0*TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE)
	p.Sprite.GetTransform().Size.X = float32(config.OBJECT_TILE_SIZE)
	p.Sprite.GetTransform().Size.Y = float32(config.OBJECT_TILE_SIZE)

	return p
}

func (b *Box) Tick(input core.Input, timer core.Timer, objects []Object) {

}

func (b *Box) Push(x float32, y float32, objects []Object) bool {
	transform := b.GetTransform()

	transform.Move(x, y)

	for _, obj := range objects {
		if obj != b {
			if _, isPlayer := obj.(*Player); !isPlayer {
				if _, isTarget := obj.(*Target); !isTarget {
					objTransform := obj.GetTransform()

					if transform.GetGlobalBound().HasIntersection(objTransform.GetGlobalBound()) {
						if transform.Forward.X == 1 && transform.Forward.Y == 0 {
							transform.Position.X = objTransform.Position.X - transform.Size.X
						} else if transform.Forward.X == -1 && transform.Forward.Y == 0 {
							transform.Position.X = objTransform.Position.X + objTransform.Size.X
						}

						if transform.Forward.X == 0 && transform.Forward.Y == 1 {
							transform.Position.Y = objTransform.Position.Y - transform.Size.Y
						} else if transform.Forward.X == 0 && transform.Forward.Y == -1 {
							transform.Position.Y = objTransform.Position.Y + objTransform.Size.Y
						}

						return false
					}
				}
			}
		}
	}

	return true
}

func (b *Box) Draw(game core.Game) {
	b.Sprite.Draw(game)
}

func (b *Box) Intersect(obj Object) {

}

func (b *Box) GetTransform() *core.Transform {
	return b.Sprite.GetTransform()
}

func (b *Box) Destroy() {

}
