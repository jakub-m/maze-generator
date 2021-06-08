package svg_test

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"maze/cell"
	"maze/svg"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rand.Seed(0)
	os.Exit(m.Run())
}

func TestSvg(t *testing.T) {
	m := cell.NewMaze(4)
	f, err := svg.FormatMaze(m, 50, 2)
	assert.NoError(t, err)
	s := string(f)
	assert.Contains(t, s, "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>")
	assert.Contains(t, s, "<svg")
	assert.Contains(t, s, "</svg>")
}
