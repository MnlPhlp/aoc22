package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = parseInput("testInput.txt")
var realInput = parseInput("input.txt")

func TestTask1TestInput(t *testing.T) {
	res1 := task1(testInput, false, 2022)
	assert.Equal(t, "3068", res1)
}

func TestTask1RealInput(t *testing.T) {
	res1 := task1(realInput, false, 2022)
	assert.Equal(t, "3068", res1)
}

func TestTask2TestInput(t *testing.T) {
	res2 := task2(testInput, false)
	assert.Equal(t, "", res2)
}

func TestTask2RealInput(t *testing.T) {
	res2 := task2(realInput, false)
	assert.Equal(t, "", res2)
}
