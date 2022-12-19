package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
	cycles int
	value  int
}

func parseInput(inputFile string) []instruction {
	instructions := []instruction{}
	input, _ := os.ReadFile(inputFile)
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		instruction := instruction{}
		if line == "noop" {
			instruction.cycles = 1
			instruction.value = 0
		} else {
			instruction.cycles = 2
			instruction.value, _ = strconv.Atoi(strings.Split(line, " ")[1])
		}
		instructions = append(instructions, instruction)
	}
	return instructions
}

func Solve(test bool) (string, string, time.Duration) {
	start := time.Now()
	inputFile := "day10/input.txt"
	if test {
		inputFile = "day10/testInput.txt"
	}
	instructions := parseInput(inputFile)

	// store value for each cycle
	cycle := 0
	value := 1
	values := make([]int, 0, len(instructions))
	for _, instruction := range instructions {
		for i := 0; i < instruction.cycles; i++ {
			cycle++
			values = append(values, value)
		}
		value += instruction.value
	}

	// Task 1
	valueSum := 0
	for c, val := range values {
		if (c-19)%40 == 0 {
			fmt.Printf("Cycle %d: %d\n", c+1, val)
			valueSum += val * (c + 1)
		}
	}
	fmt.Printf("Value sum: %d\n", valueSum)
	res1 := strconv.Itoa(valueSum)

	// Task 2
	screenWidth := 40
	cycle = 0
	screen := [6][40]bool{}
	for c, pos := range values {
		// check if printed pixel and sprite overlap
		if c%screenWidth-pos >= -1 && c%screenWidth-pos <= 1 {
			screen[c/screenWidth][c%screenWidth] = true
		}
	}
	res2 := ""
	// print screen
	for _, row := range screen {
		for _, pixel := range row {
			if pixel {
				res2 += "##"
			} else {
				res2 += ".."
			}
		}
		res2 += "\n"
	}
	fmt.Printf("Screen:\n%s", res2)
	return res1, "<img>", time.Since(start)
}
