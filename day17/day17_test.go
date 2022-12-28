package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = parseInput("testInput.txt")
var realInput = parseInput("input.txt")

func TestTask1TestInput(t *testing.T) {
	res1 := simulateDrops(testInput, false, 2022)
	assert.Equal(t, "3068", res1)
}

func TestTask1RealInput(t *testing.T) {
	res1 := simulateDrops(realInput, false, 2022)
	assert.Equal(t, "3109", res1)
}
