package core

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Game interface {
	Run(scene Scene)
	GetRenderer() *sdl.Renderer
	Destroy()
}

type GameImp struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	options  *GameOptions
	input    Input
}

type GameOptions struct {
	Title        string
	WindowHeight int32
	WindowWidth  int32
}

func NewGame(options *GameOptions, input Input) Game {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(options.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, options.WindowWidth, options.WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}

	return &GameImp{
		window:   window,
		renderer: renderer,
		options:  options,
		input:    input,
	}
}

func (g *GameImp) Run(scene Scene) {
	scene.Start()

	timer := NewTimer()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			default:
				g.input.Handle(event)
			}
		}

		timer.Tick()

		scene.Tick(g.input, timer)
		g.input.Reset()

		renderer := g.renderer
		textureBuff, err := renderer.CreateTexture(uint32(sdl.PIXELFORMAT_RGBA32), sdl.TEXTUREACCESS_TARGET, g.options.WindowWidth, g.options.WindowHeight)
		if err != nil {
			panic(err)
		}
		defer textureBuff.Destroy()

		renderer.SetRenderTarget(textureBuff)
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		scene.Draw()

		renderer.SetRenderTarget(nil)
		renderer.Copy(textureBuff, nil, nil)
		renderer.Present()
	}

}

func (g *GameImp) GetRenderer() *sdl.Renderer {
	return g.renderer
}

func (g *GameImp) Destroy() {
	if g.window != nil {
		g.window.Destroy()
	}

	if g.renderer != nil {
		g.renderer.Destroy()
	}

	sdl.Quit()
}
