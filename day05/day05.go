package day05

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseStacks(input []byte) [][]byte {
	stacks := make([][]byte, 0)
	stackLines := make([]string, 0)
	stackCount := 0

	for _, l := range strings.Split(string(input), "\n") {
		if strings.HasPrefix(l, "[") {
			stackLines = append(stackLines, l)
		} else {
			stackCount = len(strings.Replace(l, " ", "", -1))
			break
		}
	}
	fmt.Printf("creating %d stacks from following lines:\n", stackCount)
	for _, line := range stackLines {
		fmt.Println(line)
	}
	for i := 0; i < stackCount; i++ {
		stacks = append(stacks, []byte{})
	}
	for i := len(stackLines) - 1; i >= 0; i-- {
		line := stackLines[i]
		for j := 0; j < stackCount; j++ {
			item := line[1+j*4]
			if item != ' ' {
				stacks[j] = append(stacks[j], item)
			}
		}
	}
	return stacks
}

func parseInput(inputFile string) ([][]byte, [][3]int) {
	input, _ := os.ReadFile(inputFile)

	stacks := parseStacks(input)
	instructions := parseInstructions(input)
	return stacks, instructions
}

func parseInstructions(input []byte) [][3]int {
	instructions := make([][3]int, 0)
	for _, l := range strings.Split(string(input), "\n") {
		if strings.HasPrefix(l, "move") {
			parts := strings.Split(l, " ")
			inst := [3]int{}
			inst[0], _ = strconv.Atoi(parts[1])
			inst[1], _ = strconv.Atoi(parts[3])
			inst[1]--
			inst[2], _ = strconv.Atoi(parts[5])
			inst[2]--
			instructions = append(instructions, inst)
		}
	}
	return instructions
}

func doMoves(stacks [][]byte, instructions [][3]int, moveAtOnce bool) {
	for _, inst := range instructions {
		count := inst[0]
		fromStack := inst[1]
		toStack := inst[2]
		if moveAtOnce {
			items := stacks[fromStack][len(stacks[fromStack])-count:]
			stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-count]
			stacks[toStack] = append(stacks[toStack], items...)
		} else {
			for i := 0; i < count; i++ {
				item := stacks[fromStack][len(stacks[fromStack])-1]
				stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-1]
				stacks[toStack] = append(stacks[toStack], item)
			}
		}
	}
}

func Solve(test bool) (string, string, time.Duration) {
	start := time.Now()
	inputFile := "day05/input.txt"
	stacks, instructions := parseInput(inputFile)
	stacks2 := make([][]byte, len(stacks))
	for i := range stacks {
		stacks2[i] = make([]byte, len(stacks[i]))
		copy(stacks2[i], stacks[i])
	}

	// move one item at a time
	doMoves(stacks, instructions, false)
	res1 := ""
	for _, s := range stacks {
		if len(s) > 0 {
			res1 += string(s[len(s)-1])
		}
	}
	fmt.Println("Task 1: ", res1)

	// move all items at once
	doMoves(stacks2, instructions, true)
	res2 := ""
	for _, s := range stacks2 {
		if len(s) > 0 {
			res2 += string(s[len(s)-1])
		}
	}
	fmt.Println("Task 1: ", res1)

	return res1, res2, time.Since(start)
}
