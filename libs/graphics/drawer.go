package graphics

import (
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/veandco/go-sdl2/sdl"
)

type DrawObject struct {
	textureBuff *sdl.Texture
}

func Draw(game core.Game, sprite Sprite) {
	renderer := game.GetRenderer()
	transform := sprite.GetTransform()
	texture := sprite.GetTexture()
	textureRect := sprite.GetTextureRect()

	quad := transform.GetGlobalBound().ToSDLRect()
	clip := textureRect.ToSDLRect()

	if (textureRect.X + textureRect.W) > float32(texture.GetWidth()) {
		f, _, _, _, _ := texture.GetSDLTexture().Query()

		tempTexture, _ := renderer.CreateTexture(f, sdl.TEXTUREACCESS_TARGET, quad.W, quad.H)
		defer tempTexture.Destroy()

		renderer.SetRenderTarget(tempTexture)

		quad.W = int32(textureRect.W - textureRect.X)
		clip.W = quad.W
		renderer.Copy(texture.GetSDLTexture(), clip, quad)

		clip.X = 0
		quad.X = int32(textureRect.W - textureRect.X)
		quad.W = int32(textureRect.X)
		clip.W = quad.W
		renderer.Copy(texture.GetSDLTexture(), clip, quad)

		quad.X = int32(transform.Position.X)
		quad.W = int32(transform.Size.X)
		renderer.SetRenderTarget(nil)
		renderer.CopyEx(tempTexture, nil, quad, transform.Rotation, nil, sdl.FLIP_NONE)
	} else {
		renderer.CopyEx(texture.GetSDLTexture(), clip, quad, transform.Rotation, nil, sdl.FLIP_NONE)
	}
}

func DrawBatch(game core.Game, texture Texture, w int32, h int32, sprites []Sprite) {
	renderer := game.GetRenderer()
	f, _, _, _, _ := texture.GetSDLTexture().Query()
	textureBuff, err := renderer.CreateTexture(f, sdl.TEXTUREACCESS_TARGET, w, h)
	if err != nil {
		panic(err)
	}
	defer textureBuff.Destroy()

	renderer.SetRenderTarget(textureBuff)
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	for _, sprite := range sprites {
		Draw(game, sprite)
	}

	renderer.SetRenderTarget(nil)
	renderer.Copy(textureBuff, nil, nil)
}

func BeginDraw(game core.Game, texture Texture, w int32, h int32) *DrawObject {
	renderer := game.GetRenderer()
	f, _, _, _, _ := texture.GetSDLTexture().Query()
	textureBuff, err := renderer.CreateTexture(f, sdl.TEXTUREACCESS_TARGET, w, h)
	if err != nil {
		panic(err)
	}

	renderer.SetRenderTarget(textureBuff)
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	return &DrawObject{
		textureBuff: textureBuff,
	}
}

func (do *DrawObject) EndDraw(game core.Game) {
	renderer := game.GetRenderer()
	renderer.SetRenderTarget(nil)
	renderer.Copy(do.textureBuff, nil, nil)

	do.textureBuff.Destroy()
}
