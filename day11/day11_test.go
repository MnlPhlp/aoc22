package day11

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	input, _ := os.ReadFile("testInput.txt")
	monkeys := parseInput(string(input))
	fmt.Printf("%v", monkeys)
	assert.Equal(t, 4, len(monkeys))
	assert.Equal(t, 19, monkeys[0].operation(1))
	assert.Equal(t, 7, monkeys[1].operation(1))
	assert.Equal(t, 4, monkeys[2].operation(2))
	assert.Equal(t, 4, monkeys[3].operation(1))
}
