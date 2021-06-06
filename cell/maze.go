package cell

import (
	"log"
	"math/rand"
)

type Cell struct {
	Dim      Dim
	Subcells []Subcell
}

type Subcell struct {
	Cell          Cell
	RelativePos   Pos
	PassageOffset int
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
	dim := Dim{Width: width, Height: height}
	var possibleSplits []splitType
	if (width > height) && canSplitVert(dim) {
		possibleSplits = append(possibleSplits, splitVert)
	} else if (height > width) && canSplitHor(dim) {
		possibleSplits = append(possibleSplits, splitHor)
	} else {
		if canSplitHor(dim) {
			possibleSplits = append(possibleSplits, splitHor)
		}
		if canSplitVert(dim) {
			possibleSplits = append(possibleSplits, splitVert)
		}
	}
	if len(possibleSplits) == 0 {
		return []Subcell{}
	}
	i := rand.Intn(len(possibleSplits))
	switch split := possibleSplits[i]; split {
	case splitVert:
		d := randomSplitDim(width)
		log.Printf("splitVert, w:%d, d:%d", width, d)
		s1 := Subcell{
			Cell:          NewDividedCell(d, height),
			RelativePos:   Pos{X: 0, Y: 0},
			PassageOffset: randomPassageOffset(height),
		}
		s2 := Subcell{
			Cell:          NewDividedCell(width-d, height),
			RelativePos:   Pos{X: d, Y: 0},
			PassageOffset: randomPassageOffset(height),
		}
		return []Subcell{s1, s2}
	case splitHor:
		d := randomSplitDim(height)
		log.Printf("splitHor, h:%d, d:%d", height, d)
		s1 := Subcell{
			Cell:          NewDividedCell(width, d),
			RelativePos:   Pos{0, 0},
			PassageOffset: randomPassageOffset(width),
		}
		s2 := Subcell{
			Cell:          NewDividedCell(width, height-d),
			RelativePos:   Pos{0, d},
			PassageOffset: randomPassageOffset(width),
		}
		return []Subcell{s1, s2}
	default:
		panic(split)
	}
}

func canSplitHor(dim Dim) bool {
	return dim.Height >= 2 && dim.Width >= 2
}

func canSplitVert(dim Dim) bool {
	return dim.Height >= 2 && dim.Width >= 2
}

func randomSplitDim(width int) int {
	return 1 + rand.Intn(width-1)
}

func randomPassageOffset(d int) int {
	return rand.Intn(d)
}