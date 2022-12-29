package day13

import (
	"fmt"
	"testing"

	"github.com/mnlphlp/aoc22/util"
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
	pairs := parseInput(util.ReadInputUnittest(13, true))
	for i, res := range []bool{true, true, false, true, false, true, false, false} {
		ok, finished := checkPair(pairs[i])
		assert.True(t, finished, "Pair %d not finished", i)
		assert.Equal(t, res, ok, "Pair %d not ok", i)
	}
}

func TestParser(t *testing.T) {
	expected := util.ReadInputUnittest(13, false)
	pairs := parseInput(expected)
	text := ""
	for _, pair := range pairs {
		text += pair.left.String()
		text += "\n"
		text += pair.right.String()
		text += "\n\n"
	}
	assert.Equal(t, expected, text)
}
