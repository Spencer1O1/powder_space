package raylib

import (
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer struct {
	Width  int32
	Height int32
}

func NewRenderer(width, height int32) *Renderer {
	return &Renderer{
		Width:  width,
		Height: height,
	}
}

func GetFrameTime() float32 {
	return rl.GetFrameTime()
}

func toRLColor(c gfxcolor.Color) rl.Color {
	return rl.NewColor(c.R, c.G, c.B, c.A)
}

func (r *Renderer) DrawText(text string, x, y, size int32, color gfxcolor.Color) {
	rl.DrawText(text, x, y, size, toRLColor(color))
}

func (r *Renderer) DrawPixel(x, y int32, c gfxcolor.Color) {
	rl.DrawPixel(x, y, toRLColor(c))
}

func (r *Renderer) DrawLine(x1, y1, x2, y2 int32, c gfxcolor.Color) {
	rl.DrawLine(x1, y1, x2, y2, toRLColor(c))
}

func (r *Renderer) DrawRect(x, y, w, h int32, c gfxcolor.Color) {
	rl.DrawRectangle(x, y, w, h, toRLColor(c))
}

func (r *Renderer) DrawCircle(x, y int32, radius float32, c gfxcolor.Color) {
	rl.DrawCircle(x, y, radius, toRLColor(c))
}

func (r *Renderer) DrawParticle(pos mathx.Vec2, radius float32, c gfxcolor.Color) {
	x := int32(pos.X)
	y := int32(pos.Y)

	if x < 0 || y < 0 || x >= r.Width || y >= r.Height {
		return
	}

	switch {
	case radius <= 0.75:
		r.DrawPixel(x, y, c)
	case radius <= 1.5:
		r.DrawRect(x-1, y-1, 2, 2, c)
	default:
		r.DrawCircle(x, y, radius, c)
	}
}
