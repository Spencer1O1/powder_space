package old_sim

const radiusPadding = 2

// func (w *sim.World) tryAbsorbParticle(p *Particle, lastPos mathx.Vec2) bool {
// 	for i := range w.Bodies {
// 		b := &w.Bodies[i]
// 		absorbRadius := b.Radius + p.Radius + radiusPadding

// 		if geo.PointInCircle(lastPos, b.Pos, absorbRadius) ||
// 			geo.PointInCircle(p.Pos, b.Pos, absorbRadius) ||
// 			geo.SegmentIntersectsCircle(lastPos, p.Pos, b.Pos, absorbRadius) {
// 			absorbParticleIntoBody(b, p)
// 			p.Alive = false
// 			return true
// 		}
// 	}

// 	return false
// }

// func absorbParticleIntoBody(b *Body, p *Particle) {
// 	if b.Composition == nil {
// 		b.Composition = content.CompositionMap{}
// 	}
// 	b.Composition[p.MaterialId] += p.M
// 	b.RecomputeDerived()
// }
