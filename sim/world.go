package sim

import "github.com/Spencer1O1/powder_space/v2/mathx"

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
		Pos:   pos,
		Vel:   vel,
		Mass:  1,
		Alive: true,
	})
}

func (w *World) Step(dt float64) {
	for i := range w.Particles {
		p := &w.Particles[i]
		if !p.Alive {
			continue
		}

		acc := mathx.Vec2{}

		for j := range w.Bodies {
			b := &w.Bodies[j]

			delta := b.Pos.Sub(p.Pos)
			distSq := delta.MagSq()
			if distSq < 25 {
				distSq = 25
			}

			dir := delta.Norm()
			a := w.G * b.Mass / distSq
			acc = acc.Add(dir.Mul(a))
		}

		// semi-implicit Euler
		p.Vel = p.Vel.Add(acc.Mul(dt))
		p.Pos = p.Pos.Add(p.Vel.Mul(dt))

		// absorb after movement
		for j := range w.Bodies {
			b := &w.Bodies[j]

			delta := b.Pos.Sub(p.Pos)
			dist := delta.Mag()

			if dist <= b.Radius {
				b.Mass += p.Mass
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
}
