package day19

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mnlphlp/aoc22/util"
)

type ressources struct {
	clay     int
	ore      int
	obsidian int
}

type blueprint struct {
	id              int
	oreRobot        ressources
	clayRobot       ressources
	obsidianRobot   ressources
	geodeRobot      ressources
	maxOreCost      int
	maxClayCost     int
	maxObsidianCost int
}

type state struct {
	clay, ore, obsidian                        int
	clayBots, oreBots, obsidianBots, geodeBots int
	geodes                                     int
	timePassed                                 int
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
		bp.maxClayCost = util.Max(bp.clayRobot.clay, bp.oreRobot.clay, bp.obsidianRobot.clay, bp.geodeRobot.clay)
		bp.maxOreCost = util.Max(bp.clayRobot.ore, bp.oreRobot.ore, bp.obsidianRobot.ore, bp.geodeRobot.ore)
		bp.maxObsidianCost = util.Max(bp.clayRobot.obsidian, bp.oreRobot.obsidian, bp.obsidianRobot.obsidian, bp.geodeRobot.obsidian)
		blueprints = append(blueprints, bp)
	}
	return blueprints
}

func getQualityNumbers(blueprints []blueprint, timeSteps int, parallel bool) []int {
	qualityNumbers := make([]int, len(blueprints))
	wg := sync.WaitGroup{}
	for i, bp := range blueprints {
		if parallel {
			wg.Add(1)
			go func(i int, bp blueprint) {
				defer wg.Done()
				qualityNumbers[i] = getMaxGeode(bp, timeSteps) * bp.id
			}(i, bp)
		} else {
			qualityNumbers[i] = getMaxGeode(bp, timeSteps) * bp.id
		}
	}
	wg.Wait()
	return qualityNumbers
}

func hashState(s state) int {
	hash := s.clay
	hash |= s.ore << 8
	hash |= s.obsidian << 16
	hash |= s.timePassed << 24
	return hash
}

func canBuild(s state, r ressources) bool {
	return s.clay >= r.clay && s.obsidian >= r.obsidian && s.ore >= r.ore
}

func consumeRessources(s state, r ressources) state {
	s.clay -= r.clay
	s.ore -= r.ore
	s.obsidian -= r.obsidian
	return s
}

func getNeighbors(s state, bp blueprint, time int) []state {
	neighbors := make([]state, 0)
	if s.timePassed == time-1 {
		return neighbors
	}
	// create doNothing state
	doNothing := s
	doNothing.timePassed++
	doNothing.clay += doNothing.clayBots
	doNothing.ore += doNothing.oreBots
	doNothing.obsidian += doNothing.obsidianBots
	// geodes are computed directly
	// if geode bot is possible only build this
	if canBuild(s, bp.geodeRobot) {
		new := consumeRessources(doNothing, bp.geodeRobot)
		new.geodeBots++
		new.geodes += time - new.timePassed
		// if max count for all other robots is reached skip next steps
		if s.clayBots == bp.maxClayCost && s.oreBots == bp.maxOreCost && s.obsidianBots == bp.maxObsidianCost {
			// one geode bot can be build per time step
			remTime := time - new.timePassed
			new.geodeBots += remTime
			new.geodes += int(float64(remTime) * float64(remTime) / 2)
			new.timePassed = time
		}
		neighbors = append(neighbors, new)
		return neighbors
	}
	// append doNothing
	neighbors = append(neighbors, doNothing)
	// append other possible bots
	if s.clayBots < bp.maxClayCost && canBuild(s, bp.clayRobot) {
		new := consumeRessources(doNothing, bp.clayRobot)
		new.clayBots++
		neighbors = append(neighbors, new)
	}
	if s.oreBots < bp.maxOreCost && canBuild(s, bp.oreRobot) {
		new := consumeRessources(doNothing, bp.oreRobot)
		new.oreBots++
		neighbors = append(neighbors, new)
	}
	if s.obsidianBots < bp.maxObsidianCost && canBuild(s, bp.obsidianRobot) {
		new := consumeRessources(doNothing, bp.obsidianRobot)
		new.obsidianBots++
		neighbors = append(neighbors, new)
	}
	return neighbors
}

func getMaxGeode(bp blueprint, timeSteps int) int {
	// calculate max number of geodes cracked
	current := state{oreBots: 1}
	open := []state{current}
	best := current
	closed := map[int]struct{}{}
	for len(open) > 0 {
		// get next state
		current = open[len(open)-1]
		open = open[:len(open)-1]
		// check if new max can be reached
		remTime := timeSteps - current.timePassed
		if current.geodes+(remTime*(remTime/2)) <= best.geodes {
			continue
		}
		closed[hashState(current)] = struct{}{}
		if current.geodes > best.geodes {
			best = current
		}
		for _, n := range getNeighbors(current, bp, timeSteps) {
			if _, ok := closed[hashState(n)]; ok {
				continue
			}
			open = append(open, n)
		}
	}
	if debug {
		fmt.Printf("Result blueprint %d: can open %d geodes\n", bp.id, best.geodes)
	}
	return best.geodes
}

func getMaxGeodes(blueprints []blueprint, timeSteps int, parallel bool) []int {
	maxGeodes := make([]int, len(blueprints))
	wg := sync.WaitGroup{}
	for i, bp := range blueprints {
		if parallel {
			wg.Add(1)
			go func(i int, bp blueprint) {
				defer wg.Done()
				maxGeodes[i] = getMaxGeode(bp, timeSteps)
			}(i, bp)
		} else {
			maxGeodes[i] = getMaxGeode(bp, timeSteps)
		}
	}
	wg.Wait()
	return maxGeodes
}

var debug = false

func Solve(input string, debugFlag bool, task int) (string, string) {
	res1, res2 := 0, 0
	debug = debugFlag
	blueprints := parseInput(input)
	if debug {
		fmt.Println(blueprints)

	}
	if task != 2 {
		qualities := getQualityNumbers(blueprints, 24, !debug)
		fmt.Printf("Quality numbers: %v\n", qualities)
		res1 = util.Sum(qualities)
	}
	if task != 1 {
		maxGeodes := getMaxGeodes(blueprints[:util.Min(3, len(blueprints))], 32, !debug)
		fmt.Printf("Max geodes: %v\n", maxGeodes)
		res2 = util.Mul(maxGeodes)
	}
	return fmt.Sprint(res1), fmt.Sprint(res2)
}
