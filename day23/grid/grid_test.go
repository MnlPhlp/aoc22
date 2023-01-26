package grid

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mnlphlp/aoc22/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInput = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func TestInsert(t *testing.T) {
	grid := ParseInput(testInput)
	testPos := []util.Pos2{}
	for i := 0; i < 10; i++ {
		testPos = append(testPos, util.Pos2{rand.Intn(200) - 100, rand.Intn(100) - 100})
	}
	fmt.Println(testPos)
	for i, pos := range testPos {
		grid = grid.Insert(pos)
		for _, test := range testPos[:i+1] {
			require.True(t, grid.Contains(test), "Error when testing %v after inserting %v", test, pos)
		}
	}
}

func TestInsert2(t *testing.T) {
	g := Grid{}
	a := util.Pos2{-3, 9}
	b := util.Pos2{-4, -1}
	g = g.Insert(a)
	require.True(t, g.Contains(a))
	g = g.Insert(b)
	require.True(t, g.Contains(a))
	require.True(t, g.Contains(b))
}

func TestContains(t *testing.T) {
	grid := ParseInput(testInput)
	testPos := []util.Pos2{
		{0, 2},
		{2, 1},
		{3, 1},
	}
	for _, pos := range testPos {
		require.True(t, grid.Contains(pos), "%v not found in grid", pos)
	}
}

func TestHasNeighbors(t *testing.T) {
	g := Grid{}
	for i := 0; i < 100; i++ {
		x, y := rand.Intn(200), rand.Intn(200)
		dir := DirectionGroups[rand.Intn(200)%4][rand.Intn(200)%3]
		p := util.Pos2{x, y}
		g = g.Insert(p)
		n := util.Pos2{x + dir.X, y + dir.Y}
		g = g.Insert(n)
		assert.True(t, g.HasNeighbor(p), "%v not registered as neighbor %v", n, p)
	}
}
