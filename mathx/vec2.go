package mathx

import "math"

type Vec2 struct {
	X float32
	Y float32
}

func V(x, y float32) Vec2 {
	return Vec2{X: x, Y: y}
}

func V0() Vec2 {
	return Vec2{}
}

func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a Vec2) Sub(b Vec2) Vec2 {
	return Vec2{X: a.X - b.X, Y: a.Y - b.Y}
}

func (a Vec2) Mul(s float32) Vec2 {
	return Vec2{X: a.X * s, Y: a.Y * s}
}

func (a Vec2) Div(s float32) Vec2 {
	return Vec2{X: a.X / s, Y: a.Y / s}
}

func (a Vec2) Dot(b Vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}

func (a Vec2) MagSq() float32 {
	return a.X*a.X + a.Y*a.Y
}

func (a Vec2) Mag() float32 {
	return float32(math.Hypot(float64(a.X), float64(a.Y)))
}

func (v Vec2) Equal(o Vec2) bool {
	return v.X == o.X && v.Y == o.Y
}

func (a Vec2) Norm() Vec2 {
	l := a.Mag()
	if l == 0 {
		return Vec2{}
	}
	return Vec2{
		X: a.X / l,
		Y: a.Y / l,
	}
}

func (a Vec2) Lerp(b Vec2, t float32) Vec2 {
	return Vec2{
		X: a.X + (b.X-a.X)*t,
		Y: a.Y + (b.Y-a.Y)*t,
	}
}
