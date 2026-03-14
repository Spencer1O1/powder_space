package spatial

import (
	"math"

	"github.com/Spencer1O1/powder_space/v2/mathx"
)

type CellCoord struct {
	X int
	Y int
}

type UniformGrid struct {
	CellSize float64
	Cells    map[CellCoord][]int
}

func NewUniformGrid(cellSize float64) *UniformGrid {
	return &UniformGrid{
		CellSize: cellSize,
		Cells:    make(map[CellCoord][]int),
	}
}

func (g *UniformGrid) Clear() {
	for k := range g.Cells {
		delete(g.Cells, k)
	}
}

func (g *UniformGrid) CellFor(pos mathx.Vec2) CellCoord {
	return CellCoord{
		X: int(math.Floor(pos.X / g.CellSize)),
		Y: int(math.Floor(pos.Y / g.CellSize)),
	}
}

func (g *UniformGrid) Insert(index int, pos mathx.Vec2) {
	cell := g.CellFor(pos)
	g.Cells[cell] = append(g.Cells[cell], index)
}
