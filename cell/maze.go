package cell

import "math/rand"

type Cell struct {
	Dim      Dim
	Subcells []Subcell
}

type Subcell struct {
	Cell        Cell
	RelativePos Pos
}

type Pos struct {
	X int
	Y int
}

func (p Pos) Add(q Pos) Pos {
	return Pos{X: p.X + q.X, Y: p.Y + q.Y}
}

type Dim struct {
	Width  int
	Height int
}

func NewDividedCell(width, height int) Cell {
	subcells := generateSubcells(width, height)
	return Cell{
		Dim:      Dim{Height: height, Width: width},
		Subcells: subcells,
	}
}

type splitType int

const (
	splitVert splitType = iota
	splitHor
)

func generateSubcells(width, height int) []Subcell {
	switch split := randomSplitType(); split {
	case splitVert:
		if width < 2 {
			return []Subcell{}
		} else {
			d := randomSplitDim(width)
			s1 := Subcell{
				Cell:        NewDividedCell(d, height),
				RelativePos: Pos{X: 0, Y: 0},
			}
			s2 := Subcell{
				Cell:        NewDividedCell(width-d, height),
				RelativePos: Pos{X: d, Y: 0},
			}
			return []Subcell{s1, s2}
		}
	case splitHor:
		if height < 2 {
			return []Subcell{}
		} else {
			d := randomSplitDim(height)
			s1 := Subcell{
				Cell:        NewDividedCell(width, d),
				RelativePos: Pos{0, 0},
			}
			s2 := Subcell{
				Cell:        NewDividedCell(width, height-d),
				RelativePos: Pos{0, d},
			}
			return []Subcell{s1, s2}
		}
	default:
		panic(split)
	}
}

func randomSplitType() splitType {
	if rand.Intn(2) == 0 {
		return splitHor
	} else {
		return splitVert
	}
}

func randomSplitDim(width int) int {
	return 1 + rand.Intn(width-1)
}
