package day21

import (
	"strconv"
	"strings"

	"github.com/mnlphlp/aoc22/util"
)

type monkey struct {
	name           string
	op             byte
	m1             *monkey
	m2             *monkey
	finished       bool
	result         int
	effectedByHumn bool
}

func (m *monkey) getResult() (int, bool) {
	if !m.m1.finished || !m.m2.finished {
		return 0, false
	}
	m.effectedByHumn = m.m1.effectedByHumn || m.m2.effectedByHumn
	switch m.op {
	case '+':
		return m.m1.result + m.m2.result, true
	case '-':
		return m.m1.result - m.m2.result, true
	case '*':
		return m.m1.result * m.m2.result, true
	case '/':
		return m.m1.result / m.m2.result, true
	}
	return 1, true
}

func parseInput(input string) map[string]*monkey {
	monkeys := make(map[string]*monkey)
	funcs := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		words := strings.Split(line, " ")
		name := words[0][:len(words[0])-1]
		m := &monkey{
			name:           name,
			effectedByHumn: name == "humn",
		}
		if res, err := strconv.Atoi(words[1]); err == nil {
			// This monkey is finished
			m.result = res
			m.finished = true
		} else {
			// store func for later parsing
			funcs[name] = words[1:]
		}
		monkeys[name] = m
	}
	for _, m := range monkeys {
		funcStr := funcs[m.name]
		if !m.finished {
			m.m1 = monkeys[funcStr[0]]
			m.op = funcStr[1][0]
			m.m2 = monkeys[funcStr[2]]

		}
	}
	return monkeys
}

func solveMonkey(monkeys map[string]*monkey, name string) int {
	monkey := monkeys[name]
	for !monkey.finished {
		for _, m := range monkeys {
			if !m.finished {
				res, ok := m.getResult()
				if ok {
					m.result = res
					m.finished = true
				}
			}
		}
	}
	return monkey.result
}

func resetEffectedMonkeys(monkeys map[string]*monkey) {
	for _, m := range monkeys {
		if m.effectedByHumn && m.name != "humn" {
			m.finished = false
		}
	}
}

func solveHumn(monkeys map[string]*monkey, debug bool) int {
	if debug {
		effected := 0
		for _, m := range monkeys {
			if m.effectedByHumn {
				effected++
			}
		}
	}
	humn := monkeys["humn"]
	root := monkeys["root"]
	goal := 0
	var variable *monkey
	if !root.m1.effectedByHumn {
		goal = root.m1.result
		variable = root.m2
	} else if !root.m2.effectedByHumn {
		goal = root.m2.result
		variable = root.m1
	} else {
		panic("can not solve this case")
	}
	lastDelta := util.Abs(variable.result - goal)
	lastSign := 1
	op := 32 * 32
	close := false
	for lastDelta > 0 {
		humn.result += op
		resetEffectedMonkeys(monkeys)
		res := solveMonkey(monkeys, variable.name)
		deltaSigned := res - goal
		delta := util.Abs(deltaSigned)
		if delta == 0 {
			break
		}
		sign := util.Sign(deltaSigned)
		if (delta > lastDelta || sign != lastSign) && close {
			if util.Abs(op) > 1 {
				op /= -2
			} else {
				op *= -1
			}
		} else {
			if !close {
				if sign != lastSign {
					op *= -1
				}
				deltaDiff := lastDelta - delta
				if deltaDiff > 0 && delta > deltaDiff {
					op *= delta / deltaDiff
				} else {
					op = delta * util.Sign(op)
					close = true
				}
			}
		}

		lastDelta = delta
		lastSign = sign
	}

	return humn.result
}

func Solve(input string, debug bool, task int) (string, string) {
	res1, res2 := 0, 0
	monkeys := parseInput(input)
	if debug {
		for _, m := range monkeys {
			println(m.name)
		}
	}
	// always compute res1 because it is used for res2
	res1 = solveMonkey(monkeys, "root")
	if task != 1 {
		res2 = solveHumn(monkeys, debug)
	}
	return strconv.Itoa(res1), strconv.Itoa(res2)
}
