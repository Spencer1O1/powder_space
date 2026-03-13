package sim

import "github.com/Spencer1O1/powder_space/v2/mathx"

type Particle struct {
	Pos   mathx.Vec2
	Vel   mathx.Vec2
	Mass  float64
	Alive bool
}
