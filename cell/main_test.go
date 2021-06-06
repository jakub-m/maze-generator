package cell_test

import (
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rand.Seed(0)
	os.Exit(m.Run())
}
