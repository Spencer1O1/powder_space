package scene

import (
	"fmt"

	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/game"
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
)

type HudParams struct {
	TimeScale float32
}

func DrawHud(r *rr.Renderer, g *game.Game, params HudParams) {
	r.DrawText(content.TitleString, 20, 20, 32, gfxcolor.White)
	r.DrawText(fmt.Sprintf("Particles: %d", len(g.World.Particles)), 20, 60, 20, gfxcolor.Gray)

	// if len(g.World.Bodies) > 0 {
	// 	body := g.World.Bodies[0]
	// 	r.DrawText(fmt.Sprintf("Body Mass: %.0f", body.Mass), 20, 85, 20, gfxcolor.Gray)
	// 	r.DrawText(fmt.Sprintf("Body Radius: %.1f", body.Radius), 20, 110, 20, gfxcolor.Gray)
	// }

	r.DrawText(
		fmt.Sprintf("Sim Speed: %.2fx", params.TimeScale*4),
		20, 135, 20, gfxcolor.Gray,
	)
}
