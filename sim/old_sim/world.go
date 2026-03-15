package old_sim

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/spatial"
)

type World struct {
	Particles []Particle
	Bodies    []Body
	G         float32

	ParticleGrid *spatial.UniformGrid
}

func NewWorld() *World {
	return &World{
		G: 200.0,
		Bodies: []Body{
			createBody(content.CompositionMap{
				content.MaterialDust: 400_000.0,
			}, mathx.V(800, 540), mathx.V(0, 400)),
			createBody(content.CompositionMap{
				content.MaterialDust: 400_000.0,
			}, mathx.V(1120, 540), mathx.V(0, -400)),
		},
		ParticleGrid: spatial.NewUniformGrid(30.0),
	}
}

func (w *World) SpawnParticle(pos mathx.Vec2, vel mathx.Vec2, material content.MaterialID, mass float32) {
	w.Particles = append(w.Particles, createParticle(
		material,
		mass,
		pos,
		vel,
		mathx.V0(),
		true,
	))
}

func (w *World) Step(dt float32) {
	// w.stepBodies(dt)

	// for i := range w.Particles {
	// 	p := &w.Particles[i]
	// 	if !p.Alive {
	// 		continue
	// 	}
	// 	lastPos := p.Pos

	// 	acc := w.particleBodyGravAcceleration(p)
	// 	w.integrateParticle(p, acc, dt)

	// 	if w.tryAbsorbParticle(p, lastPos) {
	// 		continue
	// 	}
	// }

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
