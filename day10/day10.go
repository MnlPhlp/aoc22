package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	cycles int
	value  int
}

func parseInput(input string) []instruction {
	instructions := []instruction{}
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

func Solve(input string, test bool, task int) (string, string) {
	instructions := parseInput(input)
	res1, res2 := "", ""
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

	if task != 2 {
		// Task 1
		valueSum := 0
		for c, val := range values {
			if (c-19)%40 == 0 {
				fmt.Printf("Cycle %d: %d\n", c+1, val)
				valueSum += val * (c + 1)
			}
		}
		fmt.Printf("Value sum: %d\n", valueSum)
		res1 = strconv.Itoa(valueSum)
	}

	if task != 1 {
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
		// print screen
		for _, row := range screen {
			for _, pixel := range row {
				if pixel {
					res2 += "##"
				} else {
					res2 += "  "
				}
			}
			res2 += "\n"
		}
		fmt.Printf("Screen:\n%s", res2)
	}
	return res1, res2
}
