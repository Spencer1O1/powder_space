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
// Bounds
// ---------------------------------------------------------

type Bounds struct {
	Center mathx.Vec2
	Size   float32
}

func NewBoundsContaining[T MassPoint](pts []T) Bounds {
	// Handle empty input defensively
	if len(pts) == 0 {
		return Bounds{Center: mathx.V0(), Size: 1}
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
	return Bounds{Center: center, Size: size}
}

// ---------------------------------------------------------
// Node
// ---------------------------------------------------------

type Node struct {
	Children int // 0 => leaf; otherwise index of first child
	Next     int

	Pos  mathx.Vec2
	Mass float32

	// Quad     Quad
	Center mathx.Vec2
	Size   float32
	SizeSq float32
}

func NewNode(next int, bounds Bounds) Node {
	return Node{
		Children: 0,
		Next:     next,
		Pos:      mathx.V0(),
		Mass:     0,
		Center:   bounds.Center,
		Size:     bounds.Size,
		SizeSq:   bounds.Size * bounds.Size,
		// Quad:     quad,
	}
}

// rust ((pos.y > center.y) as usize) << 1 | (pos.x > center.x) as usize
func (n Node) FindQuadrant(pos mathx.Vec2) int {
	ybit := 0
	if pos.Y > n.Center.Y {
		ybit = 1
	}
	xbit := 0
	if pos.X > n.Center.X {
		xbit = 1
	}
	return (ybit << 1) | xbit
}

// ---------------------------------------------------------
// Quadtree (Barnes-Hut style) — optimized build (no Propagate)
// ---------------------------------------------------------

type Quadtree[T MassPoint] struct {
	TSq          float32
	ESq          float32
	G            float32
	Nodes        []Node
	MaxNodesUsed int
}

const Root = 0

// epsilon: softening. prevents gravity from becoming absurdly huge. Larger = more softening
func NewQuadtree[T MassPoint](theta, epsilon, G float32) *Quadtree[T] {
	return &Quadtree[T]{
		TSq:   theta * theta,
		ESq:   epsilon * epsilon,
		G:     G,
		Nodes: make([]Node, 0, 1024),
	}
}

// Helper to optimize allocations
func (qt *Quadtree[T]) ReserveForPoints(n int) {
	// Root + expected branch/child nodes.
	// Start with a heuristic, not worst-case.
	want := 1 + (7*n)/2
	if qt.MaxNodesUsed > want {
		want = qt.MaxNodesUsed
	}

	if cap(qt.Nodes) < want {
		qt.Nodes = make([]Node, 0, want)
	}
}

func (qt *Quadtree[T]) Clear(bounds Bounds) {
	qt.Nodes = qt.Nodes[:0]
	qt.Nodes = append(qt.Nodes, NewNode(0, bounds))
}

func (qt *Quadtree[T]) Build(pts []T) {
	qt.ReserveForPoints(len(pts))
	qt.Clear(NewBoundsContaining(pts))
	for _, p := range pts {
		qt.Insert(p.Position(), p.Mass()) // Insert maintains COM incrementally
	}
	if len(qt.Nodes) > qt.MaxNodesUsed {
		qt.MaxNodesUsed = len(qt.Nodes)
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

	parentCenter := qt.Nodes[node].Center
	childSize := qt.Nodes[node].Size * 0.5
	childSizeSq := childSize * childSize

	// quads := qt.Nodes[node].Quad.Subdivide()

	for i := 0; i < 4; i++ {
		center := parentCenter
		center.X += (float32(i&1) - 0.5) * childSize
		center.Y += (float32(i>>1) - 0.5) * childSize

		qt.Nodes = append(qt.Nodes, Node{
			Children: 0,
			Next:     nexts[i],
			Pos:      mathx.V0(),
			Mass:     0,
			Center:   center,
			Size:     childSize,
			SizeSq:   childSizeSq,
		})

		// qt.Nodes = append(qt.Nodes, NewNode(nexts[i], quads[i]))
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
	qt.insertAt(Root, pos, mass)
}

// insertAt inserts a point mass starting at the given node,
// updating COM/mass along the traversed subtree.
func (qt *Quadtree[T]) insertAt(node int, pos mathx.Vec2, mass float32) {
	for {
		n := &qt.Nodes[node]

		// if is Branch: accumulate and descend
		if n.Children != 0 {
			accumulate(n, pos, mass)
			node = n.Children + n.FindQuadrant(pos)
			continue
		}

		// if is Leaf
		if n.Mass == 0 {
			// empty leaf: store point
			n.Pos = pos
			n.Mass = mass
			return
		}

		// occupied leaf: same position => merge mass
		if pos.Equal(n.Pos) {
			n.Mass += mass
			return
		}

		// occupied leaf with different position => split
		oldPos, oldMass := n.Pos, n.Mass
		children := qt.subdivide(node)
		n = &qt.Nodes[node] // reacquire after possible slice reallocation

		// current node becomes branch aggregate
		n.Pos = mathx.V0()
		n.Mass = 0
		accumulate(n, oldPos, oldMass)
		accumulate(n, pos, mass)

		// place old point recursively
		qt.insertAt(children+n.FindQuadrant(oldPos), oldPos, oldMass)

		// continue loop with new point
		node = children + n.FindQuadrant(pos)
		continue
	}
}

func (qt *Quadtree[T]) Acc(pos mathx.Vec2) mathx.Vec2 {
	acc := mathx.V0()

	// Cache locally
	tSq := qt.TSq
	eSq := qt.ESq
	g := qt.G

	node := Root
	for {
		n := qt.Nodes[node]

		d := n.Pos.Sub(pos)
		dSq := d.MagSq()

		// Is leaf or sizeSq < dSq*tSq
		if n.Children == 0 || n.SizeSq < dSq*tSq {
			denom := (dSq + eSq) * mathx.Sqrt(dSq)
			if denom != 0 {
				acc = acc.Add(d.Mul(g * n.Mass / denom))
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
