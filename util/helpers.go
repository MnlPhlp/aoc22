package util

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"golang.org/x/exp/constraints"
)

func ReadInput(day int, test bool) string {
	inputFile := fmt.Sprintf("day%02d/input.txt", day)
	if test {
		inputFile = fmt.Sprintf("day%02d/testInput.txt", day)
	}
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(string(input), "\r\n", "\n")
	}
	return string(input)
}

func Sum[T constraints.Ordered](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
