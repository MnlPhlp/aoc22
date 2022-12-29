package day17

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask1TestInput(t *testing.T) {
	inputStr, _ := os.ReadFile("testInput.txt")
	input := parseInput(string(inputStr))
	res1 := simulateDrops(input, false, 2022)
	assert.Equal(t, "3068", res1)
}

func TestTask1RealInput(t *testing.T) {
	inputStr, _ := os.ReadFile("input.txt")
	input := parseInput(string(inputStr))
	res1 := simulateDrops(input, false, 2022)
	assert.Equal(t, "3109", res1)
}

func TestTask2TestInput(t *testing.T) {
	inputStr, _ := os.ReadFile("testInput.txt")
	input := parseInput(string(inputStr))
	res2 := simulateDrops(input, true, 1000000000000)
	assert.Equal(t, "1514285714288", res2)
}

func TestTask2RealInput(t *testing.T) {
	inputStr, _ := os.ReadFile("input.txt")
	input := parseInput(string(inputStr))
	res2 := simulateDrops(input, false, 1000000000000)
	assert.Equal(t, "1541449275365", res2)
}
