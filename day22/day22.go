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
	// parse grid
	for y, line := range strings.Split(gridStr, "\n") {
		grid = append(grid, make([]byte, len(line)))
		for x, c := range line {
			switch c {
			case '#':
				grid[y][x] = BLOCKED
			case '.':
				grid[y][x] = OPEN
			default:
				grid[y][x] = EMPTY
			}
		}
		grid[y] = append(grid[y], EMPTY)
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

func doMoves(grid [][]byte, moves []move, debug bool, wrapMap map[util.Pos2]util.Pos2) (util.Pos3, []util.Pos3) {
	// find start cell
	startx, starty := 0, 0
	// find first open tile of first row
	for x, c := range grid[0] {
		if c == OPEN {
			startx = x
			break
		}
	}
	// record moves in path
	var wayPoints []util.Pos3
	pos := util.Pos2{startx, starty}
	dir := RIGHT
	var next util.Pos2
	for i, m := range moves {
		// do steps
		for j := 0; j < m.steps; j++ {
			switch dir {
			case RIGHT:
				// right
				next = findPos(grid, pos, util.Pos2{1, 0}, util.Pos2{0, pos.Y})
			case DOWN:
				// down
				next = findPos(grid, pos, util.Pos2{0, 1}, util.Pos2{pos.X, 0})
			case LEFT:
				// left
				next = findPos(grid, pos, util.Pos2{-1, 0}, util.Pos2{len(grid[pos.Y]) - 1, pos.Y})
			case UP:
				// up
				next = findPos(grid, pos, util.Pos2{0, -1}, util.Pos2{pos.X, len(grid) - 1})

			}
			if next.X == pos.X && next.Y == pos.Y {
				break
			}
			pos = next
		}
		// do rotation if not in last move
		if i != len(moves)-1 {
			if m.rotLeft {
				dir = (dir + 3) % 4
			} else {
				dir = (dir + 1) % 4
			}
		}
	}
	return util.Pos3{pos.X, pos.Y, dir}, wayPoints

}

func findPos(grid [][]byte, start, move, wrap util.Pos2) util.Pos2 {
	//  wrap around if needed
	pos := start.Add(move)
	if pos.Y < 0 || pos.Y >= len(grid) || pos.X < 0 || pos.X >= len(grid[pos.Y]) || grid[pos.Y][pos.X] == EMPTY {
		// wrap around
		pos = wrap
		// move to next  cell
		for pos.Y < 0 || pos.Y >= len(grid) || pos.X < 0 || pos.X >= len(grid[pos.Y]) || grid[pos.Y][pos.X] == EMPTY {
			pos = pos.Add(move)
		}
	}

	// check if next cell is blocked
	if grid[pos.Y][pos.X] == BLOCKED {
		return start
	}
	return pos
}

func part1(grid [][]byte, moves []move, debug bool) int {
	end, path := doMoves(grid, moves, debug, map[util.Pos2]util.Pos2{})
	if debug {
		printPath(grid, path)
	}
	return 1000*(end.Y+1) + 4*(end.X+1) + end.Z
}

func part2(grid [][]byte, moves []move, debug bool) int {
	end, path := doMoves(grid, moves, debug, map[util.Pos2]util.Pos2{})
	if debug {
		printPath(grid, path)
	}
	return 1000*(end.Y+1) + 4*(end.X+1) + end.Z
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
		res2 = part2(grid, path, debug)
	}
	return fmt.Sprint(res1), fmt.Sprint(res2)
}
