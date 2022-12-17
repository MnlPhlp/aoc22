package day11

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	number     int
	items      []int
	throwTrue  int
	throwFalse int
	divTest    int
	operation  func(int) int
	inspected  int
}

func (m monkey) String() string {
	return fmt.Sprintf(`Monkey %d: 
	items: %v
	throwTrue: %d
	throwFalse: %d
	divTest: %d
`, m.number, m.items, m.throwTrue, m.throwFalse, m.divTest)
}

func (m *monkey) playRound(monkeys []monkey, modVal int) {
	for i := range m.items {
		throwTo := 0
		m.items[i] = m.operation(m.items[i])
		if modVal == 0 {
			m.items[i] /= 3
		} else {
			m.items[i] = m.items[i] % modVal
		}
		if m.items[i]%m.divTest == 0 {
			throwTo = m.throwTrue
		} else {
			throwTo = m.throwFalse
		}
		monkeys[throwTo].items = append(monkeys[throwTo].items, m.items[i])
		m.inspected++
	}
	m.items = []int{}
}

func parseMonkey(input string) monkey {
	// parse input
	m := monkey{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		lineParts := strings.Split(line, " ")
		if lineParts[0] == "Starting" {
			for _, itemStr := range strings.Split(strings.Split(line, ": ")[1], ", ") {
				item, _ := strconv.Atoi(strings.TrimSpace(itemStr))
				m.items = append(m.items, item)
			}
		} else if lineParts[0] == "Operation:" {
			operator := lineParts[4]
			operand, _ := strconv.Atoi(lineParts[5])
			switch operator {
			case "+":
				if lineParts[5] == "old" {
					m.operation = func(old int) int { return old + old }
				} else {
					m.operation = func(old int) int { return old + operand }
				}
			case "*":
				if lineParts[5] == "old" {
					m.operation = func(old int) int { return old * old }
				} else {
					m.operation = func(old int) int { return old * operand }
				}
			}
		} else if lineParts[0] == "Test:" {
			m.divTest, _ = strconv.Atoi(lineParts[3])
		} else if lineParts[0] == "If" {
			if lineParts[1] == "true:" {
				m.throwTrue, _ = strconv.Atoi(lineParts[5])
			} else {
				m.throwFalse, _ = strconv.Atoi(lineParts[5])
			}
		} else if line != "" {
			m.number, _ = strconv.Atoi(line[:len(line)-1])
		}
	}
	return m
}

func parseInput(input string) []monkey {
	monkeys := []monkey{}
	// split input into monkeys
	for _, monkeyInput := range strings.Split(input, "Monkey ") {
		if monkeyInput == "" {
			continue
		}
		// for each monkey, parse input
		monkeys = append(monkeys, parseMonkey(monkeyInput))
	}
	return monkeys
}

func Solve(test bool) {
	var input []byte
	if test {
		input, _ = os.ReadFile("day11/testInput.txt")
	} else {
		input, _ = os.ReadFile("day11/input.txt")
	}
	monkeys := parseInput(string(input))
	if test {
		for _, m := range monkeys {
			fmt.Println(m)
		}
	}

	// Task 1
	// run for 20 rounds
	for i := 0; i < 20; i++ {
		for m := range monkeys {
			monkeys[m].playRound(monkeys, 0)
		}
		if i == 0 && test {
			for _, m := range monkeys {
				fmt.Printf("Monkey %d: %v\n", m.number, m.items)
			}
		}
	}

	// find two most active monkeys
	if test {
		for _, m := range monkeys {
			fmt.Printf("Monkey %d inspected items %d times\n", m.number, m.inspected)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	fmt.Println("Result Task 1: ", monkeys[0].inspected*monkeys[1].inspected)

	// Task 2
	// reset monkeys
	monkeys = parseInput(string(input))
	modVal := 1
	for _, m := range monkeys {
		if modVal%m.divTest != 0 {
			modVal *= m.divTest
		}
	}
	// run for 10000 rounds
	for i := 0; i < 10000; i++ {
		for m := range monkeys {
			monkeys[m].playRound(monkeys, modVal)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	fmt.Println("Result Task 2: ", monkeys[0].inspected*monkeys[1].inspected)

}
