package sim

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/mathx"
)

type ParticleInteractionConfig struct {
	BaseAttractionStrength float64
	CollisionRestitution   float64
	PairwiseDamping        float64
	OverlapRatio           float64
	BaseInfluenceRadius    float64
	RadiusInfluenceCoeff   float64
}

var config = ParticleInteractionConfig{
	// Base attraction scale for nearby particles.
	// Low values = weak or no clumping.
	// High values = particles snap together violently.
	BaseAttractionStrength: 10.0,

	// Collision restitution along the collision normal.
	// 0 = perfectly inelastic/sticky, 1 = perfectly bouncy.
	CollisionRestitution: 0.0,

	// Damping applied to relative velocity for nearby particles.
	// Helps clumps settle instead of jittering forever.
	PairwiseDamping: 1.0, // (scaled by dt)

	// Allows particles to visually overlap a bit when together.
	// Makes clumps look more compact so they do not visually shrink as much
	// when later converted into bodies.
	OverlapRatio: 0.6,

	// Minimum particle influence radius for local clumping
	BaseInfluenceRadius: 2.0,

	// Scalar to increase / decrease the particle influence radius based on the particle's mass
	// 0 = no additional influence radius (only base)
	RadiusInfluenceCoeff: 1.8,
}

// handles all pairwise particle-particle interactions for the physics step
// Naive O(n^2) all-pairs loop
// Later, can be accelerated using a uniform spatial grid.
func (w *World) resolveParticleInteractions(dt float64) {
	n := len(w.Particles)

	for i := 0; i < n; i++ {
		a := &w.Particles[i]
		if !a.Alive {
			continue
		}

		for j := i + 1; j < n; j++ {
			b := &w.Particles[j]
			if !b.Alive {
				continue
			}

			w.resolveParticlePair(a, b, dt)
		}
	}
}

// handles the interaction between a single particle pair.
//  1. if particles overlap, resolve collision/separation and damp normal force
//  2. otherwise, if particles are within clump range, apply attraction
//     and pairwise relative damping so they can settle into clusters.
func (w *World) resolveParticlePair(a, b *Particle, dt float64) {
	delta, dist := particleSeparation(a, b)
	normal := delta.Mul(1.0 / dist)

	restDist := particleRestDistance(a, b)
	clumpRadius := particleClumpRadius(a, b, restDist)

	// 1) overlap -> collision/separation
	if dist < restDist {
		resolveParticleOverlap(a, b, normal, dist, restDist)
		return
	}

	// 2) nearby -> attraction
	if dist < clumpRadius {
		// stronger when closer to rest distance, weaker toward edge
		t := particleClumpInfluence(dist, restDist, clumpRadius)

		applyParticleAttraction(a, b, normal, dist, dt)
		applyParticlePairwiseDamping(a, b, t, dt)
	}
}

// particleSeparation returns the separation vector and distance between two particles.
//
// A tiny deterministic nudge is used if the particles are exactly coincident
// to avoid zero-length normal vectors and divide-by-zero problems.
func particleSeparation(a, b *Particle) (mathx.Vec2, float64) {
	delta := b.Pos.Sub(a.Pos)
	distSq := delta.MagSq()

	if distSq < 1e-12 {
		delta = mathx.V(0.001, 0)
	}

	return delta, delta.Mag()
}

// particleRestDistance returns the effective "resting contact distance"
// between two particles.
//
// This is smaller than the full sum of radii so that clumps look visually
// tighter and more compact.
func particleRestDistance(a, b *Particle) float64 {
	return (a.Radius + b.Radius) * (1.0 - config.OverlapRatio)
}

// particleInfluenceRadius returns how far a particle's local clumping influence extends.
//
// Heavier particles influence a larger neighborhood, using cube-root scaling
// so the influence grows sublinearly with mass.
func particleInfluenceRadius(p *Particle) float64 {
	return config.BaseInfluenceRadius * config.RadiusInfluenceCoeff * math.Cbrt(p.Mass)
}

