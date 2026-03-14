package raylib

import (
	"fmt"

	"github.com/Spencer1O1/powder_space/v2/game"
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	"github.com/Spencer1O1/powder_space/v2/mathx"
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

func (r *Renderer) DrawLine(x1, y1, x2, y2 int32, c gfxcolor.Color) {
	rl.DrawLine(x1, y1, x2, y2, toRLColor(c))
}

func (r *Renderer) DrawGame(g *game.Game, pointerPos mathx.Vec2) {
	for _, body := range g.World.Bodies {
		r.DrawCircle(
			int32(body.Pos.X),
			int32(body.Pos.Y),
			float32(body.Radius),
			gfxcolor.Blue,
		)
	}

	// Particles get drawn above mouse control lines
	if g.AnchorSet {
		// anchor marker
		r.DrawCircle(int32(g.Anchor.X), int32(g.Anchor.Y), 4, gfxcolor.Gray)

		// sling line from anchor to current mouse
		r.DrawLine(
			int32(g.Anchor.X),
			int32(g.Anchor.Y),
			int32(pointerPos.X),
			int32(pointerPos.Y),
			gfxcolor.Gray,
		)

		// predicted velocity direction line
		vel := g.LaunchVelocityFromPosition(pointerPos)

		end := pointerPos.Add(vel.Mul(0.6))

		r.DrawLine(
			int32(pointerPos.X),
			int32(pointerPos.Y),
			int32(end.X),
			int32(end.Y),
			gfxcolor.White,
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

	if len(g.World.Bodies) > 0 {
		body := g.World.Bodies[0]
		r.DrawText(fmt.Sprintf("Body Mass: %.0f", body.Mass), 20, 85, 20, gfxcolor.Gray)
		r.DrawText(fmt.Sprintf("Body Radius: %.1f", body.Radius), 20, 110, 20, gfxcolor.Gray)
	}

	// r.DrawText(
	// 	fmt.Sprintf("Sim Speed: %.2fx", a.timeScale),
	// 	20, 140, 24, gfxcolor.White,
	// )
}
