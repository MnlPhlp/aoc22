package day13

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParts(t *testing.T) {
	parts := getParts("1,[2,[3,[4,[5,6,7]]]],8,9")
	if len(parts) != 4 {
		t.Errorf("Expected 4 parts, got %d", len(parts))
	}
	expected := []string{"1", "[2,[3,[4,[5,6,7]]]]", "8", "9"}
	for i := range expected {
		if parts[i] != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], parts[i])
		}
	}
}

func printItem(i item) string {
	if i.leaf {
		return fmt.Sprintf("%d", i.value)
	}
	text := "["
	for _, item := range i.values {
		text += printItem(item)
		text += ","
	}
	if len(text) > 1 {
		text = text[:len(text)-1]
	}
	text += "]"
	return text
}

func TestExample(t *testing.T) {
	pairs := parseInput("testInput.txt")
	for i, res := range []bool{true, true, false, true, false, true, false, false} {
		ok, finished := checkPair(pairs[i])
		assert.True(t, finished, "Pair %d not finished", i)
		assert.Equal(t, res, ok, "Pair %d not ok", i)
	}
}

func TestParser(t *testing.T) {
	expected, _ := os.ReadFile("input.txt")
	pairs := parseInput("input.txt")
	text := ""
	for _, pair := range pairs {
		text += printItem(pair.left)
		text += "\n"
		text += printItem(pair.right)
		text += "\n\n"
	}
	os.WriteFile("output.txt", []byte(text), 0644)
	assert.Equal(t, string(expected), text)
}
