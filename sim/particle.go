package sim

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/mathx/geo"
)

type Particle struct {
	// "HOT" values (used by barnes-hut)
	M   float32
	Pos mathx.Vec2
	Vel mathx.Vec2
	Acc mathx.Vec2

	MaterialId content.MaterialID
	Alive      bool

	// Derived
	Radius float32
}

func (p Particle) Position() mathx.Vec2 { return p.Pos }
func (p Particle) Mass() float32        { return p.M }
func (p *Particle) Integrate(dt float32) {
	// semi-implicit euler
	p.Vel = p.Vel.Add(p.Acc.Mul(dt))
	p.Pos = p.Pos.Add(p.Vel.Mul(dt))
}

func (p *Particle) RecomputeDerived() {
	mat := content.Materials[p.MaterialId]
	volume := p.M / mat.Density
	p.Radius = geo.SphericalRadiusFromVolume(volume)
}

func createParticle(
	material content.MaterialID,
	mass float32,
	pos, vel, acc mathx.Vec2,
	alive bool,
) Particle {
	newParticle := Particle{
		MaterialId: material,
		M:          mass,
		Pos:        pos,
		Vel:        vel,
		Acc:        acc,
		Alive:      alive,
	}
	newParticle.RecomputeDerived()

	return newParticle
}
