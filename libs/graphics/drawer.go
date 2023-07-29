package graphics

import (
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/veandco/go-sdl2/sdl"
)

func Draw(renderer *sdl.Renderer, texture Texture, textureRect *sdl.Rect, transform *core.Transform) {
	quad := sdl.Rect{
		X: int32(transform.Position.X),
		Y: int32(transform.Position.Y),
		W: int32(transform.Size.X),
		H: int32(transform.Size.Y),
	}
	clip := sdl.Rect{
		X: textureRect.X,
		Y: textureRect.Y,
		W: textureRect.W,
		H: textureRect.H,
	}

	if (textureRect.X + textureRect.W) > texture.GetWidth() {
		f, _, _, _, _ := texture.GetSDLTexture().Query()

		tempTexture, _ := renderer.CreateTexture(f, sdl.TEXTUREACCESS_TARGET, quad.W, quad.H)
		defer tempTexture.Destroy()

		renderer.SetRenderTarget(tempTexture)

		quad.W = textureRect.W - textureRect.X
		clip.W = quad.W
		renderer.Copy(texture.GetSDLTexture(), &clip, &quad)

		clip.X = 0
		quad.X = textureRect.W - textureRect.X
		quad.W = textureRect.X
		clip.W = quad.W
		renderer.Copy(texture.GetSDLTexture(), &clip, &quad)

		quad.X = int32(transform.Position.X)
		quad.W = int32(transform.Size.X)
		renderer.SetRenderTarget(nil)
		renderer.CopyEx(tempTexture, nil, &quad, transform.Rotation, nil, sdl.FLIP_NONE)
	} else {
		renderer.CopyEx(texture.GetSDLTexture(), &clip, &quad, transform.Rotation, nil, sdl.FLIP_NONE)
	}
}
