package sim

import (
	"runtime"
	"sync"

	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	"github.com/Spencer1O1/powder_space/v2/spatial"
)

type World struct {
	Particles []Particle
	Quadtree  spatial.Quadtree[Particle]

	// Smaller = more accurate simulation.
	// Determines when a whole node can be treated as a combined mass.
	Theta float32

	// Larger = more softening
	// Prevents gravity from becoming absurdly huge.
	Epsilon float32

	// Gravitational constant
	G float32
}

func NewWorld() *World {
	const theta = 0.7
	const epsilon = 2.0
	const G = 1000.0

	qt := spatial.NewQuadtree[Particle](theta, epsilon, G)

	return &World{
		Particles: make([]Particle, 0, 1024),
		Quadtree:  *qt,
		Theta:     theta,
		Epsilon:   epsilon,
		G:         G,
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
	w.Iterate(dt)
	w.Collide(dt)

	w.compactParticles()

	w.Attract(dt)
}

func (w *World) Iterate(dt float32) {
	for i := range w.Particles {
		if !w.Particles[i].Alive {
			continue
		}
		w.Particles[i].Integrate(dt)
	}
}

func (w *World) Collide(dt float32) {
	// TODO: implement particle collision for clumping and body formation
}

func (w *World) Attract(dt float32) {
	if len(w.Particles) == 0 {
		return
	}

	// Build tree single-threaded
	w.Quadtree.Build(w.Particles)

	// Parallelize
	workers := runtime.GOMAXPROCS(0)
	if workers > len(w.Particles) {
		workers = len(w.Particles)
	}
	if workers < 1 {
		workers = 1
	}

	n := len(w.Particles)
	chunk := (n + workers - 1) / workers

	var wg sync.WaitGroup
	wg.Add(workers)

	for worker := 0; worker < workers; worker++ {
		start := worker * chunk
		end := start + chunk
		if end > n {
			end = n
		}

		go func(start, end int) {
			defer wg.Done()

			for i := start; i < end; i++ {
				w.Particles[i].Acc = w.Quadtree.Acc(w.Particles[i].Pos)
			}
		}(start, end)
	}

	wg.Wait()
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
