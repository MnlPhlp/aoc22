package day17

import (
	"fmt"

	"github.com/mnlphlp/aoc22/util"
)

// coordinate system starts at lower left corner to make it easier to extend it to the top

type rock struct {
	// bottom left corner of a rectangle around the rock
	Pos util.Pos
	// relative position of top right corner of the rectangle
	TopRight util.Pos
	// relative positions of all filled cells
	Shape []util.Pos
}

var rocks = []rock{
	// horizontal line
	{util.Pos{0, 0}, util.Pos{3, 0}, []util.Pos{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	// plus shape
	{util.Pos{0, 0}, util.Pos{2, 2}, []util.Pos{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}},
	// inverted L shape
	{util.Pos{0, 0}, util.Pos{2, 2}, []util.Pos{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}},
	// vertical line
	{util.Pos{0, 0}, util.Pos{0, 3}, []util.Pos{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
	// cube
	{util.Pos{0, 0}, util.Pos{1, 1}, []util.Pos{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
}

func (r rock) Draw() {
	filled := make(map[util.Pos]bool)
	for _, pos := range r.Shape {
		filled[util.Pos{r.Pos.X + pos.X, r.Pos.Y + pos.Y}] = true
	}
	for y := r.Pos.Y + r.TopRight.Y; y >= r.Pos.Y; y-- {
		for x := r.Pos.X; x <= r.Pos.X+r.TopRight.X; x++ {
			if filled[util.Pos{x, y}] {
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

func simulateDrops(input []util.Move, test bool, drops int) string {
	rockType := -1
	move := -1
	topRock := rock{}
	maxFloor := 0
	floor := [7]int{}
	filled := make(map[util.Pos]bool)
	for x := 0; x < 7; x++ {
		filled[util.Pos{x, 0}] = true
	}
	for i := 0; i < drops; i++ {
		if i%10000 == 0 {
			fmt.Printf("%f%% done. Filled size %d\n", float64(i*1000/drops)/10, len(filled))
		}
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
					if filled[util.Pos{topRock.Pos.X + pos.X, topRock.Pos.Y + pos.Y}] {
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
			filled[util.Pos{x, y}] = true
			// check if closed floor and remove everything below
			newFloor := true
			for x1 := 0; x1 < 7; x1++ {
				if !filled[util.Pos{x1, y}] {
					newFloor = false
					break
				}
			}
			if newFloor {
				// remove everything below
				for y1 := 0; y1 < y; y1++ {
					for x1 := 0; x1 < 7; x1++ {
						delete(filled, util.Pos{x1, y1})
					}
				}
			}
			// update floor
			if y > floor[x] {
				floor[x] = y
			}
			if y > maxFloor {
				maxFloor = y
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
	return res1, res2
	if task != 1 {
		res2 = simulateDrops(input, test, 1000000000000)
		fmt.Println("Result 2: ", res2)
	}
	return res1, res2
}
