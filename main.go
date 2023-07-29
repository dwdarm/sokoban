package main

import (
	"github.com/dwdarm/sokoban/config"
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/scenes"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Sokoban", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(config.SCREEN_WIDTH), int32(config.SCREEN_HEIGHT), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	input := core.NewInput()
	input.SetBindValue("horizontal", 0.0)
	input.SetBindValue("vertical", 0.0)

	scene := scenes.NewGameplayScene(renderer)
	scene.Start()

	timer := core.NewTimer()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_a:
					input.SetBindValue("horizontal", -1.0)
					input.SetBindValue("vertical", 0)
					break
				case sdl.K_d:
					input.SetBindValue("horizontal", 1.0)
					input.SetBindValue("vertical", 0)
					break
				case sdl.K_w:
					input.SetBindValue("vertical", -1.0)
					input.SetBindValue("horizontal", 0)
					break
				case sdl.K_s:
					input.SetBindValue("vertical", 1.0)
					input.SetBindValue("horizontal", 0)
					break
				}
				break
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		timer.Tick()

		scene.Tick(input, timer)
		input.Reset()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		scene.Draw()

		renderer.Present()
	}
}
