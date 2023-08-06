package objects

import (
	"github.com/dwdarm/sokoban/libs/core"
	"github.com/dwdarm/sokoban/libs/graphics"
)

type Text struct {
	font graphics.Font
	text graphics.Text
}

func NewText(game core.Game, fontPath string, txt string, size int) *Text {
	font := graphics.NewFont()
	if err := font.LoadFromFile(fontPath, size); err != nil {
		panic(err)
	}

	text := graphics.NewText(game, font, txt, &core.Color{
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
	text.SetOutlineSize(2)
	text.BuildSurface()

	return &Text{
		font: font,
		text: text,
	}
}

func (t *Text) Draw() {
	t.text.Draw()
}

func (t *Text) GetTextHandle() graphics.Text {
	return t.text
}

func (t *Text) Destroy() {
	if t.font != nil {
		t.font.Destory()
	}

	if t.text != nil {
		t.text.Destroy()
	}
}
