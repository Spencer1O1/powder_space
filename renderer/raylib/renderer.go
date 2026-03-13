package raylib

import (
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func GetFrameTime() float64 {
	return float64(rl.GetFrameTime())
}

func (r *Renderer) DrawText(text string, x, y, size int32, color gfxcolor.Color) {
	rl.DrawText(text, x, y, size, toRLColor(color))
}

func toRLColor(c gfxcolor.Color) rl.Color {
	return rl.NewColor(c.R, c.G, c.B, c.A)
}
