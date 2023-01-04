package day21

import (
	"testing"

	"github.com/mnlphlp/aoc22/util"
)

var input = util.ReadInputUnittest(21, false)

func BenchmarkParsing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInput(input)
	}
}

func BenchmarkTask1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		monkeys := parseInput(input)
		b.StartTimer()
		solveMonkey(monkeys, "root")
	}
}

func BenchmarkTask2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		monkeys := parseInput(input)
		solveMonkey(monkeys, "root")
		b.StartTimer()
		solveHumn(monkeys, false)
	}
}
