package scene

import (
	"github.com/Spencer1O1/powder_space/v2/game"
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
)

type GameViewParams struct {
	PointerPos mathx.Vec2
}

func DrawGame(r *rr.Renderer, g *game.Game, params GameViewParams) {
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
			int32(params.PointerPos.X),
			int32(params.PointerPos.Y),
			gfxcolor.Gray,
		)

		// predicted velocity direction line
		vel := g.LaunchVelocityFromPosition(params.PointerPos)

		end := params.PointerPos.Add(vel.Mul(0.6))

		r.DrawLine(
			int32(params.PointerPos.X),
			int32(params.PointerPos.Y),
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
}
