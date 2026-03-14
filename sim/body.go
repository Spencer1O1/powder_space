package sim

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
)

type Body struct {
	Pos    mathx.Vec2
	Vel    mathx.Vec2
	Mass   float64
	Radius float64

	Composition map[content.MaterialID]float64
}
