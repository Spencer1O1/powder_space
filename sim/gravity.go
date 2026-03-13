package sim

import "github.com/Spencer1O1/powder_space/v2/mathx"

const gravSoftening = 20.0

func (w *World) particleGravAcceleration(p *Particle) mathx.Vec2 {
	acc := mathx.Vec2{}

	for i := range w.Bodies {
		b := &w.Bodies[i]

		delta := b.Pos.Sub(p.Pos)
		distSq := delta.MagSq()

		// Softened gravity
		denom := distSq + gravSoftening*gravSoftening

		if distSq > 1e-9 {
			dir := delta.Norm()
			a := w.G * b.Mass / denom
			acc = acc.Add(dir.Mul(a))
		}
	}

	return acc
}
