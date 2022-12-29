package day18

import (
	"testing"

	"github.com/mnlphlp/aoc22/util"
	"github.com/stretchr/testify/assert"
)

var droplet = parseInput(util.ReadInputUnittest(18, false))
var testDroplet = parseInput(util.ReadInputUnittest(18, true))

func TestTask1TestInput(t *testing.T) {
	assert.Equal(t, 64, surfaceArea(testDroplet))
}

func TestTask1RealInput(t *testing.T) {
	assert.Equal(t, 3466, surfaceArea(droplet))
}

func TestTask2TestInput(t *testing.T) {
	assert.Equal(t, 64-58, pocketSurfaceArea(testDroplet))
}

func TestTask2RealInput(t *testing.T) {
	assert.Equal(t, 3466-2012, pocketSurfaceArea(droplet))
}

func TestSolve(t *testing.T) {
	res1, res2 := Solve(util.ReadInputUnittest(18, false), false, 0)
	assert.Equal(t, "3466", res1)
	assert.Equal(t, "2012", res2)
}
