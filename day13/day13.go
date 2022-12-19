package day13

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type pair struct {
	left  item
	right item
}

func (p pair) String() string {
	return fmt.Sprintf("%v : %v", p.left, p.right)
}

type item struct {
	leaf   bool
	value  int
	values []item
}

func (i item) String() string {
	if i.leaf {
		return fmt.Sprintf("%d", i.value)
	}
	text := "["
	for _, item := range i.values {
		text += item.String()
		text += ","
	}
	if len(text) > 1 {
		text = text[:len(text)-1]
	}
	text += "]"
	return text
}

func getParts(line string) []string {
	parts := make([]string, 0)
	depth := 0
	partTmp := ""
	for _, c := range line {
		partTmp += string(c)
		if c == '[' {
			depth++
		} else if c == ']' {
			depth--
		} else if c == ',' && depth == 0 {
			parts = append(parts, partTmp[:len(partTmp)-1])
			partTmp = ""
		}
	}
	if len(partTmp) > 0 {
		parts = append(parts, partTmp)
	}
	return parts
}

func minLength[T any](a, b []T) int {
	if len(a) < len(b) {
		return len(a)
	}
	return len(b)
}

func parseItem(line string) item {
	item := item{}
	i, err := strconv.Atoi(line)
	if err == nil {
		item.leaf = true
		item.value = i
	} else if len(line) == 0 {
		return item
	} else if line[0] == '[' {
		// parse array
		if line == "[]" {
			return item
		}
		parts := getParts(line[1 : len(line)-1])
		for _, part := range parts {
			item.values = append(item.values, parseItem(part))
		}
	}
	return item
}

func parseInput(inputFile string) []pair {
	pairs := make([]pair, 0)
	input, _ := os.ReadFile(inputFile)
	for _, part := range strings.Split(string(input), "\n\n") {
		if len(part) == 0 {
			continue
		}
		pair := pair{}
		for i, line := range strings.Split(part, "\n") {
			if len(line) == 0 {
				continue
			}
			item := parseItem(line)
			if i == 0 {
				pair.left = item
			} else {
				pair.right = item
			}
		}
		pairs = append(pairs, pair)
	}
	return pairs
}

func parse2(inputFile string) []item {
	items := make([]item, 0)
	input, _ := os.ReadFile(inputFile)
	for _, line := range strings.Split(string(input), "\n") {
		if len(line) == 0 {
			continue
		}
		items = append(items, parseItem(line))
	}
	return items
}

func checkPair(p pair) (bool, bool) {
	// both are leafs
	if p.left.leaf && p.right.leaf {
		return p.left.value <= p.right.value, p.left.value != p.right.value
	}
	// mixed types
	if p.left.leaf {
		p.left.leaf = false
		p.left.values = append(p.left.values, item{leaf: true, value: p.left.value})
	}
	if p.right.leaf {
		p.right.leaf = false
		p.right.values = append(p.right.values, item{leaf: true, value: p.right.value})
	}
	// both are arrays
	min := minLength(p.left.values, p.right.values)
	for i := 0; i < min; i++ {
		if ok, finished := checkPair(pair{left: p.left.values[i], right: p.right.values[i]}); ok && finished {
			return true, true
		} else if !ok {
			return false, true
		}
	}
	return len(p.left.values) <= len(p.right.values), len(p.left.values) != len(p.right.values)
}

func checkOrder(pairs []pair) []bool {
	orders := make([]bool, len(pairs))
	for i, pair := range pairs {
		orders[i], _ = checkPair(pair)
	}
	return orders
}

func Solve(test bool) (string, string, time.Duration) {
	start := time.Now()
	inputFile := "day13/input.txt"
	if test {
		inputFile = "day13/testInput.txt"
	}
	pairs := parseInput(inputFile)

	if test {
		fmt.Println(pairs)
	}
	orders := checkOrder(pairs)
	result := 0
	for i, order := range orders {
		if order {
			fmt.Print(i+1, ";")
			result += i + 1
		}
	}
	fmt.Println()
	fmt.Println("result task 1: ", result)
	res1 := strconv.Itoa(result)
	checkPair(pairs[7])

	// task 2
	items := parse2(inputFile)
	items = append(items, parseItem("[[2]]"))
	items = append(items, parseItem("[[6]]"))

	sort.Slice(items, func(i, j int) bool {
		ok, _ := checkPair(pair{left: items[i], right: items[j]})
		return ok
	})

	key1, key2 := 0, 0
	for i, item := range items {
		if item.String() == "[[2]]" {
			key1 = i + 1
		}
		if item.String() == "[[6]]" {
			key2 = i + 1
		}
	}
	fmt.Println("Key 1: ", key1)
	fmt.Println("Key 2: ", key2)
	fmt.Println("result task 2: ", key1*key2)
	res2 := strconv.Itoa(key1 * key2)
	return res1, res2, time.Since(start)
}
