package day08

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	// read input into 2d array
	f, _ := os.ReadFile("day08/input.txt")
	lines := strings.Split(string(f), "\n")
	grid := make([][]int, len(lines))
	for row, line := range lines {
		grid[row] = make([]int, len(line))
		for col, char := range line {
			grid[row][col] = int(char - '0')
		}
	}
	// remove empty last line
	if len(grid[len(grid)-1]) == 0 {
		grid = grid[:len(grid)-1]
	}
	return grid
}

func isVisible(grid [][]int, row, col int) (bool, int) {
	height := grid[row][col]
	visible := false
	score := 1
	viewDistance := 1
	if row == 22 && col == 56 {
		fmt.Println("here")
	}
	// check top
	for r := row - 1; r >= 0 && grid[r][col] < height; r-- {
		viewDistance++
		if r == 0 {
			visible = true
			viewDistance -= 1
		}
	}
	score *= viewDistance
	viewDistance = 1
	// check right
	for c := col + 1; c < len(grid[row]) && grid[row][c] < height; c++ {
		viewDistance++
		if c == len(grid[row])-1 {
			visible = true
			viewDistance -= 1
		}
	}
	score *= viewDistance
	viewDistance = 1
	// check bottom
	for r := row + 1; r < len(grid) && grid[r][col] < height; r++ {
		viewDistance++
		if r == len(grid)-1 {
			visible = true
			viewDistance -= 1
		}
	}
	score *= viewDistance
	viewDistance = 1
	// check left
	for c := col - 1; c >= 0 && grid[row][c] < height; c-- {
		viewDistance++
		if c == 0 {
			visible = true
			viewDistance -= 1
		}
	}
	score *= viewDistance
	return visible, score
}

func countVisible(grid [][]int) (int, int, [2]int) {
	visible := 0
	maxScenic := 0
	scenicPos := [2]int{}
	// add borders and remove double counted corners
	visible += len(grid[0])*2 + len(grid)*2 - 4
	// add all other visible trees
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			v, scenic := isVisible(grid, row, col)
			if v {
				visible++
			}
			if scenic > maxScenic {
				maxScenic = scenic
				scenicPos = [2]int{row, col}
			}
		}
	}
	return visible, maxScenic, scenicPos
}

func Solve(test bool) (string, string) {
	// read input into 2d array
	grid := parseInput()
	// count visible trees
	visible, mostScenic, scenicPos := countVisible(grid)
	fmt.Println("Task 1: ", visible)
	fmt.Println("Task 2: ", mostScenic, " at ", scenicPos[0], ",", scenicPos[1])
	return strconv.Itoa(visible), strconv.Itoa(mostScenic)
}
