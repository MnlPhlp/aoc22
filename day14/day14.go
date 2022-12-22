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
			if p.y < cave.min.y {
				cave.min.y = p.y
			}
			if p.y > cave.max.y {
				cave.max.y = p.y
			}
		}
	}
	return cave
}

func dropSand(cave *cave, p point, r *recorder, record bool) int {
	drops := 0
	i := 0
	touchingFloor := false
	for {
		// drop particle
		pos := p
		for {
			i++
			if record {
				cave.grid[pos] = 'o'
				r.AddFrame(*cave, i, touchingFloor)
				delete(cave.grid, pos)
			}
			// return if falling into void
			if pos.y > cave.max.y {
				return drops
			}
			// stop at floor
			if cave.floor != 0 && pos.y == cave.floor-1 {
				touchingFloor = true
				goto stop
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
				goto stop
			}
			continue
		stop:
			drops++
			cave.grid[pos] = 'o'
			if pos.x < cave.min.x {
				cave.min.x = pos.x
			} else if pos.x > cave.max.x {
				cave.max.x = pos.x
			}
			if pos == p {
				return drops
			}
			break
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
	maxDrops := dropSand(&cave, point{500, 0}, &recorder{}, false)
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
	r := NewRecorder("day14/frames", cave.min.y, cave.max.y, cave.floor)
	maxDrops = dropSand(&cave, point{500, 0}, r, test)
	r.DrawFrames()
	if test {
		fmt.Println(cave)
	}
	fmt.Printf("Task 2: %d\n", maxDrops)
	res2 := strconv.Itoa(maxDrops)

	return res1, res2, time.Since(start)
}

type recorder struct {
	frameDir string
	frames   []map[point]byte
	minx     int
	maxx     int
	miny     int
	maxy     int
	floor    int
}

func NewRecorder(frameDir string, miny, maxy, floor int) *recorder {
	return &recorder{
		frameDir: frameDir,
		frames:   make([]map[point]byte, 0),
		minx:     500,
		maxx:     500,
		miny:     miny,
		maxy:     maxy,
		floor:    floor,
	}
}

func (r *recorder) AddFrame(cave cave, i int, touchingFloor bool) {
	if touchingFloor && i%100 != 0 {
		return
	}
	frame := make(map[point]byte)
	for k, v := range cave.grid {
		frame[k] = v
	}
	r.frames = append(r.frames, frame)
	if cave.min.x < r.minx {
		r.minx = cave.min.x
	}
	if cave.max.x > r.maxx {
		r.maxx = cave.max.x
	}
}

func (r recorder) DrawFrames() {
	colors := map[byte]color.Color{
		'#': color.RGBA{0, 0, 0, 255},
		'o': color.RGBA{255, 255, 0, 255},
	}

	for i, frame := range r.frames {
		frameFile := fmt.Sprintf("%s/frame%05d.png", r.frameDir, i)
		img := image.NewRGBA(image.Rect(r.minx-1, r.miny, r.maxx+1, r.maxy))
		for x := r.minx - 1; x <= r.maxx+1; x++ {
			for y := r.miny; y <= r.maxy; y++ {
				if v, ok := frame[point{x, y}]; ok {
					img.Set(x, y, colors[v])
				} else if r.floor != 0 && y == r.floor {
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
}
