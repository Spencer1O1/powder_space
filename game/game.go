package game

import (
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/sim"
)

type Game struct {
	World *sim.World
}

func NewGame() *Game {
	return &Game{
		World: sim.NewWorld(),
	}
}

func (g *Game) Update(dt float64) {
	g.World.Step(dt)
}

func (g *Game) SpawnDust(x, y float64) {
	g.World.SpawnParticle(mathx.V(x, y), mathx.V(0, 0))
}
