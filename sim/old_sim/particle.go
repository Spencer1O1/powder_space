package old_sim

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/mathx/geo"
)

type Particle struct {
	MaterialId content.MaterialID
	M          float32
	Pos        mathx.Vec2
	Vel        mathx.Vec2
	Acc        mathx.Vec2
	Alive      bool

	// Derived
	Radius          float32
	InvM            float32
	InfluenceRadius float32
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

	if p.M > 0 {
		p.InvM = 1.0 / p.M
	} else {
		p.InvM = 0
	}

	p.recomputeInfluenceRadius()
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

// Determines how far a particle's local clumping influence extends and caches
// the value in the Particle struct for use by particle interactions
//
// Heavier particles influence a larger neighborhood, using cube-root scaling
// so the influence grows sublinearly with mass.
func (p *Particle) recomputeInfluenceRadius() {
	p.InfluenceRadius = config.BaseInfluenceRadius * config.RadiusInfluenceCoeff * float32(math.Cbrt(float64(p.M)))
}
