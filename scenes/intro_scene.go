package scenes

import (
	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
	"github.com/dwdarm/sokoban/objects"
)

type Title struct {
	font graphics.Font
	text graphics.Text
}

func NewTitle(game core.Game, fontPath string) *Title {
	font := graphics.NewFont()
	if err := font.LoadFromFile(config.FONT_PATH, 40); err != nil {
		panic(err)
	}

	text := graphics.NewText(game, font, "SOKOBAN", &core.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	})
	text.SetOutlineColor(&core.Color{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	})
	text.SetOutlineSize(4)
	text.BuildSurface()

	text.SetPosition(&core.Vector2{
		X: (float32(game.GetOptions().WindowWidth) - text.GetSize().X) / 2,
		Y: 128,
	})

	return &Title{
		font: font,
		text: text,
	}
}

func (t *Title) Draw() {
	t.text.Draw()
}

func (t *Title) Destroy() {
	if t.font != nil {
		t.font.Destory()
	}

	if t.text != nil {
		t.text.Destroy()
	}
}

type IntroScene struct {
	game     core.Game
	texture  graphics.Texture
	title    *objects.Text
	subtitle *objects.Text
	floors   []objects.Object
	objs     []objects.Object
}

func NewIntroScene(game core.Game) core.Scene {
	texture := graphics.NewTexture()
	if err := texture.LoadFromFile(game, config.TEXTURE_PATH); err != nil {
		panic(err)
	}

	title := objects.NewText(game, config.FONT_PATH, "SOKOBAN", 40)
	title.GetTextHandle().SetPosition(
		&core.Vector2{
			X: (float32(game.GetOptions().WindowWidth) - title.GetTextHandle().GetSize().X) / 2,
			Y: 128,
		},
	)

	subtitle := objects.NewText(game, config.FONT_THIN_PATH, "Press Space to play", 16)
	subtitle.GetTextHandle().SetPosition(
		&core.Vector2{
			X: (float32(game.GetOptions().WindowWidth) - subtitle.GetTextHandle().GetSize().X) / 2,
			Y: float32(game.GetOptions().WindowHeight) - 128,
		},
	)

	objs := []objects.Object{}
	obj := objects.NewBox(texture)
	obj.GetTransform().Position = core.Vector2{
		X: 48,
		Y: 48,
	}
	obj.GetTransform().Rotation = 30
	objs = append(objs, obj)

	obj = objects.NewBox(texture)
	obj.GetTransform().Position = core.Vector2{
		X: float32(game.GetOptions().WindowWidth) - 128,
		Y: 48,
	}
	obj.GetTransform().Rotation = -20
	objs = append(objs, obj)

	obj = objects.NewBox(texture)
	obj.GetTransform().Position = core.Vector2{
		X: 96,
		Y: 224,
	}
	obj.GetTransform().Rotation = 10
	objs = append(objs, obj)

	obj = objects.NewBox(texture)
	obj.GetTransform().Position = core.Vector2{
		X: float32(game.GetOptions().WindowWidth) - 200,
		Y: 300,
	}
	obj.GetTransform().Rotation = -24
	objs = append(objs, obj)

	return &IntroScene{
		game:     game,
		texture:  texture,
		title:    title,
		subtitle: subtitle,
		objs:     objs,
	}
}

func (s *IntroScene) Start() {
	s.GenerateFloors()
}

func (s *IntroScene) Tick(input core.Input, timer core.Timer) {
	if input.GetValue("confirm") == 1.0 {
		s.game.GetSceneManager().SwitchScene(NewGameplayScene(s.game))
	}
}

func (s *IntroScene) Draw() {
	s.DrawFloors()
	s.DrawObjs()
	s.title.Draw()
	s.subtitle.Draw()
}

func (s *IntroScene) Exit() {
	if s.texture != nil {
		defer s.texture.Destory()
	}
}

func (s *IntroScene) GenerateFloors() {
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

func (s *IntroScene) DrawFloors() {
	for _, floor := range s.floors {
		floor.Draw(s.game)
	}
}

func (s *IntroScene) DrawObjs() {
	for _, obj := range s.objs {
		obj.Draw(s.game)
	}
}
