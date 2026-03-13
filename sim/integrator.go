package sim

import "github.com/Spencer1O1/powder_space/v2/mathx"

// semi-implicit Euler
func (w *World) integrateParticle(p *Particle, acc mathx.Vec2, dt float64) {
	p.Vel = p.Vel.Add(acc.Mul(dt))
	p.Pos = p.Pos.Add(p.Vel.Mul(dt))
}
