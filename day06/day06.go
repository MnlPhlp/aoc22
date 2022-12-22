package day06

import (
	"fmt"
	"os"
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

func getIndex(input []byte, count int) int {
	chars := make([]byte, count)
	for i, c := range input {
		chars[i%count] = c
		if i >= count && !hasDuplicates(chars) {
			return i + 1
		}
	}
	return -1
}

func Solve(test bool) (string, string) {
	inputFile := "day06/input.txt"
	input, _ := os.ReadFile(inputFile)

	// Part 1
	idx1 := getIndex(input, 4)
	fmt.Println("Index for count 4: ", idx1)
	res1 := fmt.Sprintf("%d", idx1)

	// Part 2
	idx2 := getIndex(input, 14)
	fmt.Println("Index for count 14: ", idx2)
	res2 := fmt.Sprintf("%d", idx2)

	return res1, res2
}
