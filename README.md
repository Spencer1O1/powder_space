# 🌌 PowderSpace
*A gravitational sandbox inspired by Powder Game*

PowderSpace is an interactive **space particle sandbox** where users can create dust clouds, throw asteroids, form planets, ignite stars, and watch entire solar systems emerge from simple rules.

The project blends ideas from:
- Powder Game / falling-sand simulations
- N-body gravitational systems
- planetary accretion models
- emergent simulation sandboxes

Players paint matter into space and observe how gravity and material interactions produce planets, stars, and debris fields.

---

# Project Goals
## Primary Goal

Create a *real-time sandbox simulation* where simple physical rules produce emergent cosmic structures such as:
*asteroid belts
*planets
*gas giants
*stars
*debris fields
*orbital systems

## Secondary Goals

- extremely interactive
- visually satisfying
- performant with thousands or millions? of particles
- easy to extend with new materials and rules
- deterministic simulation

---

## Core Design Philosophy
### 1. Simulation First

The simulation engine drives gameplay. The world evolves continuously whether the player interacts or not.

### 2. Emergent Complexity

Small rules produce large-scale behavior.

Examples:
- dust clumps → asteroids
- asteroids collide → planets
- massive gas bodies → stars

### 3. Playable Physics

The simulation is **not intended to be astrophysically perfect**.

Instead, the physics is tuned for:

- stability
- interesting outcomes
- intuitive player interaction

---

## Key Concepts

The simulation contains **two primary entity types**.

### Particles

Particles represent **small pieces of matter**.

Examples:
- dust
- rock fragments
- gas
- ice
- plasma
- debris
- bombs 😏

Particles are numerous and lightweight.

#### Particle Properties
```go
type Particle struct {
    Position Vec2
    Velocity Vec2
    Acceleration Vec2
    Mass float64
    Material Material
}
```

Particles may merge, collide, or get absorbed by larger bodies.

---

### Bodies

Bodies represent **large gravitational objects**.

Examples:
- asteroids
- planets
- stars
- black holes

Bodies are fewer but exert strong gravitational influence.

#### Body Properties
```go
type Body struct {
    Position Vec2
    Velocity Vec2
    Mass float64
    Radius float64
    Type BodyType
    Temperature float64
}
```

Bodies grow by absorbing particles or merging with other bodies.

---

## Simulation Overview

Each simulation frame performs the following steps:
1. Apply gravitational forces
2. Integrate particle and body motion
3. Detect collisions
4. Merge/accrete matter
5. Promote objects to new body types
6. Render the updated world

---

## Physics Model
### Gravity

Gravitational force between masses:
𝐹=𝐺(𝑚1*𝑚2)/𝑟^2

To prevent extreme forces at short distances, the simulation uses softened gravity:
𝐹=𝐺(𝑚1*𝑚2)/(𝑟^2+𝜖^2)

#### Gravity Interactions
| Interaction	| Enabled |
| ----------- | ------- |
| body ↔ body |	yes |
| body ↔ particle |	yes |
| particle ↔ particle |	optional |

Particle–particle gravity may be approximated or ignored for performance, but keeping it is preferred.

---

### Integration Method

The simulation uses **semi-implicit Euler** integration.
```go
velocity += acceleration * dt
position += velocity * dt
```

Later versions may use:
- Velocity Verlet
- symplectic integration

to improve long-term orbital stability.

---

## Accretion Model

Planet formation is driven by **accretion**.

### Particle → Body

When a particle enters a body's radius:
```go
body.mass += particle.mass
// remove particle
```

### Body → Body

Body collisions produce one of three outcomes depending on impact energy AND body composition.

| Condition |	Result |
| --------- | ------ |
| low velocity | merge |
| medium velocity |	merge + eject debris |
| high velocity |	fragmentation |

---

## Body Classification

Bodies change type depending on mass and composition.

| Mass Range | Type |
| ---------- | ---- |
| small |	asteroid |
| medium | planet |
| large |	gas giant |
| very large | gas star |
| extremely large | neutron star |
| insanely large | black hole (undergoes supernova first) |

Exact thresholds are configurable.

---

## Materials

