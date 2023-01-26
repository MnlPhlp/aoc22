package day23

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/mnlphlp/aoc22/util"
)

var timing struct {
	insert, contains, canMove, nextMove, hash, remove, Items time.Duration
}

type Grid struct {
	field   [][]bool
	offsetX int
	offsetY int
}

func (g Grid) Contains(p util.Pos2) bool {
	start := time.Now()
	defer func() { timing.contains += time.Since(start) }()
	p.X += g.offsetX
	p.Y += g.offsetY
	return p.X >= 0 && p.Y >= 0 && p.X < len(g.field) && p.Y < len(g.field[p.X]) && g.field[p.X][p.Y]
}
func (g Grid) Insert(p util.Pos2) Grid {
	start := time.Now()
	defer func() { timing.insert += time.Since(start) }()
	if diff := -(p.X + g.offsetX); diff > 0 {
		g.offsetX += diff
		for i := 0; i < diff; i++ {
			g.field = append(g.field, []bool{})
		}
		// move values back
		for i := len(g.field) - diff - 1; i >= 0; i-- {
			g.field[i+diff] = g.field[i]
		}
		for i := 0; i < diff; i++ {
			g.field[i] = []bool{}
		}
	}
	p.X += g.offsetX
	for len(g.field) <= p.X {
		g.field = append(g.field, []bool{})
	}
	if diff := -(p.Y + g.offsetY); diff > 0 {
		g.offsetY += diff
		for x := 0; x < len(g.field); x++ {
			// extend grid
			for i := 0; i < diff; i++ {
				g.field[x] = append(g.field[x], false)
			}
			// move values back
			for i := len(g.field[x]) - diff - 1; i >= 0; i-- {
				g.field[x][i+diff] = g.field[x][i]
			}
			for i := 0; i < diff; i++ {
				g.field[x][i] = false
			}
		}
	}
	p.Y += g.offsetY
	for len(g.field[p.X]) <= p.Y {
		g.field[p.X] = append(g.field[p.X], false)
	}
	g.field[p.X][p.Y] = true
	return g
}
func (g Grid) Remove(p util.Pos2) Grid {
	start := time.Now()
	defer func() { timing.remove += time.Since(start) }()
	g.field[p.X+g.offsetX][p.Y+g.offsetY] = false
	return g
}
func (g Grid) Items() []util.Pos2 {
	start := time.Now()
	defer func() { timing.Items += time.Since(start) }()
	ret := make([]util.Pos2, 0)
	for x := 0; x < len(g.field); x++ {
		for y := 0; y < len(g.field[x]); y++ {
			if g.field[x][y] {
				ret = append(ret, util.Pos2{X: x - g.offsetX, Y: y - g.offsetY})
			}
		}
	}
	return ret
}

func (g Grid) String() string {
	maxX := len(g.field)
	maxY := 0
	for _, row := range g.field {
		if len(row) > maxY {
			maxY = len(row)
		}
	}
	s := ""
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if g.Contains(util.Pos2{x, y}) {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (g Grid) Hash() string {
	start := time.Now()
	defer func() { timing.hash += time.Since(start) }()
	items := g.Items()
	positions := make([]int, 0, len(items))
	for _, p := range items {
		positions = append(positions, p.X<<32+p.Y)
	}
	sort.Ints(positions)
	return fmt.Sprint(positions)
}

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
	start := time.Now()
	defer func() { timing.canMove += time.Since(start) }()
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				// skip own field
				continue
			}
			if elves.Contains(e.Add(util.Pos2{x, y})) {
				return true
			}
		}
	}
	return false
}

func nextPos(e util.Pos2, elves Grid, startDir int) (util.Pos2, bool) {
	start := time.Now()
	defer func() { timing.nextMove += time.Since(start) }()
	// check all 4 directions
	for i := 0; i < 4; i++ {
		dir := (startDir + i) % 4
		ok := true
		for _, d := range DirectionGroups[dir] {
			if elves.Contains(e.Add(d)) {
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
		for _, e := range elves.Items() {
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
		if debug {
			fmt.Println("Proposed:\n", proposed)
		}
		for _, old := range proposed {
			elves = elves.Remove(old)
		}
		if debug {
			fmt.Println("After Delete\n", elves)
		}
		for new := range proposed {
			elves = elves.Insert(new)
			if debug && i == 0 {
				fmt.Printf("insert %v\n%v\n", new, elves)
			}
		}
		if debug {
			fmt.Printf("Round  %d: %d\n", i+1, countEmpty(elves))
			fmt.Println(elves)
		}
		startDir = (startDir + 1) % 4
	}
	return elves, rounds
}

func countEmpty(g Grid) int {
	min, max := util.Pos2{X: 1 << 62, Y: 1 << 62}, util.Pos2{X: -(1 << 62), Y: -(1 << 62)}
	for _, p := range g.Items() {
		min = util.Pos2{X: util.Min(min.X, p.X), Y: util.Min(min.Y, p.Y)}
		max = util.Pos2{X: util.Max(max.X, p.X), Y: util.Max(max.Y, p.Y)}
	}
	return ((max.X - min.X + 1) * (max.Y - min.Y + 1)) - len(g.Items())
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
				g = g.Insert(util.Pos2{X: x, Y: y})
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
	fmt.Println("canMove: ", timing.canMove)
	fmt.Println("contains: ", timing.contains)
	fmt.Println("hash: ", timing.hash)
	fmt.Println("insert: ", timing.insert)
	fmt.Println("nextMove: ", timing.nextMove)
	fmt.Println("remove: ", timing.remove)
	fmt.Println("Items: ", timing.Items)

	return strconv.Itoa(res1), strconv.Itoa(res2)
}
