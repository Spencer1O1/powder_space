package spatial

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/mathx"
)

type MassPoint interface {
	Position() mathx.Vec2
	Mass() float32
}

// ---------------------------------------------------------
// Quad
// ---------------------------------------------------------

type Quad struct {
	Center mathx.Vec2
	Size   float32
}

func NewQuadContaining[T MassPoint](pts []T) Quad {
	// Handle empty input defensively
	if len(pts) == 0 {
		return Quad{Center: mathx.V0(), Size: 1}
	}

	minX, minY := float32(math.MaxFloat32), float32(math.MaxFloat32)
	maxX, maxY := -minX, -minY

	for _, p := range pts {
		pos := p.Position()
		minX = mathx.Min(minX, pos.X)
		minY = mathx.Min(minY, pos.Y)
		maxX = mathx.Max(maxX, pos.X)
		maxY = mathx.Max(maxY, pos.Y)
	}

	center := mathx.Vec2{
		X: (minX + maxX) * 0.5,
		Y: (minY + maxY) * 0.5,
	}
	size := mathx.Max(maxX-minX, maxY-minY)

	// (Optional) if size == 0, keep it 0; duplicate points will just sum mass at a leaf.
	return Quad{Center: center, Size: size}
}

// rust ((pos.y > center.y) as usize) << 1 | (pos.x > center.x) as usize
func (q Quad) FindQuadrant(pos mathx.Vec2) int {
	ybit := 0
	if pos.Y > q.Center.Y {
		ybit = 1
	}
	xbit := 0
	if pos.X > q.Center.X {
		xbit = 1
	}
	return (ybit << 1) | xbit
}

func (q Quad) IntoQuadrant(quadrant int) Quad {
	q.Size *= 0.5
	q.Center.X += (float32(quadrant&1) - 0.5) * q.Size
	q.Center.Y += (float32(quadrant>>1) - 0.5) * q.Size
	return q
}

func (q Quad) Subdivide() [4]Quad {
	return [4]Quad{
		q.IntoQuadrant(0),
		q.IntoQuadrant(1),
		q.IntoQuadrant(2),
		q.IntoQuadrant(3),
	}
}

// ---------------------------------------------------------
// Node
// ---------------------------------------------------------

type Node struct {
	Children int // 0 => leaf; otherwise index of first child
	Next     int
	Pos      mathx.Vec2
	Mass     float32
	Quad     Quad
}

func NewNode(next int, quad Quad) Node {
	return Node{
		Children: 0,
		Next:     next,
		Pos:      mathx.V0(),
		Mass:     0,
		Quad:     quad,
	}
}

func (n Node) IsLeaf() bool   { return n.Children == 0 }
func (n Node) IsBranch() bool { return n.Children != 0 }
func (n Node) IsEmpty() bool  { return n.Mass == 0 }

// ---------------------------------------------------------
// Quadtree (Barnes-Hut style) — optimized build (no Propagate)
// ---------------------------------------------------------

type Quadtree[T MassPoint] struct {
	TSq   float32
	ESq   float32
	G     float32
	Nodes []Node
}

const Root = 0

// epsilon: softening. prevents gravity from becoming absurdly huge. Larger = more softening
func NewQuadtree[T MassPoint](theta, epsilon, G float32) *Quadtree[T] {
	return &Quadtree[T]{
		TSq: theta * theta,
		ESq: epsilon * epsilon,
		G:   G,
	}
}

func (qt *Quadtree[T]) Clear(quad Quad) {
	qt.Nodes = qt.Nodes[:0]
	qt.Nodes = append(qt.Nodes, NewNode(0, quad))
}

func (qt *Quadtree[T]) Build(pts []T) {
	qt.Clear(NewQuadContaining(pts))
	for _, p := range pts {
		qt.Insert(p.Position(), p.Mass()) // Insert maintains COM incrementally
	}
}

