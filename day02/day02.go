package day02

import (
	"fmt"
	"os"
	"strings"
)

const (
	ROCK = iota
	PAPER
	SCISSORS
)

func getMove(move byte) byte {
	if move < 'D' {
		return move - 'A'
	} else {
		return move - 'X'
	}
}

func calcScore(line string) int {
	opponent := getMove(line[0])
	self := getMove(line[2])
	// score for pick
	score := int(self + 1)
	// score for draw
	if self == opponent {
		score += 3
	}
	// score for winning
	if (self+2)%3 == opponent {
		score += 6
	}
	return score
}

func decideMove(opponent, self byte) byte {
	// start with draw move
	move := opponent
	switch self {
	case 'X':
		// pick losing move
		move = (opponent + 2) % 3
	case 'Z':
		// pick winning move
		move = (opponent + 1) % 3
	}
	return move
}

func calcScore2(line string) int {
	opponent := getMove(line[0])
	self := decideMove(opponent, line[2])
	// score for pick
	score := int(self + 1)
	// score for draw
	if self == opponent {
		score += 3
	}
	// score for winning
	if (self+2)%3 == opponent {
		score += 6
	}
	return score
}

func Solve(test bool) (string, string) {
	input, _ := os.ReadFile("day02/input.txt")
	lines := strings.Split(string(input), "\n")
	score := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		score += calcScore(line)
	}
	fmt.Printf("Score 1: %v\n", score)
	res1 := fmt.Sprintf("%d", score)

	score = 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		score += calcScore2(line)
	}
	fmt.Printf("Score 2: %v\n", score)
	res2 := fmt.Sprintf("%d", score)
	return res1, res2
}
