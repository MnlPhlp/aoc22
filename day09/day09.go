package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vec2 struct {
	x, y int
}

func (v *vec2) move(m vec2) {
	v.x += m.x
	v.y += m.y
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func (v1 vec2) touching(v2 vec2) bool {
	return abs(v1.x-v2.x) <= 1 && abs(v1.y-v2.y) <= 1
}
func (v1 *vec2) follow(v2 vec2) {
	dir := vec2{v2.x - v1.x, v2.y - v1.y}
	if dir.x != 0 {
		v1.x += dir.x / abs(dir.x)
	}
	if dir.y != 0 {
		v1.y += dir.y / abs(dir.y)
	}
}

func NewVec(dir byte) vec2 {
	switch dir {
	case 'U':
		return vec2{0, 1}
	case 'D':
		return vec2{0, -1}
	case 'R':
		return vec2{1, 0}
	case 'L':
		return vec2{-1, 0}
	}
	return vec2{0, 0}
}

type move struct {
	direction vec2
	distance  int
}

func parseMoves(input string) []move {
	lines := strings.Split(input, "\n")
	moves := make([]move, len(lines))
	var err error
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		moves[i].direction = NewVec(line[0])
		moves[i].distance, err = strconv.Atoi(line[2:])
		if err != nil {
			panic(err)
		}
	}
	return moves
}

func printKnots(recording [][]vec2) {
	var minx, miny, maxx, maxy int
	for _, knots := range recording {
		for _, k := range knots {
			if k.x < minx {
				minx = k.x
			}
			if k.x > maxx {
				maxx = k.x
			}
			if k.y < miny {
				miny = k.y
			}
			if k.y > maxy {
				maxy = k.y
			}
		}
	}
	for _, knots := range recording {
		for y := maxy; y >= miny; y-- {
			for x := minx; x <= maxx; x++ {
				found := -1
				for i, k := range knots {
					if k.x == x && k.y == y {
						found = i
						break
					}
				}
				if found >= 0 {
					fmt.Printf("%d", found)
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func makeMoves(moves []move, knotCount int, test bool) map[vec2]bool {
	visited := make(map[vec2]bool)
	knots := make([]vec2, knotCount)
	recording := make([][]vec2, 0)
	visited[vec2{}] = true
	for _, m := range moves {
		for rep := 0; rep < m.distance; rep++ {
			// move head
			knots[0].move(m.direction)
			// move following knots
			prev := knots[0]
			for i := 1; i < knotCount; i++ {
				for !knots[i].touching(prev) {
					knots[i].follow(prev)
					if i == knotCount-1 {
						visited[knots[i]] = true
					}
				}
				prev = knots[i]
			}
			if test {
				recording = append(recording, make([]vec2, knotCount))
				copy(recording[len(recording)-1], knots)
			}
		}
	}
	if test {
		printKnots(recording)
	}
	return visited
}

func main() {
	test := flag.Bool("test", false, "Run test")
	flag.Parse()
	var input []byte
	if *test {
		input, _ = os.ReadFile("testInput.txt")
	} else {
		input, _ = os.ReadFile("input.txt")
	}
	moves := parseMoves(string(input))
	fmt.Printf("Doing %d moves with 2 knots\n", len(moves))
	visited := makeMoves(moves, 2, *test)
	fmt.Printf("Visited %d locations\n", len(visited))
	fmt.Printf("Doing %d moves with 10 knots\n", len(moves))
	visited = makeMoves(moves, 10, *test)
	fmt.Printf("Visited %d locations\n", len(visited))
}
