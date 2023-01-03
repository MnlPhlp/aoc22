package day20

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mnlphlp/aoc22/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMoveRight(t *testing.T) {
	f := File{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	f.Move(0, 1)
	assert.Equal(t, "[{0 1} {1 0} {2 2}]", fmt.Sprint([]Item(f)))
}

func TestMoveRightWrap(t *testing.T) {
	f := File{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	f.Move(2, 1)
	assert.Equal(t, "[{0 0} {1 2} {2 1}]", fmt.Sprint([]Item(f)))
}

func TestMoveLeft(t *testing.T) {
	f := File{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	f.Move(2, -1)
	assert.Equal(t, "[{0 0} {1 2} {2 1}]", fmt.Sprint([]Item(f)))
}

func TestMoveLeftWrap(t *testing.T) {
	f := File{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	f.Move(2, -2)
	assert.Equal(t, "[{0 0} {1 1} {2 2}]", fmt.Sprint([]Item(f)))
}

func TestMoveStartLeft(t *testing.T) {
	f := File{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	f.Move(1, -1)
	assert.Equal(t, "[0 2 1]", f.String())
	assert.False(t, f.HasDuplicatePos())
}

func TestMoveEnd(t *testing.T) {
	f := File{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	f.Move(1, 1)
	assert.Equal(t, "[1 0 2]", f.String())
}

func TestGet(t *testing.T) {
	f := File{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	assert.Equal(t, "[1 2 0]", fmt.Sprint(f.Get(1, 2, 3)))
}

func TestGetWrapOtherStart(t *testing.T) {
	f := File{
		{1, 0},
		{0, 1},
		{2, 2},
	}
	assert.Equal(t, "[2 1 0]", fmt.Sprint(f.Get(1, 2, 3)))
}

func TestExampleMoves(t *testing.T) {
	file := File{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 4}}
	file.Move(3, 4)
	assert.Equal(t, "[1 2 3 4 5]", file.String())
}

func TestParser(t *testing.T) {
	input := util.ReadInputUnittest(20, false)
	file := parseInput(input)
	for i, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}
		assert.Equal(t, l, fmt.Sprint(file[i].Val))
		assert.Equal(t, i, file[i].Pos)
	}
}

func TestSolve(t *testing.T) {
	input := util.ReadInputUnittest(20, true)
	res1, res2 := Solve(input, false, 0)
	require.Equal(t, "3", res1)
	require.Equal(t, "1623178306", res2)
	input = util.ReadInputUnittest(20, false)
	res1, res2 = Solve(input, false, 0)
	assert.Equal(t, "2275", res1)
	assert.Equal(t, "4090409331120", res2)
}
