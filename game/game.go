package game

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/sim"
)

type Game struct {
	World        *sim.World
	SelectedTool Tool

	AnchorSet bool
	Anchor    mathx.Vec2

	SelectedMaterial content.MaterialID
	SpawnMass        float32
	LaunchStrength   float32
	MaxLaunchSpeed   float32
}

func NewGame() *Game {
	return &Game{
		World:            sim.NewWorld(),
		SelectedTool:     ToolPowder,
		SelectedMaterial: content.MaterialDust,
		SpawnMass:        10.0,
		LaunchStrength:   3.0,
		MaxLaunchSpeed:   500.0,
	}
}

func (g *Game) Update(dt float32) {
	// frame-based, optional for now
}

func (g *Game) FixedUpdate(dt float32) {
	g.World.Step(dt)
}

func (g *Game) ClearParticles() {
	g.World.ClearParticles()
}

func (g *Game) Reset() {
	g.World = sim.NewWorld()
}

func (g *Game) SetAnchor(pos mathx.Vec2) {
	g.Anchor = pos
	g.AnchorSet = true
}

func (g *Game) ResetAnchor() {
	g.Anchor = mathx.Vec2{}
	g.AnchorSet = false
}

func (g *Game) LaunchVelocityFromPosition(pos mathx.Vec2) mathx.Vec2 {
	if !g.AnchorSet {
		return mathx.Vec2{}
	}

	// slingshot: pull away from anchor, launch in opposite direction
	v := g.Anchor.Sub(pos).Mul(g.LaunchStrength)

	speed := v.Mag()
	if speed > g.MaxLaunchSpeed && speed > 0 {
		v = v.Norm().Mul(g.MaxLaunchSpeed)
	}

	return v
}

func (g *Game) SpawnPowder(pos mathx.Vec2) {
	vel := g.LaunchVelocityFromPosition(pos)

	g.World.SpawnParticle(
		g.Anchor,
		vel,
		g.SelectedMaterial,
		g.SpawnMass,
	)
}
