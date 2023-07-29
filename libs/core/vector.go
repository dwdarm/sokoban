package core

type Vector2 struct {
	X float32
	Y float32
}

func (v *Vector2) Clamp(min float32, max float32) {
	if v.X < min {
		v.X = min
	}

	if v.X > max {
		v.X = max
	}

	if v.Y < min {
		v.Y = min
	}

	if v.Y > max {
		v.Y = max
	}
}
