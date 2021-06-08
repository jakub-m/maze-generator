package cell

type Maze struct {
	RootCell    Cell
	Width       int
	Height      int
	OutterWalls []Wall
}

func NewMaze(size int) Maze {
	dim := Dim{size, size}
	c := NewDividedCell(size, size)
	outter := OutterWalls(dim)
	return Maze{
		RootCell:    c,
		Height:      size,
		Width:       size,
		OutterWalls: outter,
	}
}
