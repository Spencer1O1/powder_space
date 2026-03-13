package sim

import (
	"math"

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
			absorbBodyMass(b, p.Mass)
			p.Alive = false
			return true
		}
	}

	return false
}

func absorbBodyMass(b *Body, addedMass float64) {
	lastMass := b.Mass
	b.Mass += addedMass
	b.Radius *= math.Cbrt(b.Mass / lastMass)
}
