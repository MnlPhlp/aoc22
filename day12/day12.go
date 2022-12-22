package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleErr[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

type pos struct {
	x, y int
}
type node struct {
	pos
	height  int
	f, g, h int
	parent  *node
}

func (n *node) calc(goal *node) {
	n.g = n.parent.g + 1
	n.h = abs(n.x-goal.x) + abs(n.y-goal.y)
	n.f = n.g + n.h
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (n node) isIn(nodes []*node) bool {
	for _, node := range nodes {
		if n.pos == node.pos {
			return true
		}
	}
	return false
}

func parse(inputFile string) (grid [][]*node, start *node, goal *node) {
	input := handleErr(os.ReadFile(inputFile))
	grid = make([][]*node, 0)
	for y, line := range strings.Split(string(input), "\n") {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, make([]*node, len(line)))
		for x, c := range line {
			grid[y][x] = &node{
				pos: pos{
					x: x,
					y: y,
				},
				height: int(c - 'a'),
				parent: nil,
			}
			if c == 'S' {
				grid[y][x].height = 0
				start = grid[y][x]
			} else if c == 'E' {
				grid[y][x].height = 25
				goal = grid[y][x]
			}
		}
	}
	return
}

func getNeighbors(grid [][]*node, current *node, closed map[pos]bool) []*node {
	steps := make([]*node, 0)
	for _, move := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		newX, newY := current.x+move[0], current.y+move[1]
		if newX < 0 || newX >= len(grid[0]) || newY < 0 || newY >= len(grid) {
			continue
		}
		newPos := grid[newY][newX]
		if closed[newPos.pos] {
			continue
		}
		newPos.parent = current
		hightDiff := grid[newPos.y][newPos.x].height - grid[current.y][current.x].height
		if hightDiff > 1 {
			continue
		}
		steps = append(steps, newPos)
	}
	return steps
}

func findPath(grid [][]*node, start *node, goal *node) ([][2]int, bool) {
	closed := make(map[pos]bool)
	open := make([]*node, 1)
	open[0] = start
	current := start
	for len(open) > 0 {
		// get node with lowest f
		min := 0
		for i := range open {
			if open[i].f < open[min].f {
				min = i
			}
		}
		current = open[min]
		open = append(open[:min], open[min+1:]...)
		closed[current.pos] = true
		// if reached goal, return path
		if current.pos == goal.pos {
			break
		}
		// get neighbors
		neighbors := getNeighbors(grid, current, closed)
		for _, n := range neighbors {
			if closed[n.pos] {
				continue
			}
			n.calc(goal)
			if !n.isIn(open) {
				open = append(open, n)
			}
		}

	}
	for _, row := range grid {
		for _, node := range row {
			parentPos := pos{-1, -1}
			if node.parent != nil {
				parentPos = node.parent.pos
			}
			if node.pos == parentPos {
				fmt.Printf("Self reference: %v/%v", node.pos, parentPos)
			}
		}
	}
	// if reached goal, return path
	if current.pos == goal.pos {
		path := make([][2]int, 0)
		for current != nil {
			path = append(path, [2]int{current.x, current.y})
			current = current.parent
		}
		return path, true
	} else {
		return nil, false
	}
}

func resetGrid(grid [][]*node) {
	for _, row := range grid {
		for _, node := range row {
			node.parent = nil
			node.f = 0
			node.g = 0
			node.h = 0
		}
	}
}

func Solve(test bool) (string, string) {
	var inputFile string
	if test {
		inputFile = "day12/testInput.txt"
	} else {
		inputFile = "day12/input.txt"
	}
	grid, start, goal := parse(inputFile)
	fmt.Println("start: ", start.pos)
	fmt.Println("goal: ", goal.pos)

	path, _ := findPath(grid, start, goal)
	if test {
		fmt.Println("Pfad: ", path)
	}

	fmt.Println("Result Task 1: ", len(path)-1)
	res1 := strconv.Itoa(len(path) - 1)

	// reset the nodes
	resetGrid(grid)
	shortestPath := 100000000
	// find all starting points
	for _, row := range grid {
		for _, node := range row {
			// if possible start node calculate path
			if node.height == 0 {
				pathTmp, ok := findPath(grid, node, goal)
				if ok && len(pathTmp) < shortestPath {
					shortestPath = len(pathTmp)
					path = pathTmp
				}
				// reset the nodes
				resetGrid(grid)
			}
		}
	}

	if test {
		fmt.Println("Pfad: ", path)
	}

	fmt.Println("Result Task 2: ", len(path)-1)
	res2 := strconv.Itoa(len(path) - 1)
	return res1, res2
}
