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

func ReadInputUnittest(day int, test bool) string {
	path, _ := os.Getwd()
	if !strings.HasSuffix(path, "aoc22") {
		os.Chdir("..")
	}
	return ReadInput(day, test)
}

func Sum[T constraints.Ordered](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

func Mul[T constraints.Integer | constraints.Float](arr []T) T {
	mul := T(1)
	for _, v := range arr {
		mul *= v
	}
	return mul
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](args ...T) T {
	max := 0
	for i := 1; i < len(args); i++ {
		if args[i] > args[max] {
			max = i
		}
	}
	return args[max]
}
