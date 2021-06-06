package maze

import "maze/cell"

type Maze struct {
	RootCell cell.Cell
	Width int
	Height int
}

func NewMaze(size int) Maze {
	c := cell.NewDividedCell(size, size)
	return Maze {
		RootCell: c,
		Height: size,
		Width: size,
	}
}
