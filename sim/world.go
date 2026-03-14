package sim

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/spatial"
)

type World struct {
	Particles []Particle
	Bodies    []Body
	G         float64

	ParticleGrid *spatial.UniformGrid
}

func NewWorld() *World {
	initialBodyMass := 400000.0
	initialBodyComposition := map[content.MaterialID]float64{
		content.MaterialDust: initialBodyMass,
	}
	initialBodyRadius := radiusFromMassAndDensity(
		initialBodyMass,
		weightedDensity(initialBodyComposition),
	)

	return &World{
		G: 200.0,
		Bodies: []Body{
			{
				Pos:         mathx.V(960, 540),
				Vel:         mathx.Vec2{},
				Mass:        initialBodyMass,
				Radius:      initialBodyRadius,
				Composition: initialBodyComposition,
			},
		},
		ParticleGrid: spatial.NewUniformGrid(30.0),
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

		acc := w.particleBodyGravAcceleration(p)
		w.integrateParticle(p, acc, dt)

		if w.tryAbsorbParticle(p, lastPos) {
			continue
		}
	}

	w.resolveParticleInteractions(dt)
	w.compactParticles()
}

func (w *World) rebuildParticleGrid() {
	w.ParticleGrid.Clear()

	for i := range w.Particles {
		p := &w.Particles[i]
		if !p.Alive {
			continue
		}
		w.ParticleGrid.Insert(i, p.Pos)
	}
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
