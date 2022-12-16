package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type secRange struct {
	start int
	end   int
}

type pair [2]secRange

func (p pair) fullyOverlap() bool {
	return p[0].start >= p[1].start && p[0].end <= p[1].end || p[1].start >= p[0].start && p[1].end <= p[0].end
}

func (p pair) overlap() bool {
	return p[0].start >= p[1].start && p[0].start <= p[1].end || p[1].start >= p[0].start && p[1].start <= p[0].end
}

func parseInput() []pair {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	pairs := make([]pair, len(lines))
	for i, line := range lines {
		ranges := strings.Split(line, ",")
		for j, r := range ranges {
			sec := strings.Split(r, "-")
			start, _ := strconv.Atoi(sec[0])
			end, _ := strconv.Atoi(sec[1])
			pairs[i][j] = secRange{start, end}
		}
	}
	return pairs
}

func main() {
	pairs := parseInput()

	count := 0
	for _, p := range pairs {
		if p.fullyOverlap() {
			count++
		}
	}
	fmt.Printf("fully overlapping pairs: %v\n", count)

	count = 0
	for _, p := range pairs {
		if p.overlap() {
			count++
		}
	}
	fmt.Printf("overlapping pairs: %v\n", count)
}
