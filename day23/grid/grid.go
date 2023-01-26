package grid

import (
	"strings"
	"time"

	"github.com/mnlphlp/aoc22/util"
)

const TIMING_ACTIVE = false

type Stat struct {
	Time  time.Duration
	Calls int
}

func (s *Stat) Add(t time.Duration) {
	s.Calls++
	s.Time += t
}

var Timing struct {
	Insert, Contains, CanMove, NextMove, Hash, Remove Stat
}

type Grid struct {
	field   [][]bool
	offsetX int
	offsetY int
}

func (g Grid) Contains(p util.Pos2) bool {
	if TIMING_ACTIVE {
		start := time.Now()
		defer func() { Timing.Contains.Add(time.Since(start)) }()
	}
	p.X += g.offsetX
	p.Y += g.offsetY
	return p.X >= 0 && p.Y >= 0 && p.X < len(g.field) && p.Y < len(g.field[p.X]) && g.field[p.X][p.Y]
}

func (g Grid) Insert(p util.Pos2) Grid {
	if TIMING_ACTIVE {
		start := time.Now()
		defer func() { Timing.Insert.Add(time.Since(start)) }()
	}
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
	if TIMING_ACTIVE {
		start := time.Now()
		defer func() { Timing.Remove.Add(time.Since(start)) }()
	}
	g.field[p.X+g.offsetX][p.Y+g.offsetY] = false
	return g
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
			if len(g.field[x]) <= y || !g.field[x][y] {
				s += "."
			} else {
				s += "#"
			}
		}
		s += "\n"
	}
	return s
}

func (g Grid) ForEach(f func(util.Pos2)) {
	x := 0
	for x = 0; x < len(g.field); x++ {
		for y := 0; y < len(g.field[x]); y++ {
			if g.field[x][y] {
				f(util.Pos2{x - g.offsetX, y - g.offsetY})
			}
		}
	}
}

type GridHash []int

func (h1 GridHash) Equals(h2 GridHash) bool {
	if len(h1) != len(h2) {
		return false
	}
	for i := range h1 {
		if h1[i] != h2[i] {
			return false
		}
	}
	return true
}

func (g Grid) Hash() GridHash {
	if TIMING_ACTIVE {
		start := time.Now()
		defer func() { Timing.Hash.Add(time.Since(start)) }()
	}
	positions := make([]int, 0)
	x := 0
	for x = 0; x < len(g.field); x++ {
		for y := 0; y < len(g.field[x]); y++ {
			if g.field[x][y] {
				positions = append(positions, x<<32+y)
			}
		}
	}
	return positions
}

func (g Grid) HasNeighbor(p util.Pos2) bool {
	if TIMING_ACTIVE {
		start := time.Now()
		defer func() { Timing.CanMove.Add(time.Since(start)) }()
	}
	p.X += g.offsetX
	p.Y += g.offsetY
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			// skip own element and invalid cases
			if x == 0 && y == 0 || p.X+x < 0 || p.Y+y < 0 || p.X+x >= len(g.field) || p.Y+y >= len(g.field[p.X+x]) {
				continue
			}
			if g.field[p.X+x][p.Y+y] {
				return true
			}
		}
	}
	return false
}

var Directions = []util.Pos2{
	{0, -1}, // North
	{0, 1},  // South
	{-1, 0}, // West
	{1, 0},  // East
}

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

func (g Grid) NextPos(p util.Pos2, startDir int) (util.Pos2, bool) {
	if TIMING_ACTIVE {
		start := time.Now()
		defer func() { Timing.NextMove.Add(time.Since(start)) }()
	}
	p.X += g.offsetX
	p.Y += g.offsetY
	// check all 4 directions
	for i := 0; i < 4; i++ {
		dir := (startDir + i) % 4
		ok := true
		for _, d := range DirectionGroups[dir] {
			// skip invalid cases
			if p.X+d.X < 0 || p.Y+d.Y < 0 || p.X+d.X >= len(g.field) || p.Y+d.Y >= len(g.field[p.X+d.X]) {
				continue
			}
			if g.field[p.X+d.X][p.Y+d.Y] {
				ok = false
			}
		}
		if ok {
			p.X = p.X + Directions[dir].X - g.offsetX
			p.Y = p.Y + Directions[dir].Y - g.offsetY
			return p, true
		}
	}
	return util.Pos2{}, false
}

func ParseInput(input string) Grid {
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
