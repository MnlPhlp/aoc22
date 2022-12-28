package day16

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type valve struct {
	name        string
	flow        int
	f           int
	timeToOpen  int
	open        bool
	connections []string
	id          int64
	inPath      int64
	curFlow     int
	hash        complex128
	openedCount int
}

func (v *valve) calc(time int, current valve) {
	v.inPath = current.inPath
	v.timeToOpen = current.timeToOpen + 1
	v.f = current.f + v.flow*(time-v.timeToOpen)
	v.curFlow = current.curFlow + v.flow
	v.openedCount = current.openedCount
	if v.open {
		v.inPath = v.inPath | v.id
		v.openedCount = current.openedCount + 1
	}
	open := int64(0)
	if v.open {
		open = 1
	}
	v.hash = complex(float64(v.inPath), float64(v.id+open))
}

func (v valve) String() string {
	return fmt.Sprintf("Valve: %s (%d), Flow: %d, Connections to: %v\n  Path: %d\n  curFlow: %v",
		v.name, v.id, v.flow, v.connections, v.inPath, v.curFlow)
}

func parseValves(file string) map[string]valve {
	valves := make(map[string]valve)
	connections := make(map[string]string)
	input, _ := os.ReadFile(file)
	for i, l := range strings.Split(string(input), "\n") {
		if l == "" {
			continue
		}
		parts := strings.Split(l, " ")
		connection := ""
		for _, p := range parts[9:] {
			connection += p
		}
		flow, _ := strconv.Atoi(parts[4][5 : len(parts[4])-1])
		id := int64(1) << i
		valves[parts[1]+"C"] = valve{
			name: parts[1] + "C",
			flow: 0,
			id:   id,
		}
		valves[parts[1]+"O"] = valve{
			name: parts[1] + "O",
			flow: flow,
			open: true,
			id:   id,
		}
		connections[parts[1]+"O"] = connection
		connections[parts[1]+"C"] = connection
	}
	for name, connection := range connections {
		val := valves[name]
		for _, c := range strings.Split(connection, ",") {
			val.connections = append(val.connections, c+"C")
		}
		val.connections = append(val.connections, name[:len(name)-1]+"C")
		if !valves[name].open {
			val.connections = append(val.connections, name[:len(name)-1]+"O")
		}
		valves[name] = val
	}
	return valves
}

func findMaxPressure(valves map[string]valve, time int, startPoint string, multipleWorkers bool) (int, map[int64]valve) {
	open := make([]valve, 0)
	openHash := make(map[complex128]bool)
	// start with closed version of first valve
	current := valves[startPoint+"C"]
	current.calc(time, valve{
		timeToOpen: -1,
	})
	open = append(open, current)
	bestGoal := current
	bestPaths := make(map[int64]valve)
	cache := make(map[int64]int)
	maxFlow := 0
	workingValves := 0
	for _, v := range valves {
		maxFlow += v.flow
		if v.flow > 0 {
			workingValves++
		}
	}
	// explore paths using a* like algorithm
	for len(open) > 0 {
		// choose best item of open list
		maxScore := 0
		for i, v := range open {
			if v.f > open[maxScore].f {
				maxScore = i
			}
		}
		current = open[maxScore]
		open = append(open[:maxScore], open[maxScore+1:]...)
		delete(openHash, current.hash)

		// check if better solution is found
		if current.timeToOpen == time-1 {
			overallPaths++
			if current.f > bestGoal.f {
				bestGoal = current
			}
			// check paths
			bestPath, ok := bestPaths[current.inPath]
			balanced := true
			if multipleWorkers {
				percentOpen := float32(current.openedCount) / float32(workingValves)
				if percentOpen < 0.3 || percentOpen > 0.7 {
					balanced = false
				}
			}
			if (!ok || current.f > bestPath.f) && balanced {
				bestPaths[current.inPath] = current
			}
			//}
		}

		for _, v := range current.connections {
			// check if trying to open already open valve or valve without flow
			if valves[v].open && (current.inPath&valves[v].id > 0 || valves[v].flow == 0) {
				continue
			}
			val := valves[v]
			val.calc(time, current)
			pressure, ok := cache[val.inPath]
			if ok && pressure > val.f {
				cacheHit++
				continue
			}
			if !multipleWorkers && (maxFlow-val.curFlow)*(time-val.timeToOpen) < bestGoal.f-val.f {
				possibilityCheck++
				continue
			}
			if val.timeToOpen > time || openHash[val.hash] {
				continue
			} else {
				cache[val.inPath] = val.f
				open = append(open, val)
				openHash[val.hash] = true
			}
		}

	}
	return bestGoal.f, bestPaths
}

func findBestDisjointPath(bestPaths map[int64]valve) (valve, valve) {
	v1, v2 := valve{}, valve{}
	for a, aVal := range bestPaths {
	inner:
		for b, bVal := range bestPaths {
			if a == b {
				continue
			}

			if aVal.inPath&bVal.inPath != 0 {
				continue inner
			}
			if aVal.f+bVal.f > v1.f+v2.f {
				v1 = aVal
				v2 = bVal
			}
		}
	}
	return v1, v2
}

var cacheHit = 0
var possibilityCheck = 0
var overallPaths = 0

func Solve(test bool, tasks int) (string, string) {
	start := time.Now()
	inputFile := "day16/input.txt"
	if test {
		inputFile = "day16/testInput.txt"
	}
	valves := parseValves(inputFile)
	if test {
		for _, v := range valves {
			fmt.Println(v)
		}
	}
	parsing := time.Now()

	ret1, ret2 := "", ""
	// Task 1
	if tasks == 0 || tasks == 1 {
		maxPressure, _ := findMaxPressure(valves, 30, "AA", false)
		fmt.Println("Part 1:")
		fmt.Println("Cache hits:", cacheHit)
		fmt.Println("Possibility checks:", possibilityCheck)
		fmt.Println("Max Pressure:", maxPressure)
		ret1 = strconv.Itoa(maxPressure)
	}
	task1 := time.Now()

	// Task 2
	if tasks == 0 || tasks == 2 {
		cacheHit = 0
		possibilityCheck = 0
		_, bestPath := findMaxPressure(valves, 26, "AA", true)
		fmt.Println("\n\nPart 2:")
		fmt.Println("Cache hits:", cacheHit)
		fmt.Println("Possibility checks:", possibilityCheck)
		fmt.Println("Possible paths: ", len(bestPath))
		fmt.Println("Overall paths: ", overallPaths)
		path1, path2 := findBestDisjointPath(bestPath)
		if test {
			fmt.Println("Opened paths:")
			fmt.Println(" PATH 1: ", path1)
			fmt.Println(" PATH 2: ", path2)
		}
		fmt.Println("Max Pressures with help: ", path1.f, path2.f)
		fmt.Println("Sum: ", path1.f+path2.f)
		fmt.Println(path1)
		fmt.Println(path2)
		ret2 = strconv.Itoa(path1.f + path2.f)
	}
	task2 := time.Now()
	fmt.Println("Time: parsing ", parsing.Sub(start), " task1 ", task1.Sub(parsing), " task2 ", task2.Sub(task1), " total ", task2.Sub(start))

	return ret1, ret2
}
