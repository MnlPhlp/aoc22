package day22

import (
	"testing"

	"github.com/mnlphlp/aoc22/util"
	"github.com/stretchr/testify/assert"
)

var input = util.ReadInputUnittest(22, false)

func BenchmarkInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInput(input, false)
	}
}

func BenchmarkPart1(b *testing.B) {
	b.StopTimer()
	grid, path := parseInput(input, false)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		part1(grid, path, false)
	}
}

func TestPart1(t *testing.T) {
	grid, path := parseInput(input, false)
	assert.Equal(t, 93226, part1(grid, path, false))
}

func TestPart2(t *testing.T) {
	grid, path := parseInput(input, false)
	assert.Equal(t, 37415, part2(grid, path, false))
}
