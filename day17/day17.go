package day17

import (
	"fmt"
	"os"

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
	{util.Pos{0, 0}, util.Pos{0, 2}, []util.Pos{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
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

func parseInput(file string) []util.Move {
	input, _ := os.ReadFile(file)
	moves := make([]util.Move, len(input))
	for i, c := range input {
		if c == '<' {
			moves[i] = util.Move{X: -1, Y: 0}
		} else if c == '>' {
			moves[i] = util.Move{X: 1, Y: 0}
		}
	}
	return moves
}

func task1(input []util.Move, test bool, drops int) string {
	rockType := 0
	topRock := rocks[rockType]
	topRock.Pos.Y = 3
	floor := [7]int{}
	for i := 0; i < drops; i++ {
	drop:
		for {
			// move rock by hot air
			topRock.Pos.Move(input[i%len(input)])
			// limit to walls
			if topRock.Pos.X < 0 {
				topRock.Pos.X = 0
			} else if topRock.Pos.X+topRock.TopRight.X > 6 {
				topRock.Pos.X = 6 - topRock.TopRight.X
			}
			// rock falls down
			topRock.Pos.Y--
			// check if rock hits the floor
			for _, pos := range topRock.Shape {
				if topRock.Pos.Y+pos.Y == floor[topRock.Pos.X+pos.X] {
					// rock hits the floor
					break drop
				}
			}
		}
		// rock hits the floor
		for _, pos := range topRock.Shape {
			x, y := topRock.Pos.X+pos.X, topRock.Pos.Y+pos.Y+1
			if y > floor[x] {
				floor[x] = y
			}
		}
		// place new rock
		rockType = (rockType + 1) % len(rocks)
		topRock = rocks[rockType]
	}
	return fmt.Sprint(floor)
}

func task2(input []util.Move, test bool) string {
	return ""
}

func Solve(test bool, tasks int) (string, string) {
	inputFile := "day17/input.txt"
	if test {
		inputFile = "day17/testInput.txt"
	}
	input := parseInput(inputFile)
	if test {
		fmt.Println("Rocks: ")
		for _, r := range rocks {
			r.Draw()
			fmt.Println()
		}
		fmt.Println("Moves:")
		fmt.Println(input)
	}
	res1 := task1(input, test, 2022)
	res2 := task2(input, test)
	return res1, res2
}
