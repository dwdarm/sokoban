package scenes

import (
	"fmt"

	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
	"github.com/dwdarm/sokoban/objects"
)

type GameplayScene struct {
	game        core.Game
	texture     graphics.Texture
	floors      []objects.Object
	objects     []objects.Object
	targets     []objects.Object
	level       int
	targetCount int
	targetTotal int
}

func NewGameplayScene(game core.Game) core.Scene {
	texture := graphics.NewTexture()
	if err := texture.LoadFromFile(game, config.TEXTURE_PATH); err != nil {
		panic(err)
	}

	return &GameplayScene{
		game:        game,
		texture:     texture,
		level:       0,
		targetCount: 0,
		targetTotal: 0,
	}
}

func (s *GameplayScene) Start() {
	s.GenerateFloors()
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

func (s *GameplayScene) GenerateFloors() {
	for y := 0; y < config.TILE_COUNT_V; y++ {
		for x := 0; x < config.TILE_COUNT_H; x++ {
			floor := objects.NewFloor(s.texture)
			floorT := floor.GetTransform()
			floorT.Size.X = float32(config.OBJECT_TILE_SIZE)
			floorT.Size.Y = float32(config.OBJECT_TILE_SIZE)
			floorT.Position.X = float32(x * config.OBJECT_TILE_SIZE)
			floorT.Position.Y = float32(y * config.OBJECT_TILE_SIZE)
			s.floors = append(s.floors, floor)
		}
	}
}

func (s *GameplayScene) DrawFloors() {
	for _, floor := range s.floors {
		floor.Draw(s.game)
	}
}

func (s *GameplayScene) DrawObjects() {
	for _, target := range s.targets {
		target.Draw(s.game)
	}

	for _, object := range s.objects {
		object.Draw(s.game)
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
	// draw background
	s.DrawFloors()

	// draw objects
	s.DrawObjects()
}

func (s *GameplayScene) Exit() {
	if s.texture != nil {
		defer s.texture.Destory()
	}
}
