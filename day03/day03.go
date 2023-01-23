package day03

import (
	"fmt"
	"strings"
)

func findDuplicate(r string) byte {
	seen := uint64(0)
	comp1 := r[:len(r)/2]
	comp2 := r[len(r)/2:]
	for i := 0; i < len(r)/2; i++ {
		// set bit for character in rucksack 1
		seen |= 1 << (comp1[i] - 65)
		// check if character bit for rucksack 2 element is set
		if seen&uint64(comp2[i]-65) > 0 {
			return byte(comp1[i])
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
	seen := make([]uint64, 3)
	// loop over groups
	for i := 0; i < len(rucksacks); i += 3 {
		// create set of seen items for each rucksack in one group
		for j := 0; j < 3; j++ {
			for _, c := range rucksacks[i+j] {
				seen[j] |= 1 << (c - 65)
			}
		}
		// search bit (character) set in all 3 rucksacks
		for r := 0; r < 64; r++ {
			if ((1 << r) & seen[0] & seen[1] & seen[2]) > 0 {
				badges[i/3] = byte(r)
				break
			}
		}
	}
	return badges
}

func Solve(input string, test bool, task int) (string, string) {
	rucksacks := strings.Split(string(input), "\n")
	if test {
		fmt.Println("Rucksacks:", rucksacks)
	}
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
