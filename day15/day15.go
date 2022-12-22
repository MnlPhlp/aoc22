package day15

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type point struct {
	x, y int
}

type sensor struct {
	point
	distToBeacon int
}

type grid struct {
	sensors []sensor
	beacons []point
}

func parseInput(inputFile string) grid {
	g := grid{
		sensors: make([]sensor, 0),
		beacons: make([]point, 0),
	}
	input, _ := os.ReadFile(inputFile)
	for _, l := range strings.Split(string(input), "\n") {
		if len(l) == 0 {
			continue
		}
		sensor := sensor{}
		beacon := point{}
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		sensor.distToBeacon = manhattanDistance(sensor.point, beacon)
		g.sensors = append(g.sensors, sensor)
		g.beacons = append(g.beacons, beacon)
	}
	return g
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func getCoveredPoints(g grid, line int) int {
	covered := 0
	isBeacon := make(map[int]bool)
	for _, b := range g.beacons {
		if b.y == line {
			isBeacon[b.x] = true
		}
	}
	start := 0
	end := 0
	for _, s := range g.sensors {
		distToLine := manhattanDistance(s.point, point{s.x, line})
		if distToLine > s.distToBeacon {
			continue
		}
		distOnLine := s.distToBeacon - distToLine
		if s.x-distOnLine < start {
			start = s.x - distOnLine
		}
		if s.x+distOnLine > end {
			end = s.x + distOnLine
		}
	}
	for i := start; i <= end; i++ {
		if !isBeacon[i] {
			covered++
		}
	}
	return covered
}

func findFreeSpot(g grid, min, max int) point {
	isBeacon := make(map[point]bool)
	for _, b := range g.beacons {
		isBeacon[b] = true
	}
	for y := min; y <= max; y++ {
	point:
		for x := min; x <= max; x++ {
			p := point{x, y}
			if isBeacon[p] {
				continue
			}
			for _, s := range g.sensors {
				dist := manhattanDistance(p, s.point)
				if dist <= s.distToBeacon {
					// move to next spot not covered by this sensor
					distOnLine := s.distToBeacon - manhattanDistance(s.point, point{s.x, y})
					if x-distOnLine < min && x+distOnLine > max {
						// move to next not completely covered line
						if s.x < max/2 {
							// right sight clears first
							y += distOnLine - (max - s.x) + 1
							x = max
						} else {
							// left sight clears first
							y += distOnLine - (s.x - min) + 1
							x = min
						}

					} else {
						x = s.x + distOnLine
					}
					continue point
				}
			}
			return p
		}
	}
	return point{0, 0}
}

func Solve(test bool) (string, string, time.Duration) {
	start := time.Now()

	// parse Input
	inputFile := "day15/input.txt"
	line := 2000000
	max := 4000000
	if test {
		inputFile = "day15/testInput.txt"
		line = 10
		max = 20
	}
	grid := parseInput(inputFile)
	if test {
		fmt.Printf("Beacons: %v\nSensors: %v\n", grid.beacons, grid.sensors)
	}

	// Task 1
	coveredPoints := getCoveredPoints(grid, line)
	fmt.Printf("Task 1: %v\n", coveredPoints)
	re1 := fmt.Sprintf("%v", coveredPoints)

	// Task 2
	freeSpot := findFreeSpot(grid, 0, max)
	res2 := fmt.Sprintf("%v", freeSpot.x*4000000+freeSpot.y)
	fmt.Printf("Free spot: %v\n", freeSpot)
	fmt.Println("Task 2:", res2)

	return re1, res2, time.Since(start)
}
