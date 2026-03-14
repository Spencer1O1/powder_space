package sim

import "github.com/Spencer1O1/powder_space/v2/content"

// resolveBodyCollisions merges overlapping bodies.
// First version: simple merge-on-overlap.
func (w *World) resolveBodyCollisions() {
	n := len(w.Bodies)

	for i := 0; i < n; i++ {
		a := &w.Bodies[i]
		if !a.Alive {
			continue
		}

		for j := i + 1; j < n; j++ {
			b := &w.Bodies[j]
			if !b.Alive {
				continue
			}

			delta := b.Pos.Sub(a.Pos)
			r := a.Radius + b.Radius

			if delta.MagSq() <= r*r {
				mergeBodies(a, b)
				b.Alive = false
			}
		}
	}
}

func (w *World) compactBodies() {
	dst := w.Bodies[:0]
	for _, b := range w.Bodies {
		if b.Alive {
			dst = append(dst, b)
		}
	}
	w.Bodies = dst
}

// mergeBodies merges b into a, conserving momentum and combining composition.
func mergeBodies(a, b *Body) {
	totalMass := a.Mass + b.Mass
	if totalMass <= 0 {
		return
	}

	// Center of mass position.
	newPos := a.Pos.Mul(a.Mass).Add(b.Pos.Mul(b.Mass)).Mul(1.0 / totalMass)

	// Conserve momentum.
	newVel := a.Vel.Mul(a.Mass).Add(b.Vel.Mul(b.Mass)).Mul(1.0 / totalMass)

	a.Pos = newPos
	a.Vel = newVel
	mergeCompositionInto(a, b)

	a.RecomputeDerived()

	// Proto beats mature only if you want that behavior.
	// For now, once merged, keep proto only if both were proto. (later proto is based on mass)
	if a.Phase == BodyPhaseMature || b.Phase == BodyPhaseMature {
		a.Phase = BodyPhaseMature
	} else {
		a.Phase = BodyPhaseProto
	}

	a.Age = 0
}

func mergeCompositionInto(dst, src *Body) {
	if dst.Composition == nil {
		dst.Composition = map[content.MaterialID]float64{}
	}

	for materialID, mass := range src.Composition {
		dst.Composition[materialID] += mass
	}
}
