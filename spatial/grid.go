package spatial

import (
	"github.com/Spencer1O1/powder_space/v2/mathx"
)

type CellCoord struct {
	X int
	Y int
}

type UniformGrid struct {
	CellSize    float32
	InvCellSize float32

	Cells       map[CellCoord][]int
	ActiveCells []CellCoord
}

func NewUniformGrid(cellSize float32) *UniformGrid {
	return &UniformGrid{
		CellSize:    cellSize,
		InvCellSize: 1.0 / cellSize,
		Cells:       make(map[CellCoord][]int),
		ActiveCells: make([]CellCoord, 0, 256),
	}
}

func (g *UniformGrid) Clear() {
	for _, cell := range g.ActiveCells {
		g.Cells[cell] = g.Cells[cell][:0]
	}
	g.ActiveCells = g.ActiveCells[:0]
}

func (g *UniformGrid) Insert(index int, pos mathx.Vec2) {
	cell := g.CellFor(pos)
	indices, exists := g.Cells[cell]

	if !exists {
		// Preallocate for 8 particles
		indices = make([]int, 0, 8)
		g.Cells[cell] = indices
	}

	if len(indices) == 0 {
		g.ActiveCells = append(g.ActiveCells, cell)
	}

	g.Cells[cell] = append(indices, index)
}

func (g *UniformGrid) ForEachNeighborCell(c CellCoord, fn func(CellCoord)) {
	neighbors := [9]CellCoord{
		{c.X - 1, c.Y - 1},
		{c.X, c.Y - 1},
		{c.X + 1, c.Y - 1},
		{c.X - 1, c.Y},
		{c.X, c.Y},
		{c.X + 1, c.Y},
		{c.X - 1, c.Y + 1},
		{c.X, c.Y + 1},
		{c.X + 1, c.Y + 1},
	}

	for _, n := range neighbors {
		fn(n)
	}
}

// If world coordinates are always positive
// func (g *UniformGrid) CellFor(pos mathx.Vec2) CellCoord {
// 	return CellCoord{
// 		X: int(pos.X * g.InvCellSize),
// 		Y: int(pos.Y * g.InvCellSize),
// 	}
// }

// If world coordinates not always positive
func (g *UniformGrid) CellFor(pos mathx.Vec2) CellCoord {
	x := pos.X * g.InvCellSize
	y := pos.Y * g.InvCellSize

	return CellCoord{
		X: mathx.FastFloorToInt(x),
		Y: mathx.FastFloorToInt(y),
	}
}
