package mathx

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func V(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a Vec2) Sub(b Vec2) Vec2 {
	return Vec2{X: a.X - b.X, Y: a.Y - b.Y}
}

func (a Vec2) Mul(s float64) Vec2 {
	return Vec2{X: a.X * s, Y: a.Y * s}
}

func (a Vec2) Dot(b Vec2) float64 {
	return a.X*b.X + a.Y*b.Y
}

func (a Vec2) MagSq() float64 {
	return a.X*a.X + a.Y*a.Y
}

func (a Vec2) Mag() float64 {
	return math.Hypot(a.X, a.Y)
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
