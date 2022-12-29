package day19

import (
	"fmt"
	"strings"
)

type ressources struct {
	clay     int
	ore      int
	obsidian int
}

type blueprint struct {
	id            int
	oreRobot      ressources
	clayRobot     ressources
	obsidianRobot ressources
	geodeRobot    ressources
}

func parseInput(input string) []blueprint {
	input = strings.Replace(input, "\n", "", -1)
	blueprints := []blueprint{}
	for _, bpStr := range strings.Split(input, "Blueprint") {
		if bpStr == "" {
			continue
		}
		bp := blueprint{}
		fmt.Sscanf(bpStr,
			" %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.id, &bp.oreRobot.ore, &bp.clayRobot.ore, &bp.obsidianRobot.ore, &bp.obsidianRobot.clay, &bp.geodeRobot.ore, &bp.geodeRobot.obsidian,
		)
		blueprints = append(blueprints, bp)
	}
	return blueprints
}

func Solve(input string, debug bool, task int) (string, string) {
	res1, res2 := 0, 0
	blueprints := parseInput(input)
	if debug {
		fmt.Println(blueprints)
	}
	return fmt.Sprint(res1), fmt.Sprint(res2)
}
