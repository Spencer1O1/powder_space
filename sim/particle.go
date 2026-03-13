package sim

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
)

type Particle struct {
	Material content.MaterialID

	Mass   float64
	Radius float64

	Pos   mathx.Vec2
	Vel   mathx.Vec2
	Alive bool
}

func radiusFromMassAndDensity(mass float64, density float64) float64 {
	volume := mass / density

	r := math.Max(math.Cbrt((3*volume)/(4*math.Pi)), 0.5)

	return r
}
