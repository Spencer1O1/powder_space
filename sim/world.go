package sim

import (
	"github.com/Spencer1O1/powder_space/v2/content"
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

func (w *World) SpawnParticle(pos mathx.Vec2, vel mathx.Vec2, material content.MaterialID, mass float64) {
	mat := content.Materials[material]

	radius := radiusFromMassAndDensity(mass, mat.Density)

	w.Particles = append(w.Particles, Particle{
		Material: material,
		Mass:     mass,
		Radius:   radius,
		Pos:      pos,
		Vel:      vel,
		Alive:    true,
	})
}

func (w *World) Step(dt float64) {
	for i := range w.Particles {
		p := &w.Particles[i]
		if !p.Alive {
			continue
		}
		lastPos := p.Pos

		acc := w.particleGravAcceleration(p)
		w.integrateParticle(p, acc, dt)

		if w.tryAbsorbParticle(p, lastPos) {
			continue
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

func (w *World) ClearParticles() {
	w.Particles = w.Particles[:0]
}
