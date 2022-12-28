package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func sumCalories(elfes []string) []int {
	calories := make([]int, len(elfes))
	for i, elfe := range elfes {
		for _, line := range strings.Split(elfe, "\n") {
			if line == "" {
				continue
			}
			cal, _ := strconv.Atoi(line)
			calories[i] += cal
		}
	}
	return calories
}

func Solve(input string, test bool, task int) (string, string) {
	res1 := ""
	res2 := ""

	elfes := strings.Split(input, "\n\n")
	calories := sumCalories(elfes)
	sort.Ints(calories)

	if task != 2 {
		maxCalories := calories[len(calories)-1]
		res1 = fmt.Sprintf("%d", maxCalories)
		fmt.Printf("Max calories: %d\n", maxCalories)
	}

	if task != 1 {
		topCalories := 0
		for _, cal := range calories[len(calories)-3:] {
			topCalories += cal
		}
		fmt.Printf("Top calories: %d\n", topCalories)
		res2 = fmt.Sprintf("%d", topCalories)
	}

	return res1, res2
}
