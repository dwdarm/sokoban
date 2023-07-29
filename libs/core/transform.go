package core

type Transform struct {
	Size     Vector2
	Position Vector2
	Scale    Vector2
	Rotation float64
	Forward  Vector2
}

func (t *Transform) Move(x float32, y float32) {
	t.Position.X += x
	t.Position.Y += y
}

func (t *Transform) GetGlobalBound() *Rect {
	return &Rect{
		X: t.Position.X,
		Y: t.Position.Y,
		W: t.Size.X * t.Scale.X,
		H: t.Size.Y * t.Scale.Y,
	}
}
