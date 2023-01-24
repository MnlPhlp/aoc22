package day23

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mnlphlp/aoc22/util"
)

type Grid map[util.Pos2]struct{}

func (g Grid) String() string {
	min, max := util.Pos2{X: 1 << 62, Y: 1 << 62}, util.Pos2{X: -(1 << 62), Y: -(1 << 62)}
	for p := range g {
		min = util.Pos2{X: util.Min(min.X, p.X), Y: util.Min(min.Y, p.Y)}
		max = util.Pos2{X: util.Max(max.X, p.X), Y: util.Max(max.Y, p.Y)}
	}
	s := fmt.Sprintf("min: %d,%d max: %d,%d\n", min.X, min.Y, max.X, max.Y)
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			if _, elve := g[util.Pos2{x, y}]; elve {
				s += "#"
			} else {
				s += "O"
			}
		}
		s += "\n"
	}
	return s
}

func (g Grid) Hash() string {
	positions := make([]int, 0, len(g))
	for p := range g {
		positions = append(positions, p.X<<32+p.Y)
	}
	sort.Ints(positions)
	return fmt.Sprint(positions)
}

var occupied struct{}

const (
	// weird order copied from task
	North = iota
	South
	West
	East
)

var Directions_N = []util.Pos2{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1}}
var Directions_E = []util.Pos2{{X: 1, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: 1}}
var Directions_S = []util.Pos2{{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}}
var Directions_W = []util.Pos2{{X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1}}

var DirectionGroups = [][]util.Pos2{
	Directions_N,
	Directions_S,
	Directions_W,
	Directions_E,
}

var Directions = []util.Pos2{
	{0, -1}, // North
	{0, 1},  // South
	{-1, 0}, // West
	{1, 0},  // East
}

func canMove(e util.Pos2, elves Grid) bool {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				// skip own field
				continue
			}
			if _, found := elves[e.Add(util.Pos2{x, y})]; found {
				return true
			}
		}
	}
	return false
}

func nextPos(e util.Pos2, elves Grid, startDir int) (util.Pos2, bool) {
	// check all 4 directions
	for i := 0; i < 4; i++ {
		dir := (startDir + i) % 4
		ok := true
		for _, d := range DirectionGroups[dir] {
			if _, found := elves[e.Add(d)]; found {
				ok = false
			}
		}
		if ok {
			return e.Add(Directions[dir]), true
		}
	}
	return util.Pos2{}, false
}

func move(elves Grid, rounds int, startDir int, debug bool) (Grid, int) {
	if debug {
		fmt.Println(elves)
	}
	lastHash := ""
	for i := 0; i < rounds || rounds == 0; i++ {
		hash := elves.Hash()
		if hash == lastHash {
			return elves, i
		}
		lastHash = hash
		proposed := make(map[util.Pos2]util.Pos2)
		conflicts := make([]util.Pos2, 0)
		for e := range elves {
			if !canMove(e, elves) {
				continue
			}
			next, nextFound := nextPos(e, elves, startDir)
			if nextFound {
				if _, conflict := proposed[next]; conflict {
					// do not move and mark other move for deletion
					conflicts = append(conflicts, next)
				} else {
					proposed[next] = e
				}
			}
		}
		for _, conflict := range conflicts {
			delete(proposed, conflict)
		}
		for _, old := range proposed {
			delete(elves, old)
		}
		for new := range proposed {
			elves[new] = occupied
		}
		if debug {
			fmt.Printf("Round %d: %d\n", i+1, countEmpty(elves))
			fmt.Println(elves)
		}
		startDir = (startDir + 1) % 4
	}
	return elves, rounds
}

func countEmpty(g Grid) int {
	min, max := util.Pos2{X: 1 << 62, Y: 1 << 62}, util.Pos2{X: -(1 << 62), Y: -(1 << 62)}
	for p := range g {
		min = util.Pos2{X: util.Min(min.X, p.X), Y: util.Min(min.Y, p.Y)}
		max = util.Pos2{X: util.Max(max.X, p.X), Y: util.Max(max.Y, p.Y)}
	}
	return ((max.X - min.X + 1) * (max.Y - min.Y + 1)) - len(g)
}

func part1(grid Grid, debug bool) int {
	grid, _ = move(grid, 10, North, debug)
	return countEmpty(grid)
}

func part2(grid Grid, startDir int, debug bool) int {
	_, lastMove := move(grid, 0, startDir, debug)
	return lastMove
}

func parseInput(input string) Grid {
	g := Grid{}
	for y, l := range strings.Split(input, "\n") {
		for x, c := range l {
			if c == '#' {
				g[util.Pos2{X: x, Y: y}] = occupied
			}
		}
	}
	return g
}

func Solve(input string, debug bool, task int) (string, string) {
	res1, res2 := 0, 0
	grid := parseInput(input)
	if task != 2 {
		res1 = part1(grid, debug)
	}
	if task != 1 {
		// if part one already ran continue from that state
		startDir := North
		if task == 0 {
			startDir = (startDir + 10) % 4
			res2 = 10
		}
		res2 += part2(grid, startDir, debug)
	}
	return strconv.Itoa(res1), strconv.Itoa(res2)
}
