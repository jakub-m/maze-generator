package cell_test

import (
	"github.com/stretchr/testify/assert"
	"maze/cell"
	"maze/pretty"
	"testing"
)

func TestWalls2(t *testing.T) {
	c := cell.NewDividedCell(2, 2)
	ws := cell.InternalWalls(c)
	assert.Equal(t, 1, len(ws), "%v\n%v", ws, pretty.Format(c))
}

func TestWalls4(t *testing.T) {
	c := cell.NewDividedCell(4, 4)
	ws := cell.InternalWalls(c)
	assert.GreaterOrEqual(t, len(ws), 3, "%v", ws)
}
