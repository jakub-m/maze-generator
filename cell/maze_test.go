package cell_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"maze/cell"
	"testing"
)

func TestSmallCell(t *testing.T) {
	ch := cell.NewDividedCell(2, 1)
	dim := ch.Dim
	assert.Equal(t, 2, dim.Width)
	assert.Equal(t, 1, dim.Height)
	assert.Empty(t, ch.Subcells)
}

func TestCellDiv2x2(t *testing.T) {
	ch := cell.NewDividedCell(2, 2)
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
	tc := []cell.Dim{
		{2, 2},
		//{3, 3},
		//{4, 4},
		//{9, 9},
		//{128, 128},
		//{111, 11},
		//{11, 111},
	}
	for _, dim := range tc {
		t.Run(fmt.Sprintf("Dim(%v)", dim), func(t *testing.T) {
			c := cell.NewDividedCell(dim.Width, dim.Height)
			testBasicCellProps(t, c)
			testLeafCellsArea(t, c)
			testNoDots(t, c)
		})
	}
}

func testBasicCellProps(t *testing.T, main cell.Cell) {
	cells := []cell.Cell{main}
	for {
		if len(cells) == 0 {
			break
		}
		c := cells[0]
		cells = cells[1:]
		subs := c.Subcells
		assert.Equal(t, 2, len(subs), "Cell not divided into two: %v\n main:%v", c, main) // assume division 2
		for _, sub := range subs {
			rel := sub.RelativePos
			assert.True(t, rel.X >= 0)
			assert.True(t, rel.Y >= 0)
			assert.True(t, rel.X+sub.Cell.Dim.Width <= c.Dim.Width)
			assert.True(t, rel.Y+sub.Cell.Dim.Height <= c.Dim.Height)
		}
	}
}

func testLeafCellsArea(t *testing.T, main cell.Cell) {
	leafs := collectLeafCells(main)
	leafArea := 0
	for _, leaf := range leafs {
		leafArea += leaf.Dim.Height * leaf.Dim.Width
	}
	mainArea := main.Dim.Height * main.Dim.Width
	assert.Equal(t, mainArea, leafArea)
}

func testNoDots(t *testing.T, c cell.Cell) {
	iterAllCells(t, c, func(t *testing.T, c cell.Cell) {
		area := c.Dim.Width * c.Dim.Height
		assert.NotEqual(t, 1, area, c)
	})
}

func collectLeafCells(c cell.Cell) []cell.Cell {
	if len(c.Subcells) == 0 {
		return []cell.Cell{c}
	} else {
		leafs := []cell.Cell{}
		for _, sub := range c.Subcells {
			leafs = append(leafs, collectLeafCells(sub.Cell)...)
		}
		return leafs
	}
}

func iterAllCells(t *testing.T, c cell.Cell, fn func(t *testing.T, c cell.Cell)) {
	fn(t, c)
	for _, sub := range c.Subcells {
		iterAllCells(t, sub.Cell, fn)
	}
}