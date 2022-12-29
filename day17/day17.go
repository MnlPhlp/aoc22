package day17

import (
	"fmt"
	"sort"

	"github.com/mnlphlp/aoc22/util"
)

// coordinate system starts at lower left corner to make it easier to extend it to the top

type rock struct {
	// bottom left corner of a rectangle around the rock
	Pos util.Pos2
	// relative position of top right corner of the rectangle
	TopRight util.Pos2
	// relative positions of all filled cells
	Shape []util.Pos2
}

var rocks = []rock{
	// horizontal line
	{util.Pos2{0, 0}, util.Pos2{3, 0}, []util.Pos2{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	// plus shape
	{util.Pos2{0, 0}, util.Pos2{2, 2}, []util.Pos2{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}},
	// inverted L shape
	{util.Pos2{0, 0}, util.Pos2{2, 2}, []util.Pos2{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}},
	// vertical line
	{util.Pos2{0, 0}, util.Pos2{0, 3}, []util.Pos2{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
	// cube
	{util.Pos2{0, 0}, util.Pos2{1, 1}, []util.Pos2{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
}

func (r rock) Draw() {
	filled := make(map[util.Pos2]bool)
	for _, pos := range r.Shape {
		filled[util.Pos2{r.Pos.X + pos.X, r.Pos.Y + pos.Y}] = true
	}
	for y := r.Pos.Y + r.TopRight.Y; y >= r.Pos.Y; y-- {
		for x := r.Pos.X; x <= r.Pos.X+r.TopRight.X; x++ {
			if filled[util.Pos2{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseInput(input string) []util.Move {
	moves := make([]util.Move, 0)
	for _, c := range input {
		if c == '<' {
			moves = append(moves, util.Move{X: -1, Y: 0})
		} else if c == '>' {
			moves = append(moves, util.Move{X: 1, Y: 0})
		}
	}
	return moves
}

type cacheState struct {
	drop     int
	maxFloor int
}

func hashState(filled map[util.Pos2]bool, rockType int, move int, maxHeight int) string {
	minFloor := maxHeight
	for pos := range filled {
		if pos.Y < minFloor {
			minFloor = pos.Y
		}
	}
	normalizedFilled := make([]util.Pos2, 0)
	for pos := range filled {
		normalizedFilled = append(normalizedFilled, util.Pos2{pos.X, pos.Y - minFloor})
	}
	sort.Slice(normalizedFilled, func(i, j int) bool {
		if normalizedFilled[i].Y == normalizedFilled[j].Y {
			return normalizedFilled[i].X < normalizedFilled[j].X
		} else {
			return normalizedFilled[i].Y < normalizedFilled[j].Y
		}
	})
	return fmt.Sprintf("%v|%d|%d", normalizedFilled, rockType, move)
}

func simulateDrops(input []util.Move, test bool, drops int) string {
	rockType := -1
	move := -1
	topRock := rock{}
	maxFloor := 0
	filled := make(map[util.Pos2]bool)
	cache := make(map[string]cacheState)
	minFloor := 0
	for x := 0; x < 7; x++ {
		filled[util.Pos2{x, 0}] = true
	}
	for i := 0; i < drops; i++ {
		// place new rock
		rockType = (rockType + 1) % len(rocks)
		topRock = rocks[rockType]
		topRock.Pos.Y = maxFloor + 4
		topRock.Pos.X = 2
	drop:
		for {
			// move rock by hot air
			move = (move + 1) % len(input)
			oldPos := topRock.Pos
			topRock.Pos.Move(input[move])
			// limit to walls
			if topRock.Pos.X < 0 || topRock.Pos.X+topRock.TopRight.X > 6 {
				// rock hits a wall
				topRock.Pos = oldPos
			} else {
				// check if rock hits another rock
				for _, pos := range topRock.Shape {
					if filled[util.Pos2{topRock.Pos.X + pos.X, topRock.Pos.Y + pos.Y}] {
						// rock hits another rock
						topRock.Pos = oldPos
					}
				}
			}
			// rock falls down
			topRock.Pos.Y--
			// check if rock hits the floor
			if topRock.Pos.Y <= maxFloor {
				for _, pos := range topRock.Shape {
					if filled[topRock.Pos.Add(pos)] {
						// rock hits the floor
						topRock.Pos.Y++
						break drop
					}
				}
			}
		}
		// rock hits the floor
		for _, pos := range topRock.Shape {
			x, y := topRock.Pos.X+pos.X, topRock.Pos.Y+pos.Y
			filled[util.Pos2{x, y}] = true
			// check if closed floor and remove everything below
			newFloor := true
			for x1 := 0; x1 < 7; x1++ {
				if !filled[util.Pos2{x1, y}] && !filled[util.Pos2{x1, y + 1}] && !filled[util.Pos2{x1, y - 1}] {
					newFloor = false
					break
				}
			}
			if newFloor {
				// remove everything below
				for y1 := minFloor; y1 < y-1; y1++ {
					for x1 := 0; x1 < 7; x1++ {
						delete(filled, util.Pos2{x1, y1})
					}
				}
				minFloor = y - 1
			}
			if y > maxFloor {
				maxFloor = y
			}
		}
		// check if cache can be used
		hash := hashState(filled, rockType, move, maxFloor)
		if cached, ok := cache[hash]; ok {
			// cache hit
			stepsGain := i - cached.drop
			heightGain := maxFloor - cached.maxFloor
			count := (drops - i) / stepsGain
			if count > 0 {
				stepsGain *= count
				heightGain *= count
				i += stepsGain
				maxFloor += heightGain
				minFloor += heightGain
				keys := make([]util.Pos2, 0, len(filled))
				for p := range filled {
					keys = append(keys, p)
				}
				for _, k := range keys {
					delete(filled, k)
					filled[util.Pos2{k.X, k.Y + heightGain}] = true
				}
			}
		} else {
			// cache miss
			cache[hash] = cacheState{
				drop:     i,
				maxFloor: maxFloor,
			}
		}

		if test && i < 10 {
			fmt.Printf("Drop %d: maxFloor: %d\n", i, maxFloor)
		}
	}
	return fmt.Sprint(maxFloor)
}

func Solve(inputStr string, test bool, task int) (string, string) {
	input := parseInput(inputStr)
	if test {
		fmt.Println("Rocks: ")
		for _, r := range rocks {
			r.Draw()
			fmt.Println()
		}
	}
	res1, res2 := "", ""
	if task != 2 {
		res1 = simulateDrops(input, test, 2022)
		fmt.Println("Result 1: ", res1)
	}
	if task != 1 {
		res2 = simulateDrops(input, test, 1000000000000)
		fmt.Println("Result 2: ", res2)
	}
	return res1, res2
}
