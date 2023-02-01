package util

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
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
	inputStr := string(input)
	if runtime.GOOS == "windows" {
		inputStr = strings.ReplaceAll(inputStr, "\r\n", "\n")
	}
	// remove trailing newline
	if inputStr[len(inputStr)-1] == '\n' {
		inputStr = inputStr[:len(inputStr)-1]
	}
	return inputStr
}

func ReadInputUnittest(day int, test bool) string {
	path, _ := os.Getwd()
	if !strings.HasSuffix(path, "aoc22") {
		os.Chdir("..")
	}
	return ReadInput(day, test)
}

func Sum[T constraints.Ordered](arr ...T) T {
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

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Abs[T constraints.Integer | constraints.Float](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Sign[T constraints.Signed](a T) T {
	if a < 0 {
		return -1
	}
	return 1
}

func Contains[T comparable](arr []T, elem T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](arr []T, elem T) int {
	for i, v := range arr {
		if v == elem {
			return i
		}
	}
	return -1
}
