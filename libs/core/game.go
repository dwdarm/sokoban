package core

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game interface {
	Run()
	GetRenderer() *sdl.Renderer
	GetSceneManager() SceneManager
	GetOptions() *GameOptions
	Destroy()
}

type GameImp struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	textureBuff  *sdl.Texture
	options      *GameOptions
	sceneManager SceneManager
	input        Input
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

	if err := ttf.Init(); err != nil {
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

	textureBuff, err := renderer.CreateTexture(uint32(sdl.PIXELFORMAT_RGBA32), sdl.TEXTUREACCESS_TARGET, options.WindowWidth, options.WindowHeight)
	if err != nil {
		panic(err)
	}

	sceneManager := NewSceneManager()

	return &GameImp{
		window:       window,
		renderer:     renderer,
		textureBuff:  textureBuff,
		options:      options,
		sceneManager: sceneManager,
		input:        input,
	}
}

func (g *GameImp) Run() {
	g.sceneManager.Start()

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

		g.sceneManager.Tick(g.input, timer)
		g.input.Reset()

		renderer := g.renderer
		textureBuff := g.textureBuff

		renderer.SetRenderTarget(textureBuff)
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		g.sceneManager.Draw()

		renderer.SetRenderTarget(nil)
		renderer.Copy(textureBuff, nil, nil)
		renderer.Present()
	}

}

func (g *GameImp) GetRenderer() *sdl.Renderer {
	return g.renderer
}

func (g *GameImp) GetSceneManager() SceneManager {
	return g.sceneManager
}

func (g *GameImp) GetOptions() *GameOptions {
	return g.options
}

func (g *GameImp) Destroy() {
	if g.window != nil {
		g.window.Destroy()
	}

	if g.renderer != nil {
		g.renderer.Destroy()
	}

	if g.textureBuff != nil {
		g.textureBuff.Destroy()
	}

	g.sceneManager.Exit()

	sdl.Quit()
}
