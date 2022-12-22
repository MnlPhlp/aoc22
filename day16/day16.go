package day16

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type valve struct {
	name        string
	flow        int
	g           int
	f           int
	open        bool
	connections []*valve
	parent      *valve
}

func (v *valve) calc(time int) {
	if v.parent != nil {
		v.g = v.parent.g + 1
	}
	v.f = v.flow * (time - v.g)
}

func (v valve) String() string {
	connections := ""
	if len(v.connections) > 0 {
		for _, c := range v.connections {
			connections += c.name + ", "
		}
		if len(connections) > 2 {
			connections = connections[:len(connections)-2]
		}
	}
	return fmt.Sprintf("Valve: %s, Flow: %d, Connections to: %s", v.name, v.flow, connections)
}

func parseValves(file string) (map[string]*valve, string) {
	valves := make(map[string]*valve)
	connections := make(map[string]string)
	input, _ := os.ReadFile(file)
	start := string(input[6:8])
	for _, l := range strings.Split(string(input), "\n") {
		if l == "" {
			continue
		}
		parts := strings.Split(l, " ")
		connection := ""
		for _, p := range parts[9:] {
			connection += p
		}
		flow, _ := strconv.Atoi(parts[4][5 : len(parts[4])-1])
		valves[parts[1]+"C"] = &valve{
			name: parts[1] + "C",
			flow: flow,
		}
		valves[parts[1]+"O"] = &valve{
			name: parts[1] + "O",
			flow: flow,
			open: true,
		}
		connections[parts[1]+"O"] = connection
		connections[parts[1]+"C"] = connection
	}
	for name, connection := range connections {
		for _, c := range strings.Split(connection, ",") {
			valves[name].connections = append(valves[name].connections, valves[c+"O"], valves[c+"C"])
		}
		if !valves[name].open {
			valves[name].connections = append(valves[name].connections, valves[name[:len(name)-1]+"O"])
		}
	}
	return valves, start
}

func findMaxPressure(valves map[string]*valve, time int, startPoint string) int {
	pressure := 0
	closed := make(map[string]bool)
	open := make([]*valve, 0)
	// start with closed version of first valve
	current := valves[startPoint+"C"]
	current.calc(time)
	open = append(open, current)

	// explore paths using a* like algorithm
	for len(open) > 0 {
		maxScore := 0
		for i, v := range open {
			if v.f > open[maxScore].f {
				maxScore = i
			}
		}
		current = open[maxScore]
		open = append(open[:maxScore], open[maxScore+1:]...)
		closed[current.name] = true

	}

	// find valve with best f
	for _, v := range valves {
		if v.f > pressure {
			pressure = v.f
		}
	}

	return pressure
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
	ret1 := strconv.Itoa(maxPressure)

	return ret1, "implemented"
}
