package sim

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/mathx"
)

type World struct {
	Particles []Particle
	Bodies    []Body
	G         float64
}

func NewWorld() *World {
	return &World{
		G: 4000,
		Bodies: []Body{
			{
				Pos:    mathx.V(640, 360),
				Mass:   5000,
				Radius: 30,
			},
		},
	}
}

func (w *World) SpawnParticle(pos mathx.Vec2, vel mathx.Vec2) {
	w.Particles = append(w.Particles, Particle{
		Radius: 1,
		Pos:    pos,
		Vel:    vel,
		Mass:   10,
		Alive:  true,
	})
}

func (w *World) Step(dt float64) {
	const gravSoftening = 20.0

	for i := range w.Particles {
		p := &w.Particles[i]
		if !p.Alive {
			continue
		}

		lastPos := p.Pos
		acc := mathx.Vec2{}

		for j := range w.Bodies {
			b := &w.Bodies[j]

			delta := b.Pos.Sub(p.Pos)
			distSq := delta.MagSq()

			// softened gravity
			denom := distSq + gravSoftening*gravSoftening

			if delta.MagSq() > 1e-9 {
				dir := delta.Norm()
				a := w.G * b.Mass / denom
				acc = acc.Add(dir.Mul(a))
			}
		}

		// semi-implicit Euler
		p.Vel = p.Vel.Add(acc.Mul(dt))
		p.Pos = p.Pos.Add(p.Vel.Mul(dt))

		// absorb after movement
		for j := range w.Bodies {
			b := &w.Bodies[j]
			absorbRadius := b.Radius + p.Radius + 2
			if mathx.PointInCircle(lastPos, b.Pos, absorbRadius) ||
				mathx.PointInCircle(p.Pos, b.Pos, absorbRadius) ||
				mathx.SegmentIntersectsCircle(lastPos, p.Pos, b.Pos, absorbRadius) {

				lastMass := b.Mass
				b.Mass += p.Mass

				// Body scales as if it were a sphere
				b.Radius = b.Radius * math.Cbrt(b.Mass/lastMass)

				p.Alive = false
				break
			}
		}
	}

	w.compactParticles()
}

func (w *World) compactParticles() {
	dst := w.Particles[:0]
	for _, p := range w.Particles {
		if p.Alive {
			dst = append(dst, p)
		}
	}
	w.Particles = dst
}
