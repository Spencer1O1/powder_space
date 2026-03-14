package sim

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/mathx/geo"
)

type Particle struct {
	Material content.MaterialID
	Mass     float64
	Pos      mathx.Vec2
	Vel      mathx.Vec2
	Alive    bool

	// Derived
	Radius          float64
	InvMass         float64
	InfluenceRadius float64
}

func (p *Particle) RecomputeDerived() {
	mat := content.Materials[p.Material]
	volume := p.Mass / mat.Density
	p.Radius = geo.SphericalRadiusFromVolume(volume)

	if p.Mass > 0 {
		p.InvMass = 1.0 / p.Mass
	} else {
		p.InvMass = 0
	}

	p.recomputeInfluenceRadius()
}

func createParticle(
	material content.MaterialID,
	mass float64,
	pos, vel mathx.Vec2,
	alive bool,
) Particle {
	newParticle := Particle{
		Material: material,
		Mass:     mass,
		Pos:      pos,
		Vel:      vel,
		Alive:    alive,
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
	p.InfluenceRadius = config.BaseInfluenceRadius * config.RadiusInfluenceCoeff * math.Cbrt(p.Mass)
}
