package day06

import (
	"fmt"
)

func hasDuplicates(input []byte) bool {
	seen := make(map[byte]bool)
	for _, c := range input {
		if seen[c] {
			return true
		}
		seen[c] = true
	}
	return false
}

func getIndex(input string, count int) int {
	chars := make([]byte, count)
	for i, c := range []byte(input) {
		chars[i%count] = c
		if i >= count && !hasDuplicates(chars) {
			return i + 1
		}
	}
	return -1
}

func Solve(input string, test bool, task int) (string, string) {
	res1, res2 := "", ""

	if task != 2 {
		// Part 1
		idx1 := getIndex(input, 4)
		fmt.Println("Index for count 4: ", idx1)
		res1 = fmt.Sprintf("%d", idx1)
	}

	if task != 1 {
		// Part 2
		idx2 := getIndex(input, 14)
		fmt.Println("Index for count 14: ", idx2)
		res2 = fmt.Sprintf("%d", idx2)
	}

	return res1, res2
}