Materials define behavior and physical properties.

Example materials:

| Material | Behavior |
| -------- | -------- |
| Dust | clumps easily |
| Rock | dense solid |
| Gas |	light, star fuel |
| Ice |	low density, fragile |
| Plasma | low density, high energy |
| Bomb | explodes on impact |

Materials affect:

- density
- accretion rate
- collision results
- star formation

---

## Spatial Optimization

The simulation must handle thousands and possibly millions of objects efficiently.

Two spatial structures are used.

---

### Uniform Grid

Used for **local interactions**:
- particle merging
- collision detection
- neighborhood queries

World space is divided into fixed-size cells.

Particles only check nearby cells.

---

### Barnes–Hut Quadtree

Used for **long-range gravity**.

Distant groups of particles are approximated as a single mass.

This reduces gravity calculations from 𝑂(𝑛^2) to approximately 𝑂(𝑛log⁡𝑛)

---

## Rendering

Rendering is performed on the **GPU**.

The simulation itself runs on the CPU (for now).

The renderer displays:
- particles
- planets
- stars
- motion trails
- glow effects
- cosmic wind?

### Visual Elements

Particles:
- small glowing points

Bodies:
- colored circles

Stars:
- bright glow

Optional:
- nebula background
- atmospheric glow
- particle trails

---

## Controls
### Tools
| Tool | Action |
| ---- | ------ |
| Brush | paint dust for selected material |
| Planet seed	| spawn body (click, drag, and release for slingshot effect) |
| Cosmic wind |	impart motion |
| Eraser | remove matter |
| Play / Pause | start and stop the simulation |

---

### Camera
- pan
- zoom
- reset view

---

## Game Loop

Pseudo-code:

```
while running:

    handleInput()
    world.applyGravity()
    world.integrateMotion()
    world.resolveCollisions()
    world.performAccretion()
    world.updateBodyTypes()
    renderer.draw()
```

## Project Structure
```
powder-space/

  cmd/
    desktop-raylib/
      main.go
    web-ebiten/
      main.go

  engine/ # runtime shell
    app.go
    camera.go # put in renderer/raylib if it directly wraps raylib's camera API or draw transforms. Keep camera state/math in engine/ or game/ and keep camera application to rendering in renderer/raylib/
    input.go
    timing.go

  renderer/
    raylib/
      renderer.go
      window.go
      texture.go
      shader.go
      input.go
    ebiten/
      renderer.go
      texture.go
      shader.go
      input.go

  sim/
    world.go
    particle.go
    body.go
    gravity.go
    collision.go
    accretion.go
    integrator.go

  spatial/
    grid.go
    quadtree.go

  mathx/
    vec2.go
  gfx/
    colors.go

  game/ # Layer between UI/rendering and raw sim
    game.go
    tools.go
    materials.go
    palette.go
    rules.go

  content/ # static data tables and config
    materials.go
    bodies.go
```

Keep dependencies flowing like:
```
cmd -> engine, game, renderer
game -> sim 
sim -> spatial, mathx
```

## Milestones
### Phase 1 – Prototype
- particles
- bodies
- gravity
- merging
- rendering
- camera

### Phase 2 – Planet Formation
- debris
- body collisions
- accretion rules
- multiple materials

### Phase 3 – Stars
- gas mechanics
- star ignition
- heat effects
- stellar gravity

### Phase 4 – Destruction
- fragmentation
- explosions
- supernovae
- black holes

### Phase 5 – Large Scale
- Barnes–Hut gravity
- optimized particle systems
- thousands of particles

---

### Potential Future Features
- atmosphere simulation
- orbital tools
- spacecraft construction
- planetary rings
- nebula formation
- magnetic fields
- multiplayer sandbox
- save/load universes

---

### Inspired by **Powder Game** and **N-body astrophysics simulations**

---

### License

MIT License

---

### Final Vision

The goal of Cosmic Powder is to create a **toy universe generator** where players can experiment with gravity and matter to build their own solar systems while flying around in a ship.

From a handful of dust particles to full stellar systems, everything emerges from simple rules interacting over time.

"Paint matter into the void and watch the universe assemble itself."