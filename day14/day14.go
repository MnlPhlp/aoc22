package day14

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}
type cave struct {
	grid  map[point]byte
	min   point
	max   point
	floor int
}

func (c cave) String() string {
	var sb strings.Builder
	for y := c.min.y; y <= c.max.y; y++ {
		for x := c.min.x - 1; x <= c.max.x+1; x++ {
			if p, ok := c.grid[point{x, y}]; ok {
				sb.WriteByte(p)
			} else if c.floor != 0 && y == c.floor {
				sb.WriteByte('~')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func line(a, b int) []int {
	if a > b {
		return line(b, a)
	}
	res := make([]int, b-a+1)
	for i := range res {
		res[i] = a + i
	}
	return res
}

func parseInput(inputFile string) cave {
	input, _ := os.ReadFile(inputFile)
	cave := cave{
		grid: make(map[point]byte),
		min:  point{500, 0},
		max:  point{500, 0},
	}
	for _, path := range strings.Split(string(input), "\n") {
		lastPoint := point{-1, -1}
		for _, str := range strings.Split(path, " -> ") {
			var p point
			fmt.Sscanf(str, "%d,%d", &p.x, &p.y)
			cave.grid[p] = '#'
			if lastPoint.x != -1 {
				// mark all points between lastPoint and p as visited
				if lastPoint.x == p.x {
					for _, i := range line(lastPoint.y, p.y) {
						cave.grid[point{lastPoint.x, i}] = '#'
					}
				} else {
					for _, i := range line(lastPoint.x, p.x) {
						cave.grid[point{i, lastPoint.y}] = '#'
					}
				}
			}
			lastPoint = p
			if p.y < cave.min.y {
				cave.min.y = p.y
			}
			if p.y > cave.max.y {
				cave.max.y = p.y
			}
		}
	}
	return cave
}

func dropSand(cave *cave, p point) int {
	drops := 0
	for {
		// drop particle
		pos := p
		for {
			// return if falling into void
			if pos.y > cave.max.y {
				return drops
			}
			// stop at floor
			if cave.floor != 0 && pos.y == cave.floor-1 {
				goto stop
			}
			// normal drop
			if _, ok := cave.grid[point{pos.x, pos.y + 1}]; !ok {
				// drop down
				pos.y++
			} else if _, ok := cave.grid[point{pos.x - 1, pos.y + 1}]; !ok {
				// drop down and left
				pos.y++
				pos.x--
			} else if _, ok := cave.grid[point{pos.x + 1, pos.y + 1}]; !ok {
				// drop down and right
				pos.y++
				pos.x++
			} else {
				goto stop
			}
			continue
		stop:
			drops++
			cave.grid[pos] = 'o'
			if pos.x < cave.min.x {
				cave.min.x = pos.x
			} else if pos.x > cave.max.x {
				cave.max.x = pos.x
			}
			if pos == p {
				return drops
			}
			break
		}
	}
}

func Solve(test bool) (string, string) {
	inputFile := "day14/input.txt"
	if test {
		inputFile = "day14/testInput.txt"
	}
	cave := parseInput(inputFile)
	if test {
		fmt.Println(cave)
	}

	// Task 1
	maxDrops := dropSand(&cave, point{500, 0})
	if test {
		fmt.Println(cave)
	}
	fmt.Printf("Task 1: %d\n", maxDrops)
	res1 := strconv.Itoa(maxDrops)

	// Task 2
	// remove sand from cave
	for k, v := range cave.grid {
		if v == 'o' {
			delete(cave.grid, k)
		}
	}
	// add floor to cave
	cave.floor = cave.max.y + 2
	cave.max.y = cave.floor
	if test {
		fmt.Println(cave)
	}
	// drop sand
	maxDrops = dropSand(&cave, point{500, 0})
	if test {
		fmt.Println(cave)
	}
	fmt.Printf("Task 2: %d\n", maxDrops)
	res2 := strconv.Itoa(maxDrops)

	return res1, res2
}
