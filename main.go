package main

import (
	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/scenes"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	input := core.NewInput()
	input.RegisterInput("horizontal", &core.KeyboardBinding{
		MinEventKey: sdl.K_a,
		MinValue:    -1.0,
		MaxEventKey: sdl.K_d,
		MaxValue:    1.0,
	})
	input.RegisterInput("vertical", &core.KeyboardBinding{
		MinEventKey: sdl.K_w,
		MinValue:    -1.0,
		MaxEventKey: sdl.K_s,
		MaxValue:    1.0,
	})
	input.RegisterInput("confirm", &core.KeyboardBinding{
		MaxEventKey: sdl.K_SPACE,
		MaxValue:    1.0,
	})

	game := core.NewGame(&core.GameOptions{
		Title:        "Sokoban",
		WindowWidth:  int32(config.SCREEN_WIDTH),
		WindowHeight: int32(config.SCREEN_HEIGHT),
	}, input)
	defer game.Destroy()

	sceneManager := game.GetSceneManager()
	sceneManager.SwitchScene(scenes.NewIntroScene(game))

	game.Run()
}
