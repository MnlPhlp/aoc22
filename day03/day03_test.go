package day03

import (
	"strings"
	"testing"

	"github.com/mnlphlp/aoc22/util"
)

var input = util.ReadInputUnittest(3, false)
var rucksacks = strings.Split(input, "\n")

func BenchmarkFindDuplicate(b *testing.B) {
	for _, line := range strings.Split(input, "\n") {
		for i := 0; i < b.N; i++ {
			findDuplicate(line)
		}
	}
}

func BenchmarkFindBadge(b *testing.B) {

	for i := 0; i < b.N; i++ {
		findBadges(rucksacks)
	}
}
