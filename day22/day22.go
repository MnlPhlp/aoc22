package day22

import (
	"fmt"
	"strings"

	"github.com/mnlphlp/aoc22/util"
)

const (
	RIGHT = iota
	DOWN
	LEFT
	UP

	EMPTY
	OPEN
	BLOCKED
)

type move struct {
	rotLeft bool
	steps   int
}

func parseInput(input string, debug bool) ([][]byte, []move) {
	grid := [][]byte{}
	path := []move{}
	gridStr := strings.Split(input, "\n\n")[0]
	pathStr := strings.Split(input, "\n\n")[1]
	if debug {
		fmt.Printf("Grid:\n%s\n", gridStr)
		fmt.Println("Path:", pathStr)
	}
	// parse grid and add border of empty fields for wrap logic
	maxLen := 0
	grid = append(grid, []byte{EMPTY})
	for y, line := range strings.Split(gridStr, "\n") {
		y++ // skip empty border
		grid = append(grid, []byte{EMPTY})
		for _, c := range line {
			switch c {
			case '#':
				grid[y] = append(grid[y], BLOCKED)
			case '.':
				grid[y] = append(grid[y], OPEN)
			default:
				grid[y] = append(grid[y], EMPTY)
			}
		}
		grid[y] = append(grid[y], EMPTY)
		if len(grid[y]) > maxLen {
			maxLen = len(grid[y])
		}
	}
	grid = append(grid, []byte{EMPTY})
	// fill all rows to same length
	for y, row := range grid {
		for i := len(row); i < maxLen; i++ {
			grid[y] = append(grid[y], EMPTY)
		}
	}

	// parse path
	numStr := ""
	for _, c := range pathStr {
		if c >= '0' && c <= '9' {
			numStr += string(c)
		} else {
			if numStr == "" {
				panic("No number before character in path: " + string(c))
			} else if c == 'L' {
				path = append(path, move{rotLeft: true, steps: util.ParseInt(numStr)})
			} else if c == 'R' {
				path = append(path, move{rotLeft: false, steps: util.ParseInt(numStr)})
			} else {
				panic("Unknown character in path: " + string(c))
			}
			numStr = ""
		}
	}
	if numStr != "" {
		path = append(path, move{steps: util.ParseInt(numStr)})
	}

	return grid, path
}

func calcPath(grid [][]byte, path []move, debug bool) []util.Pos3 {
	moved := []util.Pos3{}
	pos := util.Pos2{0, 0}
	dir := RIGHT
	// find first open tile of first row
outer:
	for y, row := range grid {
		for x, c := range row {
			if c == OPEN {
				pos = util.Pos2{x, y}
				break outer
			}
		}
	}
	if debug {
		fmt.Println("Start pos:", pos)
		fmt.Println("Start dir:", dir)
		steps := 0
		for _, m := range path {
			steps += m.steps
		}
		fmt.Println("total steps:", steps)
	}
	moved = append(moved, util.Pos3{pos.X, pos.Y, dir})
	// follow path
	for i, m := range path {
		move := util.Pos2{0, 0}
		switch dir {
		case RIGHT:
			move = util.Pos2{1, 0}
		case DOWN:
			move = util.Pos2{0, 1}
		case LEFT:
			move = util.Pos2{-1, 0}
		case UP:
			move = util.Pos2{0, -1}
		}
		for i := 0; i < m.steps; i++ {
			newPos := pos.Add(move)
			// stop if blocked
			if grid[newPos.Y][newPos.X] == BLOCKED {
				break
			}
			// wrap around
			if grid[newPos.Y][newPos.X] == EMPTY {
				switch dir {
				case RIGHT:
					newPos.X = 0
				case DOWN:
					newPos.Y = 0
				case LEFT:
					newPos.X = len(grid[0]) - 1
				case UP:
					newPos.Y = len(grid) - 1
				}
				for grid[newPos.Y][newPos.X] == EMPTY {
					newPos = newPos.Add(move)
				}
				if grid[newPos.Y][newPos.X] == BLOCKED {
					break
				}
			}
			// actually move
			pos = newPos
			if debug {
				moved = append(moved, util.Pos3{pos.X, pos.Y, dir})
			}
		}
		// do rotation if not in last move
		if i != len(path)-1 {
			if m.rotLeft {
				dir = (dir + 3) % 4
			} else {
				dir = (dir + 1) % 4
			}
		}
		moved = append(moved, util.Pos3{pos.X, pos.Y, dir})
	}
	return moved
}

func printPath(grid [][]byte, path []util.Pos3) {
	visited := make(map[util.Pos2]int)
	for _, p := range path {
		visited[util.Pos2{p.X, p.Y}] = p.Z
	}
	chars := []string{">", "v", "<", "^"}
	for y, row := range grid {
		for x, c := range row {
			if c == BLOCKED {
				fmt.Print("#")
			} else if c == EMPTY {
				fmt.Print(" ")
			} else {
				p := util.Pos2{x, y}
				if char, ok := visited[p]; ok {
					fmt.Print(chars[char])
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func part1(grid [][]byte, moves []move, debug bool) int {
	path := calcPath(grid, moves, debug)
	pos := path[len(path)-1]
	ret := 1000*pos.Y + 4*pos.X + pos.Z
	if debug {
		printPath(grid, path)
		fmt.Printf("Result 1: Row: %d, Col: %d, Dir: %d, Code: %d", pos.Y, pos.X, pos.Z, ret)
	}
	return ret
}

func Solve(input string, debug bool, task int) (string, string) {
	res1, res2 := 0, 0
	grid, path := parseInput(input, debug)
	if debug {
		fmt.Println("Grid:", grid)
		fmt.Println("Path:", path)
	}
	if task != 2 {
		res1 = part1(grid, path, debug)
	}
	if task != 1 {

	}
	return fmt.Sprint(res1), fmt.Sprint(res2)
}
