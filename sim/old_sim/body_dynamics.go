package old_sim

// const bodyGravSoftening = 10.0

// // stepBodies advances body-body gravity, integration, collisions, and aging.

// func (w *World) stepBodies(dt float32) {
// 	if len(w.Bodies) == 0 {
// 		return
// 	}

// 	accels := make([]mathx.Vec2, len(w.Bodies))

// 	w.computeBodyAccelerations(accels)
// 	w.integrateBodies(accels, dt)
// 	w.resolveBodyCollisions()
// 	w.compactBodies()
// 	w.ageBodies(dt)
// }

// // computeBodyAccelerations computes pairwise body-body gravitational acceleration.
// // Equal and opposite effects are applied to each pair.
// func (w *World) computeBodyAccelerations(accels []mathx.Vec2) {
// 	n := len(w.Bodies)

// 	for i := 0; i < n; i++ {
// 		if !w.Bodies[i].Alive {
// 			continue
// 		}

// 		for j := i + 1; j < n; j++ {
// 			if !w.Bodies[j].Alive {
// 				continue
// 			}

// 			a := &w.Bodies[i]
// 			b := &w.Bodies[j]

// 			delta := b.Pos.Sub(a.Pos)
// 			distSq := delta.MagSq()
// 			if distSq < 1e-12 {
// 				continue
// 			}

// 			denom := distSq + bodyGravSoftening*bodyGravSoftening
// 			dist := mathx.Sqrt(distSq)
// 			normal := delta.Mul(1.0 / dist)

// 			// Force magnitude divided by mass gives acceleration.
// 			accelOnA := w.G * b.Mass / denom
// 			accelOnB := w.G * a.Mass / denom

// 			accels[i] = accels[i].Add(normal.Mul(accelOnA))
// 			accels[j] = accels[j].Sub(normal.Mul(accelOnB))
// 		}
// 	}
// }

// // integrateBodies advances body positions and velocities using semi-implicit Euler.
// func (w *World) integrateBodies(accels []mathx.Vec2, dt float32) {
// 	for i := range w.Bodies {
// 		b := &w.Bodies[i]
// 		if !b.Alive {
// 			continue
// 		}

// 		b.Vel = b.Vel.Add(accels[i].Mul(dt))
// 		b.Pos = b.Pos.Add(b.Vel.Mul(dt))
// 	}
// }

// func (w *World) ageBodies(dt float32) {
// 	for i := range w.Bodies {
// 		if !w.Bodies[i].Alive {
// 			continue
// 		}
// 		w.Bodies[i].Age += dt
// 	}
// }
