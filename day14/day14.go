package day14

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}
type cave struct {
	grid  map[point]byte
	min   point
	max   point
	floor int
}

func (c cave) String() string {
	var sb strings.Builder
	for y := c.min.y; y <= c.max.y; y++ {
		for x := c.min.x - 1; x <= c.max.x+1; x++ {
			if p, ok := c.grid[point{x, y}]; ok {
				sb.WriteByte(p)
			} else if c.floor != 0 && y == c.floor {
				sb.WriteByte('~')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func line(a, b int) []int {
	if a > b {
		return line(b, a)
	}
	res := make([]int, b-a+1)
	for i := range res {
		res[i] = a + i
	}
	return res
}

func parseInput(inputFile string) cave {
	input, _ := os.ReadFile(inputFile)
	cave := cave{
		grid: make(map[point]byte),
		min:  point{500, 0},
		max:  point{500, 0},
	}
	for _, path := range strings.Split(string(input), "\n") {
		lastPoint := point{-1, -1}
		for _, str := range strings.Split(path, " -> ") {
			var p point
			fmt.Sscanf(str, "%d,%d", &p.x, &p.y)
			cave.grid[p] = '#'
			if lastPoint.x != -1 {
				// mark all points between lastPoint and p as visited
				if lastPoint.x == p.x {
					for _, i := range line(lastPoint.y, p.y) {
						cave.grid[point{lastPoint.x, i}] = '#'
					}
				} else {
					for _, i := range line(lastPoint.x, p.x) {
						cave.grid[point{i, lastPoint.y}] = '#'
					}
				}
			}
			lastPoint = p
			if p.x < cave.min.x {
				cave.min.x = p.x
			}
			if p.y < cave.min.y {
				cave.min.y = p.y
			}
			if p.x > cave.max.x {
				cave.max.x = p.x
			}
			if p.y > cave.max.y {
				cave.max.y = p.y
			}
		}
	}
	return cave
}

func drawFrame(cave *cave, frameDir string, frame int) {
	frameFile := fmt.Sprintf("%s/frame%05d.png", frameDir, frame)
	colors := map[byte]color.Color{
		'#': color.RGBA{0, 0, 0, 255},
		'o': color.RGBA{255, 255, 0, 255},
	}
	img := image.NewRGBA(image.Rect(cave.min.x-1, cave.min.y, cave.max.x+1, cave.max.y))
	for x := cave.min.x - 1; x <= cave.max.x+1; x++ {
		for y := cave.min.y; y <= cave.max.y; y++ {
			if v, ok := cave.grid[point{x, y}]; ok {
				img.Set(x, y, colors[v])
			} else if cave.floor != 0 && y == cave.floor {
				img.Set(x, y, colors['#'])
			} else {
				img.Set(x, y, color.White)
			}
		}
	}
	f, _ := os.Create(frameFile)
	png.Encode(f, img)
	f.Close()
}

func dropSand(cave *cave, p point, frameDir string) int {
	drops := 0
	i := 0
	frame := 0
	for {
		// drop particle
		pos := p
		for {
			i++
			if len(frameDir) > 0 && i > 10000 {
				frame++
				i = 0
				cave.grid[pos] = 'o'
				drawFrame(cave, frameDir, frame)
				delete(cave.grid, pos)
			}
			// return if falling into void
			if pos.y > cave.max.y {
				return drops
			}
			// stop at floor
			if cave.floor != 0 && pos.y == cave.floor-1 {
				drops++
				cave.grid[pos] = 'o'
				if pos.x < cave.min.x {
					cave.min.x = pos.x
				} else if pos.x > cave.max.x {
					cave.max.x = pos.x
				}
				break
			}
			// normal drop
			if _, ok := cave.grid[point{pos.x, pos.y + 1}]; !ok {
				// drop down
				pos.y++
			} else if _, ok := cave.grid[point{pos.x - 1, pos.y + 1}]; !ok {
				// drop down and left
				pos.y++
				pos.x--
			} else if _, ok := cave.grid[point{pos.x + 1, pos.y + 1}]; !ok {
				// drop down and right
				pos.y++
				pos.x++
			} else {
				// stop
				drops++
				cave.grid[pos] = 'o'
				if pos == p {
					return drops
				}
				break
			}
		}
	}
}

func Solve(test bool) (string, string, time.Duration) {
	start := time.Now()
	inputFile := "day14/input.txt"
	if test {
		inputFile = "day14/testInput.txt"
	}
	cave := parseInput(inputFile)
	if test {
		fmt.Println(cave)
	}

	// Task 1
	maxDrops := dropSand(&cave, point{500, 0}, "")
	if test {
		fmt.Println(cave)
	}
	fmt.Printf("Task 1: %d\n", maxDrops)
	res1 := strconv.Itoa(maxDrops)

	// Task 2
	// remove sand from cave
	for k, v := range cave.grid {
		if v == 'o' {
			delete(cave.grid, k)
		}
	}
	// add floor to cave
	cave.floor = cave.max.y + 2
	cave.max.y = cave.floor
	if test {
		fmt.Println(cave)
	}
	// drop sand
	maxDrops = dropSand(&cave, point{500, 0}, "day14/frames")
	if test {
		fmt.Println(cave)
	}
	fmt.Printf("Task 2: %d\n", maxDrops)
	res2 := strconv.Itoa(maxDrops)

	return res1, res2, time.Since(start)
}
