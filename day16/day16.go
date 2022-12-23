package day16

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type valve struct {
	flow        int
	f           int
	timeToOpen  int
	open        bool
	connections []string
	id          int
	inPath      int
	curFlow     int
}

func (v *valve) calc(time int, current valve) {
	v.timeToOpen = current.timeToOpen + 1
	v.f = current.f + v.flow*(time-v.timeToOpen)
	v.curFlow = current.curFlow + v.flow
}

func (v valve) String() string {
	return fmt.Sprintf("Valve: %d, Flow: %d, Connections to: %v", v.id, v.flow, v.connections)
}

func (v *valve) IsInSlice(slice []valve) bool {
	for _, s := range slice {
		if s.id == v.id && s.open == v.open && s.f == v.f {
			return true
		}
	}
	return false
}

func parseValves(file string) (map[string]valve, string) {
	valves := make(map[string]valve)
	connections := make(map[string]string)
	input, _ := os.ReadFile(file)
	start := string(input[6:8])
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
		id := 1 << i
		valves[parts[1]+"C"] = valve{
			flow: 0,
			id:   id,
		}
		valves[parts[1]+"O"] = valve{
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
		if !valves[name].open {
			val.connections = append(val.connections, name[:len(name)-1]+"O")
		}
		valves[name] = val
	}
	return valves, start
}

func findMaxPressure(valves map[string]valve, time int, startPoint string) valve {
	open := make([]valve, 0)
	// start with closed version of first valve
	current := valves[startPoint+"C"]
	current.calc(time, valve{})
	open = append(open, current)
	bestGoal := current
	maxFlow := 0
	for _, v := range valves {
		maxFlow += v.flow
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

		// skip item if max score is not reachable even with all valves open
		if current.f+(maxFlow-current.curFlow)*(time-current.timeToOpen) < bestGoal.f {
			continue
		} else if current.f > bestGoal.f {
			bestGoal = current
		}

		for _, v := range current.connections {
			val := valves[v]
			if !val.open || current.inPath&val.id == 0 {
				val.inPath = current.inPath
				if current.open {
					val.inPath |= current.id
				}
				val.calc(time, current)
				if val.timeToOpen < time && !val.IsInSlice(open) {
					open = append(open, val)
				}
			}
		}

	}
	return bestGoal
}

func Solve(test bool) (string, string) {
	inputFile := "day16/input.txt"
	if test {
		inputFile = "day16/testInput.txt"
	}
	valves, startValve := parseValves(inputFile)
	if test {
		fmt.Println("Start valve:", startValve)
		for _, v := range valves {
			fmt.Println(v)
		}
	}

	// Task 1
	maxPressure := findMaxPressure(valves, 30, startValve)
	fmt.Println("Max Pressure", maxPressure.f)
	ret1 := strconv.Itoa(maxPressure.f)

	return ret1, "implemented"
}
