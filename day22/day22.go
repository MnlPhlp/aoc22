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
		visited[util.Pos2{X: p.X, Y: p.Y}] = p.Z
	}
	chars := []string{">", "v", "<", "^"}
	for y, row := range grid {
		for x, c := range row {
			if c == BLOCKED {
				fmt.Print("#")
			} else if c == EMPTY {
				fmt.Print(" ")
			} else {
				p := util.Pos2{X: x, Y: y}
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

func doMoves(grid [][]byte, moves []move, debug bool, cube bool) (util.Pos3, []util.Pos3) {
	// find start cell
	startX, startY := 0, 0
	// find first open tile of first row
	for x, c := range grid[0] {
		if c == OPEN {
			startX = x
			break
		}
	}
	// record moves in path
	var wayPoints []util.Pos3
	pos := util.Pos2{X: startX, Y: startY}
	dir := RIGHT
	for i, m := range moves {
		// do steps
		for j := 0; j < m.steps; j++ {
			next, nextDir := findPos(grid, pos, dir, cube)
			if next.X == pos.X && next.Y == pos.Y {
				break
			}
			pos = next
			dir = nextDir
			if debug {
				wayPoints = append(wayPoints, util.Pos3{X: pos.X, Y: pos.Y, Z: dir})
			}
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
	return util.Pos3{X: pos.X, Y: pos.Y, Z: int(dir)}, wayPoints
}

func getMove(dir int) util.Pos2 {
	switch dir {
	case RIGHT:
		return util.Pos2{X: 1, Y: 0}
	case DOWN:
		return util.Pos2{X: 0, Y: 1}
	case LEFT:
		return util.Pos2{X: -1, Y: 0}
	case UP:
		return util.Pos2{X: 0, Y: -1}
	}
	panic("Unknown direction: " + fmt.Sprint(dir))
}

func getCubeFace(pos util.Pos2) int {
	// returns number of the face of the cube
	// numbers on the input are defined as follows:
	//   0 1
	//   2
	// 3 4
	// 5
	if pos.Y < 50 {
		// to row
		if pos.X < 100 {
			// left
			return 0
		} else {
			// right
			return 1
		}
	} else if pos.Y < 100 {
		// second row
		return 2
	} else if pos.Y < 150 {
		// third row
		if pos.X < 50 {
			// left
			return 3
		} else {
			// right
			return 4
		}
	} else {
		// fourth row
		return 5
	}
}

func getWrap(grid [][]byte, pos util.Pos2, dir int, cube bool) util.Pos3 {
	if !cube {
		switch dir {
		case RIGHT:
			return util.Pos3{X: 0, Y: pos.Y, Z: dir}
		case DOWN:
			return util.Pos3{X: pos.X, Y: 0, Z: dir}
		case LEFT:
			return util.Pos3{X: len(grid[pos.Y]) - 1, Y: pos.Y, Z: dir}
		case UP:
			return util.Pos3{X: pos.X, Y: len(grid) - 1, Z: dir}
		}
		panic("Unknown direction: " + fmt.Sprint(dir))
	} else {
		face := getCubeFace(pos)
		// faces are defined as follows (for my input):
		//   0 1
		//   2
		// 3 4
		// 5
		switch dir {
		case RIGHT:
			switch face {
			case 1:
				return util.Pos3{X: 99, Y: 149 - pos.Y, Z: LEFT} // move to face 4 and rotate 180
			case 2:
				return util.Pos3{X: pos.Y + 50, Y: 49, Z: UP} // move to face 1 and rotate -90 (or 270)
			case 4:
				return util.Pos3{X: 149, Y: 49 - (pos.Y - 100), Z: LEFT} // move to face 1 and rotate 180
			case 5:
				return util.Pos3{X: pos.Y - 100, Y: 149, Z: UP} // move to face 4 and rotate -90
			}
		case DOWN:
			switch face {
			case 1:
				return util.Pos3{X: 99, Y: pos.X - 50, Z: LEFT} // move to face 2 and rotate 90
			case 4:
				return util.Pos3{X: 49, Y: pos.X + 100, Z: LEFT} // move to face 5 and rotate 90
			case 5:
				return util.Pos3{X: pos.X + 100, Y: 0, Z: DOWN} // move to face 1 and rotate 0
			}
		case LEFT:
			switch face {
			case 0:
				return util.Pos3{X: 0, Y: 149 - pos.Y, Z: RIGHT} // move to face 3 and rotate 180
			case 2:
				return util.Pos3{X: pos.Y - 50, Y: 100, Z: DOWN} // move to face 3 and rotate -90 (or 270)
			case 3:
				return util.Pos3{X: 50, Y: 49 - (pos.Y - 100), Z: RIGHT} // move to face 0 and rotate 180
			case 5:
				return util.Pos3{X: pos.Y - 100, Y: 0, Z: DOWN} // move to face 0 and rotate -90
			}
		case UP:
			switch face {
			case 0:
				return util.Pos3{X: 0, Y: pos.X + 100, Z: RIGHT} // move to face 5 and rotate 90
			case 1:
				return util.Pos3{X: pos.X - 100, Y: 199, Z: UP} // move to face 5 and rotate 0
			case 3:
				return util.Pos3{X: 50, Y: pos.X + 50, Z: RIGHT} // move to face 2 and rotate 90
			}
		}
		// no wrap
		return util.Pos3{X: pos.X, Y: pos.Y, Z: dir}
	}
}

func findPos(grid [][]byte, start util.Pos2, dirStart int, cube bool) (util.Pos2, int) {
	move := getMove(dirStart)
	//  wrap around if needed
	pos := start.Add(move)
	dir := dirStart
	if pos.Y < 0 || pos.Y >= len(grid) || pos.X < 0 || pos.X >= len(grid[pos.Y]) || grid[pos.Y][pos.X] == EMPTY {
		// wrap around
		wrap := getWrap(grid, start, dirStart, cube)
		pos.X = wrap.X
		pos.Y = wrap.Y
		dir = wrap.Z
		// move to next  cell
		for pos.Y < 0 || pos.Y >= len(grid) || pos.X < 0 || pos.X >= len(grid[pos.Y]) || grid[pos.Y][pos.X] == EMPTY {
			pos = pos.Add(move)
		}
	}

	// check if next cell is blocked
	if grid[pos.Y][pos.X] == BLOCKED {
		return start, dirStart
	}
	return pos, dir
}

func part1(grid [][]byte, moves []move, debug bool) int {
	end, path := doMoves(grid, moves, debug, false)
	if debug {
		printPath(grid, path)
	}
	return 1000*(end.Y+1) + 4*(end.X+1) + end.Z
}

func part2(grid [][]byte, moves []move, debug bool) int {
	// part two only works correctly for the real input, not the test
	// because of the wrapping in the getWrap function
	if len(grid) != 200 {
		fmt.Println("Part 2 only works for the real input, not the test input")
		return 0
	}
	end, path := doMoves(grid, moves, debug, true)
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
