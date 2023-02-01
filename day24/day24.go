package day24

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mnlphlp/aoc22/util"
)

type node struct {
	util.Pos2
	hasWall, hasBlizzard bool
	f, g, h              int
	parent               *node
}

func (n *node) calc(goal *node) {
	n.g = n.parent.g + 1
	n.h = util.Abs(n.X-goal.X) + util.Abs(n.Y-goal.Y)
	n.f = n.g + n.h
}

func (n *node) isIn(nodes []*node) bool {
	return util.Contains(nodes, n)
}

func (n *node) clone() *node {
	return &node{
		Pos2:        n.Pos2,
		hasWall:     n.hasWall,
		hasBlizzard: n.hasBlizzard,
		f:           n.f,
		g:           n.g,
		h:           n.h,
		parent:      n.parent,
	}
}

var blizzardSigns = []rune{'^', '>', 'v', '<'}

type Grid struct {
	field              [][]*node
	blizzards          []util.Pos2
	blizzardDirections []int
	start, end         *node
}

func parseInput(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := Grid{
		field:              make([][]*node, len(lines)),
		blizzards:          make([]util.Pos2, 0),
		blizzardDirections: make([]int, 0),
	}
	for y, line := range lines {
		grid.field[y] = make([]*node, len(line))
		for x, c := range line {
			grid.field[y][x] = &node{
				Pos2:        util.Pos2{X: x, Y: y},
				hasWall:     c == '#',
				hasBlizzard: util.Contains(blizzardSigns, c),
			}
			if grid.field[y][x].hasBlizzard {
				grid.blizzards = append(grid.blizzards, util.Pos2{X: x, Y: y})
				grid.blizzardDirections = append(grid.blizzardDirections, util.IndexOf(blizzardSigns, c))
			}
		}
	}
	grid.start = grid.field[0][1]
	grid.end = grid.field[len(grid.field)-1][len(grid.field[0])-2]
	return grid
}

func getNeighbors(grid [][]*node, current *node, closed map[*node]bool) []*node {
	neighbors := make([]*node, 0)
	for _, move := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		newX, newY := current.X+move[0], current.X+move[1]
		if newX < 0 || newX >= len(grid[0]) || newY < 0 || newY >= len(grid) {
			continue
		}
		newPos := grid[newY][newX]
		if closed[newPos] {
			continue
		}
		if newPos.hasWall || newPos.hasBlizzard {
			continue
		}
		newPos.parent = current
		neighbors = append(neighbors, newPos)
	}
	if len(neighbors) == 0 {
		// allow to wait
		cloned := current.clone()
		cloned.parent = current
		neighbors = append(neighbors, cloned)
	}
	return neighbors
}

func moveBlizzards(grid Grid) {
	for _, row := range grid.field {
		for _, node := range row {
			if node.hasBlizzard {
			}
		}
	}
}

func getPath(g Grid) []util.Pos2 {
	closed := make(map[*node]bool)
	open := make([]*node, 1)
	open[0] = g.start
	current := g.start
	waiting := 0
	for len(open) > 0 {
		moveBlizzards(g)
		// get node with lowest f
		min := 0
		for i := range open {
			if open[i].f < open[min].f {
				min = i
			}
		}
		// keep track of waiting time
		if open[min].Pos2 == current.Pos2 {
			waiting++
		} else {
			waiting = 0
		}
		if waiting > 100 {
			panic("Infinite loop")
		}
		current = open[min]
		open = append(open[:min], open[min+1:]...)
		closed[current] = true
		// if reached goal, return path
		if current.Pos2 == g.end.Pos2 {
			break
		}
		// get neighbors
		neighbors := getNeighbors(g.field, current, closed)
		for _, n := range neighbors {
			if closed[n] {
				continue
			}
			n.calc(g.end)
			if !n.isIn(open) {
				open = append(open, n)
			}
		}

	}
	for _, row := range g.field {
		for _, node := range row {
			parentPos := util.Pos2{X: -1, Y: -1}
			if node.parent != nil {
				parentPos = node.parent.Pos2
			}
			if node.Pos2 == parentPos {
				fmt.Printf("Self reference: %v/%v", node.Pos2, parentPos)
			}
		}
	}
	// if reached goal, return path
	if current.Pos2 == g.end.Pos2 {
		path := make([]util.Pos2, 0)
		for current != nil {
			path = append(path, util.Pos2{X: current.X, Y: current.Y})
			current = current.parent
		}
		return path
	} else {
		panic("No path found")
	}
}

func part1(grid Grid) int {
	path := getPath(grid)
	return len(path)
}

func Solve(input string, debug bool, task int) (string, string) {
	ret1, ret2 := 0, 0
	grid := parseInput(input)
	if task != 2 {
		ret1 = part1(grid)
	}

	return strconv.Itoa(ret1), strconv.Itoa(ret2)
}
