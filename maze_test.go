package maze

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rand.Seed(0)
	os.Exit(m.Run())
}

//func TestPrint(t *testing.T) {
//	c := NewDividedCell(12, 12)
//	json, err := json.MarshalIndent(c, "", " ")
//	assert.NoError(t, err)
//	fmt.Print(string(json) + "\n")
//}

func TestSmallCell(t *testing.T) {
	ch := NewDividedCell(2, 1)
	dim := ch.Dim
	assert.Equal(t, 2, dim.Width)
	assert.Equal(t, 1, dim.Height)
	assert.Empty(t, ch.Subcells)
}

func TestCellDiv2x2(t *testing.T) {
	ch := NewDividedCell(2, 2)
	dim := ch.Dim
	assert.Equal(t, 2, dim.Width)
	assert.Equal(t, 2, dim.Height)
	assert.Len(t, ch.Subcells, 2)
	for _, subcell := range ch.Subcells {
		r := subcell.RelativePos
		assert.True(t, (r.X == 0 && r.Y == 0) || (r.X == 1 && r.Y == 0) || (r.X == 0 && r.Y == 1), "relative pos wrong: %v\nmain: %v", r, ch)
	}
}

func TestCellDivBasicProps(t *testing.T) {
	tc := []Dim{
		{2, 2},
		{3, 3},
		{4, 4},
		{9, 9},
		{128, 128},
		{111, 11},
		{11, 111},
	}
	for _, dim := range tc {
		t.Run(fmt.Sprintf("Dim(%v)", dim), func (t *testing.T) {
			cell := NewDividedCell(dim.Width, dim.Height)
			testBasicCellProps(t, cell)
			testLeafCellsArea(t, cell)
		})
	}
}

func testBasicCellProps(t *testing.T, main Cell) {
	cells := []Cell{main}
	for {
		if len(cells) == 0 {
			break
		}
		cell := cells[0]
		cells = cells[1:]
		subs := cell.Subcells
		assert.Equal(t, 2, len(subs), "Cell not divided into two: %v\n main:%v", cell, main) // assume division 2
		for _, sub := range subs {
			rel := sub.RelativePos
			assert.True(t, rel.X >= 0)
			assert.True(t, rel.Y >= 0)
			assert.True(t, rel.X + sub.Cell.Dim.Width <= cell.Dim.Width)
			assert.True(t, rel.Y + sub.Cell.Dim.Height <= cell.Dim.Height)
		}
	}
}

func testLeafCellsArea(t *testing.T, main Cell) {
	leafs := collectLeafCells(main)
	leafArea := 0
	for _, leaf := range leafs {
		leafArea += leaf.Dim.Height * leaf.Dim.Width
	}
	mainArea := main.Dim.Height * main.Dim.Width
	assert.Equal(t, mainArea, leafArea)
}

func collectLeafCells(c Cell) []Cell {
	if len(c.Subcells) == 0 {
		return []Cell{c}
	} else {
		leafs := []Cell{}
		for _, sub := range c.Subcells {
			leafs = append(leafs, collectLeafCells(sub.Cell)...)
		}
		return leafs
	}
}