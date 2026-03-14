package sim

import (
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
