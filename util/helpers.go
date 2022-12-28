package util

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func ReadInput(day int, test bool) string {
	inputFile := fmt.Sprintf("day%02d/input.txt", day)
	if test {
		inputFile = fmt.Sprintf("day%02d/testInput.txt", day)
	}
	input, _ := os.ReadFile(inputFile)
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(string(input), "\r\n", "\n")
	}
	return string(input)
}
