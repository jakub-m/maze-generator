package cell

type Maze struct {
	RootCell    Cell
	Width       int
	Height      int
	OutterWalls []Wall
}

func NewMaze(width, height int) Maze {
	dim := Dim{width, height}
	c := NewDividedCell(width, height)
	outter := OutterWalls(dim)
	return Maze{
		RootCell:    c,
		Height:      height,
		Width:       width,
		OutterWalls: outter,
	}
}
