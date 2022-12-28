package day03

import (
	"fmt"
	"strings"
)

func findDuplicate(r string) byte {
	seen := make(map[rune]bool)
	comp1 := r[:len(r)/2]
	comp2 := r[len(r)/2:]
	for _, c := range comp1 {
		seen[c] = true
	}
	for _, c := range comp2 {
		if seen[c] {
			return byte(c)
		}
	}
	return 0
}

func getPriority(duplicate byte) int {
	if duplicate >= 'a' && duplicate <= 'z' {
		return int(duplicate-'a') + 1
	} else {
		return int(duplicate-'A') + 27
	}
}

func findBadges(rucksacks []string) []byte {
	badges := make([]byte, len(rucksacks)/3)
	seen := make([]map[rune]bool, 3)
	for i := 0; i < len(rucksacks); i += 3 {
		// create set of seen items
		for j := 0; j < 3; j++ {
			seen[j] = make(map[rune]bool)
			for _, c := range rucksacks[i+j] {
				seen[j][c] = true
			}
		}
		// store badge
		for r, _ := range seen[2] {
			if seen[0][r] && seen[1][r] {
				badges[i/3] = byte(r)
				break
			}
		}
	}
	return badges
}

func Solve(input string, test bool, task int) (string, string) {
	rucksacks := strings.Split(string(input), "\n")
	// remove empty line
	rucksacks = rucksacks[:len(rucksacks)-1]
	res1, res2 := "", ""
	if task != 2 {
		sum := 0
		for _, r := range rucksacks {
			duplicate := findDuplicate(r)
			sum += getPriority(duplicate)
		}
		fmt.Println("Result 1:", sum)
		res1 = fmt.Sprint(sum)
	}
	if task != 1 {
		badgeSum := 0
		badges := findBadges(rucksacks)
		for _, b := range badges {
			badgeSum += getPriority(b)
		}
		for _, b := range badges {
			fmt.Print(string(b))
		}
		fmt.Println()
		fmt.Println("Result 2:", badgeSum)
		res2 = fmt.Sprint(badgeSum)
	}
	return res1, res2
}
