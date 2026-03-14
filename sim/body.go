package sim

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
)

type Body struct {
	Composition content.CompositionMap
	Pos         mathx.Vec2
	Vel         mathx.Vec2

	// Derived
	Mass   float64
	Radius float64
}

func createBody(
	initialComposition content.CompositionMap,
	pos, vel mathx.Vec2,
) Body {
	b := Body{
		Pos:         pos,
		Vel:         vel,
		Composition: initialComposition,
	}

	b.RecomputeDerived()
	return b
}

func (b *Body) RecomputeDerived() {
	mass, radius, _, _ := b.Composition.GetSphericalDerivedValues()

	b.Mass = mass
	b.Radius = radius
}
