package raylib

import (
	"github.com/Spencer1O1/powder_space/v2/gfx"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) DrawText(text string, x, y, size int32, color gfx.Color) {
	rl.DrawText(text, x, y, size, toRLColor(color))
}

func toRLColor(c gfx.Color) rl.Color {
	return rl.NewColor(c.R, c.G, c.B, c.A)
}
