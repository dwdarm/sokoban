package objects

import (
	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
)

var CHAR_SIZE int32 = 36
var CHAR_WALK_SPEED int32 = 320

var TEXTURE_TILE_SIZE int32 = int32(config.TEXTURE_TILE_SIZE)

type Player struct {
	Sprite     graphics.Sprite
	Texture    graphics.Texture
	Animation  core.Vector2
	AnimationX float32
	Movement   core.Vector2
}

func NewPlayer(texture graphics.Texture) Object {
	p := &Player{
		Texture:    texture,
		AnimationX: 0,
	}

	p.Sprite = graphics.NewSprite()
	p.Sprite.SetTexture(p.Texture)
	p.Sprite.SetTextureRect(0*TEXTURE_TILE_SIZE, 4*TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE)
	p.Sprite.GetTransform().Size.X = float32(CHAR_SIZE)
	p.Sprite.GetTransform().Size.Y = float32(CHAR_SIZE)

	p.Animation.X = 0.0
	p.Animation.Y = 0.0

	return p
}

func (p *Player) Tick(input core.Input, timer core.Timer, objects []Object) {
	transform := p.GetTransform()

	p.Movement = core.Vector2{
		X: float32(input.GetValue("horizontal")) * float32(CHAR_WALK_SPEED) * float32(timer.DeltaTime()),
		Y: float32(input.GetValue("vertical")) * float32(CHAR_WALK_SPEED) * float32(timer.DeltaTime()),
	}

	p.Movement.Clamp(float32(-CHAR_WALK_SPEED), float32(CHAR_WALK_SPEED))

	if p.Movement.X != 0.0 {
		p.Movement.Y = 0.0
	} else if p.Movement.Y != 0.0 {
		p.Movement.X = 0.0
	}

	if p.Movement.X > 0.0 {
		p.Animation.X = 0
		p.Animation.Y = 6
		transform.Forward.X = 1
		transform.Forward.Y = 0
	} else if p.Movement.X < 0.0 {
		p.Animation.X = 3
		p.Animation.Y = 6
		transform.Forward.X = -1
		transform.Forward.Y = 0
	} else if p.Movement.Y > 0.0 {
		p.Animation.X = 0
		p.Animation.Y = 4
		transform.Forward.X = 0
		transform.Forward.Y = 1
	} else if p.Movement.Y < 0.0 {
		p.Animation.X = 3
		p.Animation.Y = 4
		transform.Forward.X = 0
		transform.Forward.Y = -1
	}

	if p.Movement.X != 0.0 || p.Movement.Y != 0.0 {
		p.AnimationX += float32(timer.DeltaTime() * 16)
		if p.AnimationX > 3.0 {
			p.AnimationX = 0.0
		}
	}

	transform.Move(p.Movement.X, p.Movement.Y)

	p.Sprite.SetTextureRect(int32(p.Animation.X+p.AnimationX)*TEXTURE_TILE_SIZE, int32(p.Animation.Y)*TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE, TEXTURE_TILE_SIZE)
	transform.Size.X = float32(CHAR_SIZE)
	transform.Size.Y = float32(CHAR_SIZE)

	for _, obj := range objects {
		objTransform := obj.GetTransform()

		if transform.GetGlobalBound().HasIntersection(objTransform.GetGlobalBound()) {
			if _, ok := obj.(*Player); !ok {
				if _, ok := obj.(*Wall); ok {
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
				} else if _, ok := obj.(*Box); ok {
					objTransform.Forward.X = transform.Forward.X
					objTransform.Forward.Y = transform.Forward.Y
					if !obj.(*Box).Push(p.Movement.X, p.Movement.Y, objects) {
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
					}
				}
			}
		}
	}
}

func (p *Player) Draw(game core.Game) {
	p.Sprite.Draw(game)
}

func (p *Player) Intersect(obj Object) {

}

func (p *Player) GetTransform() *core.Transform {
	return p.Sprite.GetTransform()
}

func (p *Player) Destroy() {

}
