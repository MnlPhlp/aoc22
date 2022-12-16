package main

import (
	"fmt"
	"os"
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

func main() {
	input, _ := os.ReadFile("input.txt")
	elfes := strings.Split(string(input), "\n\n")
	calories := sumCalories(elfes)
	sort.Ints(calories)
	maxCalories := calories[len(calories)-1]
	fmt.Printf("Max calories: %d\n", maxCalories)
}
