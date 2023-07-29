package scenes

import (
	"fmt"

	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
	"github.com/dwdarm/sokoban/objects"
	"github.com/veandco/go-sdl2/sdl"
)

type GameplayScene struct {
	renderer    *sdl.Renderer
	texture     graphics.Texture
	objects     []objects.Object
	targets     []objects.Object
	level       int
	targetCount int
	targetTotal int
}

func NewGameplayScene(renderer *sdl.Renderer) core.Scene {
	texture := graphics.NewTexture()
	if err := texture.LoadFromFile(renderer, config.TEXTURE_PATH); err != nil {
		panic(err)
	}

	return &GameplayScene{
		renderer:    renderer,
		texture:     texture,
		level:       0,
		targetCount: 0,
		targetTotal: 0,
	}
}

func (s *GameplayScene) Start() {
	s.SpawnObjects()
}

func (s *GameplayScene) NextLevel() {
	s.objects = nil
	s.targets = nil
	s.targetCount = 0
	s.targetTotal = 0
	s.level++
	s.SpawnObjects()
}

func (s *GameplayScene) Tick(input core.Input, timer core.Timer) {
	s.targetCount = 0

	for _, object := range s.objects {
		object.Tick(input, timer, s.objects)

		if _, isBox := object.(*objects.Box); isBox {
			for _, target := range s.targets {
				boxT := object.GetTransform()
				targetT := target.GetTransform()
				if boxT.Position.X == targetT.Position.X && boxT.Position.Y == targetT.Position.Y {
					s.targetCount++
				}
			}
		}
	}

	if s.targetCount == s.targetTotal {
		if s.level < len(config.LEVELS) {
			s.NextLevel()
		} else {
			fmt.Println("FINISH")
		}
	}

}

func (s *GameplayScene) DrawBackground() {
	for y := 0; y < config.TILE_COUNT_V; y++ {
		for x := 0; x < config.TILE_COUNT_H; x++ {
			clip := sdl.Rect{
				X: 11 * int32(config.TEXTURE_TILE_SIZE),
				Y: 6 * int32(config.TEXTURE_TILE_SIZE),
				W: int32(config.TEXTURE_TILE_SIZE),
				H: int32(config.TEXTURE_TILE_SIZE),
			}
			quad := sdl.Rect{
				X: int32(x) * int32(config.OBJECT_TILE_SIZE),
				Y: int32(y) * int32(config.OBJECT_TILE_SIZE),
				W: int32(config.OBJECT_TILE_SIZE),
				H: int32(config.OBJECT_TILE_SIZE),
			}

			s.renderer.Copy(s.texture.GetSDLTexture(), &clip, &quad)
		}
	}
}

func (s *GameplayScene) DrawObjects() {
	for _, target := range s.targets {
		target.Draw(s.renderer)
	}

	for _, object := range s.objects {
		object.Draw(s.renderer)
	}
}

func (s *GameplayScene) SpawnObjects() {
	for y := 0; y < config.TILE_COUNT_V; y++ {
		for x := 0; x < config.TILE_COUNT_H; x++ {
			objectType := config.LEVELS[s.level][x+(y*config.TILE_COUNT_V)]
			if objectType == config.LEVEL_BOX_ON_TARGET {
				s.AppendObject(config.LEVEL_BOX, x, y)
				s.AppendObject(config.LEVEL_TARGET, x, y)
			} else {
				s.AppendObject(objectType, x, y)
			}
		}
	}
}

func (s *GameplayScene) AppendObject(objectType int, x int, y int) {
	var obj objects.Object
	switch objectType {
	case config.LEVEL_PLAYER:
		obj = objects.NewPlayer(s.texture)
	case config.LEVEL_WALL:
		obj = objects.NewWall(s.texture)
	case config.LEVEL_BOX:
		obj = objects.NewBox(s.texture)
	case config.LEVEL_TARGET:
		obj = objects.NewTarget(s.texture)
		s.targetTotal++
	}

	if obj != nil {
		obj.GetTransform().Position.X = float32(x * config.OBJECT_TILE_SIZE)
		obj.GetTransform().Position.Y = float32(y * config.OBJECT_TILE_SIZE)

		if _, ok := obj.(*objects.Target); ok {
			s.targets = append(s.targets, obj)
		} else {
			s.objects = append(s.objects, obj)
		}
	}
}

func (s *GameplayScene) Draw() {
	f, _, _, _, _ := s.texture.GetSDLTexture().Query()
	textureBg, err := s.renderer.CreateTexture(f, sdl.TEXTUREACCESS_TARGET, int32(config.SCREEN_WIDTH), int32(config.SCREEN_HEIGHT))
	if err != nil {
		panic(err)
	}
	defer textureBg.Destroy()

	s.renderer.SetRenderTarget(textureBg)
	s.renderer.SetDrawColor(0, 0, 0, 255)
	s.renderer.Clear()

	// draw background
	s.DrawBackground()

	// draw objects
	s.DrawObjects()

	s.renderer.SetRenderTarget(nil)
	s.renderer.Copy(textureBg, nil, nil)
}