// particleClumpRadius returns the outer radius where two particles can still
// weakly attract each other into a clump.
func particleClumpRadius(a, b *Particle, restDist float64) float64 {
	return restDist + particleInfluenceRadius(a) + particleInfluenceRadius(b)
}

// particleClumpInfluence returns a normalized falloff factor in [0, 1]
// describing how strongly two nearby particles should interact.
//
// 1 = strongest near rest distance
// 0 = no effect at the edge of the clump radius
func particleClumpInfluence(dist, restDist, clumpRadius float64) float64 {
	t := 1.0 - (dist-restDist)/(clumpRadius-restDist)
	if t < 0 {
		return 0
	}
	if t > 1 {
		return 1
	}
	return t
}

// resolveParticleOverlap separates overlapping particles and damps their
// closing normal velocity.
//
// This acts like a sticky/inelastic collision response:
// - particles are pushed apart so they no longer overlap
// - bounce along the collision normal is reduced based on restitution
func resolveParticleOverlap(a, b *Particle, normal mathx.Vec2, dist, restDist float64) {
	penetration := restDist - dist

	totalMass := a.Mass + b.Mass
	if totalMass <= 0 {
		totalMass = 1
	}

	// Lighter particles move more during position correction.
	aMove := penetration * (b.Mass / totalMass)
	bMove := penetration * (a.Mass / totalMass)

	a.Pos = a.Pos.Sub(normal.Mul(aMove))
	b.Pos = b.Pos.Add(normal.Mul(bMove))

	relVel := b.Vel.Sub(a.Vel)
	relNormalSpeed := relVel.Dot(normal)

	// Only resolve if particles are moving toward each other.
	if relNormalSpeed < 0 {
		impulseMag := -(1.0 - config.CollisionRestitution) * relNormalSpeed

		aImpulse := impulseMag * (b.Mass / totalMass)
		bImpulse := impulseMag * (a.Mass / totalMass)

		a.Vel = a.Vel.Sub(normal.Mul(aImpulse))
		b.Vel = b.Vel.Add(normal.Mul(bImpulse))
	}
}

// applyParticleAttraction applies a gentle local attractive interaction
// between two nearby particles.
//
// This is not full gravity. It is a local clumping force used to help
// particles gather into visible clusters.
//
// In the real world space dust clumps due to static electricity, not only gravity
func applyParticleAttraction(a, b *Particle, normal mathx.Vec2, influence, dt float64) {
	pairStrength := config.BaseAttractionStrength * math.Sqrt(a.Mass*b.Mass)
	forceMag := pairStrength * influence

	accelA := forceMag / a.Mass
	accelB := forceMag / b.Mass

	a.Vel = a.Vel.Add(normal.Mul(accelA * dt))
	b.Vel = b.Vel.Sub(normal.Mul(accelB * dt))
}

// applyParticlePairwiseDamping reduces relative velocity between nearby particles
// while preserving the pair's center-of-mass motion as much as possible.
//
// This is the key to stable clumping:
// particles in a forming cluster stop jittering against each other, but the
// whole clump can still orbit or drift together.
func applyParticlePairwiseDamping(a, b *Particle, influence, dt float64) {
	relVel := b.Vel.Sub(a.Vel)

	totalMass := a.Mass + b.Mass
	if totalMass <= 0 {
		totalMass = 1
	}

	dampFactor := config.PairwiseDamping * influence * dt
	if dampFactor > 1.0 {
		dampFactor = 1.0
	}

	deltaVel := relVel.Mul(dampFactor)

	aShare := b.Mass / totalMass
	bShare := a.Mass / totalMass

	a.Vel = a.Vel.Add(deltaVel.Mul(aShare))
	b.Vel = b.Vel.Sub(deltaVel.Mul(bShare))
}
