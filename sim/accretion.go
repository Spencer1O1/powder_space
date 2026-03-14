package sim

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/mathx/geo"
)

const radiusPadding = 2

func (w *World) tryAbsorbParticle(p *Particle, lastPos mathx.Vec2) bool {
	for i := range w.Bodies {
		b := &w.Bodies[i]
		absorbRadius := b.Radius + p.Radius + radiusPadding

		if geo.PointInCircle(lastPos, b.Pos, absorbRadius) ||
			geo.PointInCircle(p.Pos, b.Pos, absorbRadius) ||
			geo.SegmentIntersectsCircle(lastPos, p.Pos, b.Pos, absorbRadius) {
			absorbParticleIntoBody(b, p)
			p.Alive = false
			return true
		}
	}

	return false
}

func absorbParticleIntoBody(b *Body, p *Particle) {
	if b.Composition == nil {
		b.Composition = map[content.MaterialID]float64{}
	}

	b.Composition[p.Material] += p.Mass
	b.Mass += p.Mass

	_, radius, _, _ := getMaterialDerivedValues(b.Composition)
	b.Radius = radius
}

func getMaterialDerivedValues(composition map[content.MaterialID]float64) (
	mass,
	radius,
	volume,
	density float64,
) {
	totalMass := 0.0
	totalVolume := 0.0

	for materialID, materialMass := range composition {
		if materialMass <= 0 {
			continue
		}

		mat := content.Materials[materialID]
		density := mat.Density

		totalMass += materialMass
		totalVolume += materialMass / density
	}

	if totalVolume <= 0 {
		return 0.0, 0.0, 0.0, 1.0
	}

	d := totalMass / totalVolume
	r := sphericalRadiusFromVolume(totalVolume)

	return totalMass, r, totalVolume, d
}

func sphericalRadiusFromVolume(v float64) float64 {
	return math.Max(math.Cbrt((3*v)/(4*math.Pi)), 0.5)
}

func sphericalRadiusFromMassAndDensity(mass float64, density float64) float64 {
	volume := mass / density
	return math.Max(math.Cbrt((3*volume)/(4*math.Pi)), 0.5)
}
