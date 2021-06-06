package maze

import "maze/cell"

type Maze struct {
	RootCell cell.Cell
	Width int
	Height int
	OutterWalls []cell.Wall
}

func NewMaze(size int) Maze {
	dim := cell.Dim{size, size}
	c := cell.NewDividedCell(size, size)
	outter := cell.OutterWalls(dim)
	return Maze {
		RootCell: c,
		Height: size,
		Width: size,
		OutterWalls: outter,
	}
}
