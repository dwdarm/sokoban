package graphics

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture interface {
	LoadFromFile(renderer *sdl.Renderer, path string) error
	GetSDLTexture() *sdl.Texture
	GetWidth() int32
	GetHeight() int32
	Destory()
}

type TextureImp struct {
	sdlTexture *sdl.Texture
	width      int32
	height     int32
}

func NewTexture() Texture {
	return &TextureImp{
		sdlTexture: nil,
		width:      -1,
		height:     -1,
	}
}

func (t *TextureImp) LoadFromFile(renderer *sdl.Renderer, path string) error {
	img, err := img.Load(path)
	if err != nil {
		return err
	}
	defer img.Free()

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return err
	}

	t.Destory()
	t.sdlTexture = tex
	t.width = img.W
	t.height = img.H

	return nil
}

func (t *TextureImp) GetSDLTexture() *sdl.Texture {
	return t.sdlTexture
}

func (t *TextureImp) GetWidth() int32 {
	return t.width
}

func (t *TextureImp) GetHeight() int32 {
	return t.height
}

func (t *TextureImp) Destory() {
	if t.sdlTexture != nil {
		t.sdlTexture.Destroy()
	}

	t.width = -1
	t.height = -1
}
