package day04

import (
	"fmt"
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

func parseInput(input string) []pair {
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

func Solve(input string, test bool, task int) (string, string) {
	pairs := parseInput(input)
	res1, res2 := "", ""
	count := 0

	if task != 2 {
		for _, p := range pairs {
			if p.fullyOverlap() {
				count++
			}
		}
		fmt.Printf("fully overlapping pairs: %v\n", count)
		res1 = strconv.Itoa(count)
	}
	if task != 1 {
		count = 0
		for _, p := range pairs {
			if p.overlap() {
				count++
			}
		}
		fmt.Printf("overlapping pairs: %v\n", count)
		res2 = strconv.Itoa(count)
	}
	return res1, res2
}
