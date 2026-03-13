package game

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/sim"
)

type Game struct {
	World        *sim.World
	SelectedTool Tool
}

func NewGame() *Game {
	return &Game{
		World:        sim.NewWorld(),
		SelectedTool: ToolDust,
	}
}

func (g *Game) Update(dt float64) {
	g.World.Step(dt)
}

func (g *Game) SpawnPowder(x, y, vx, vy float64) {
	g.World.SpawnParticle(
		mathx.V(x, y),
		mathx.V(vx, vy),
		content.MaterialDust,
		10.0,
	)
}

func (g *Game) ClearParticles() {
	g.World.ClearParticles()
}

func (g *Game) Reset() {
	g.World = sim.NewWorld()
}