// subdivide(node) -> children index
func (qt *Quadtree[T]) subdivide(node int) int {
	children := len(qt.Nodes)
	qt.Nodes[node].Children = children

	nexts := [4]int{
		children + 1,
		children + 2,
		children + 3,
		qt.Nodes[node].Next,
	}
	quads := qt.Nodes[node].Quad.Subdivide()

	for i := 0; i < 4; i++ {
		qt.Nodes = append(qt.Nodes, NewNode(nexts[i], quads[i]))
	}
	return children
}

// Update center-of-mass and mass: newCOM = (oldCOM*oldM + pos*m) / (oldM+m)
func accumulate(n *Node, pos mathx.Vec2, m float32) {
	if n.Mass == 0 {
		n.Pos = pos
		n.Mass = m
		return
	}
	oldM := n.Mass
	newM := oldM + m
	n.Pos = n.Pos.Mul(oldM).Add(pos.Mul(m)).Div(newM)
	n.Mass = newM
}

// Insert stores points in leaves; branches store COM. No second pass needed.
func (qt *Quadtree[T]) Insert(pos mathx.Vec2, mass float32) {
	node := Root

	for {
		n := &qt.Nodes[node]

		// Branch: accumulate and descend
		if n.IsBranch() {
			accumulate(n, pos, mass)
			q := n.Quad.FindQuadrant(pos)
			node = n.Children + q
			continue
		}

		// Leaf cases
		if n.IsEmpty() {
			// empty leaf: store point
			n.Pos = pos
			n.Mass = mass
			return
		}

		// occupied leaf: same position => just add mass
		if pos.Equal(n.Pos) {
			n.Mass += mass
			return
		}

		// occupied leaf with different position => split
		oldPos, oldMass := n.Pos, n.Mass

		children := qt.subdivide(node)

		// This node becomes a branch representing {old, new}
		n.Pos = mathx.V0()
		n.Mass = 0
		accumulate(n, oldPos, oldMass)
		accumulate(n, pos, mass)

		// Reinsert the old point into children (without touching ancestors above this node)
		qt.insertFrom(children+n.Quad.FindQuadrant(oldPos), oldPos, oldMass)
		qt.insertFrom(children+n.Quad.FindQuadrant(pos), pos, mass)
		return
	}
}

// insertFrom inserts into subtree rooted at `node`, updating COM/mass along that subtree only.
func (qt *Quadtree[T]) insertFrom(node int, pos mathx.Vec2, mass float32) {
	for {
		n := &qt.Nodes[node]

		if n.IsBranch() {
			accumulate(n, pos, mass)
			q := n.Quad.FindQuadrant(pos)
			node = n.Children + q
			continue
		}

		if n.IsEmpty() {
			n.Pos = pos
			n.Mass = mass
			return
		}

		if pos.Equal(n.Pos) {
			n.Mass += mass
			return
		}

		// split leaf
		oldPos, oldMass := n.Pos, n.Mass
		children := qt.subdivide(node)

		n.Pos = mathx.V0()
		n.Mass = 0
		accumulate(n, oldPos, oldMass)
		accumulate(n, pos, mass)

		qt.insertFrom(children+n.Quad.FindQuadrant(oldPos), oldPos, oldMass)
		qt.insertFrom(children+n.Quad.FindQuadrant(pos), pos, mass)
		return
	}
}

func (qt *Quadtree[T]) Acc(pos mathx.Vec2) mathx.Vec2 {
	acc := mathx.V0()

	node := Root
	for {
		n := qt.Nodes[node]

		d := n.Pos.Sub(pos)
		dSq := d.MagSq()

		if n.IsLeaf() || (n.Quad.Size*n.Quad.Size) < (dSq*qt.TSq) {
			denom := (dSq + qt.ESq) * mathx.Sqrt(dSq)
			if denom != 0 {
				acc = acc.Add(d.Mul(qt.G * n.Mass / denom))
			}
			if n.Next == 0 {
				break
			}
			node = n.Next
		} else {
			node = n.Children
		}
	}

	return acc
}
