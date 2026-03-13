package raylib

import (
	"fmt"

	"github.com/Spencer1O1/powder_space/v2/game"
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

func toRLColor(c gfxcolor.Color) rl.Color {
	return rl.NewColor(c.R, c.G, c.B, c.A)
}

func (r *Renderer) DrawText(text string, x, y, size int32, color gfxcolor.Color) {
	rl.DrawText(text, x, y, size, toRLColor(color))
}

func (r *Renderer) DrawCircle(x, y int32, radius float32, c gfxcolor.Color) {
	rl.DrawCircle(x, y, radius, toRLColor(c))
}

func (r *Renderer) DrawGame(g *game.Game) {
	for _, body := range g.World.Bodies {
		r.DrawCircle(
			int32(body.Pos.X),
			int32(body.Pos.Y),
			float32(body.Radius),
			gfxcolor.Blue,
		)
	}

	for _, p := range g.World.Particles {
		if !p.Alive {
			continue
		}

		r.DrawCircle(
			int32(p.Pos.X),
			int32(p.Pos.Y),
			float32(p.Radius),
			gfxcolor.White,
		)
	}

	r.DrawText(fmt.Sprintf("Particles: %d", len(g.World.Particles)), 20, 60, 20, gfxcolor.Gray)
}
