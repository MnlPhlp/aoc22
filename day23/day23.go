package day23

import (
	"fmt"
	"strconv"

	"github.com/mnlphlp/aoc22/day23/grid"
	"github.com/mnlphlp/aoc22/util"
)

const (
	// weird order copied from task
	North = iota
	South
	West
	East
)

func move(elves grid.Grid, rounds int, startDir int, debug bool) (grid.Grid, int) {
	if debug {
		fmt.Println(elves)
	}
	lastHash := elves.Hash()
	for i := 0; i < rounds || rounds == 0; i++ {
		proposed := make(map[util.Pos2]util.Pos2)
		conflicts := make([]util.Pos2, 0)
		elves.ForEach(func(p util.Pos2) {
			if !elves.HasNeighbor(p) {
				return
			}
			next, nextFound := elves.NextPos(p, startDir)
			if nextFound {
				if _, conflict := proposed[next]; conflict {
					// do not move and mark other move for deletion
					conflicts = append(conflicts, next)
				} else {
					proposed[next] = p
				}
			}
		})
		for _, conflict := range conflicts {
			delete(proposed, conflict)
		}
		for _, old := range proposed {
			elves = elves.Remove(old)
		}
		for new := range proposed {
			elves = elves.Insert(new)
		}
		if debug {
			fmt.Printf("Round  %d: %d\n", i+1, countEmpty(elves))
			fmt.Println(elves)
		}
		startDir = (startDir + 1) % 4
		hash := elves.Hash()
		if hash.Equals(lastHash) && rounds == 0 {
			return elves, i + 1
		}
		lastHash = hash
	}
	return elves, rounds
}

func countEmpty(g grid.Grid) int {
	min, max := util.Pos2{X: 1 << 62, Y: 1 << 62}, util.Pos2{X: -(1 << 62), Y: -(1 << 62)}
	count := 0
	g.ForEach(func(p util.Pos2) {
		min = util.Pos2{X: util.Min(min.X, p.X), Y: util.Min(min.Y, p.Y)}
		max = util.Pos2{X: util.Max(max.X, p.X), Y: util.Max(max.Y, p.Y)}
		count++
	})
	return ((max.X - min.X + 1) * (max.Y - min.Y + 1)) - count
}

func part1(grid grid.Grid, debug bool) int {
	grid, _ = move(grid, 10, North, debug)
	return countEmpty(grid)
}

func part2(grid grid.Grid, startDir int, debug bool) int {
	_, lastMove := move(grid, 0, startDir, debug)
	return lastMove
}

func Solve(input string, debug bool, task int) (string, string) {
	res1, res2 := 0, 0
	g := grid.ParseInput(input)
	if task != 2 {
		res1 = part1(g, debug)
	}
	if task != 1 {
		// if part one already ran continue from that state
		startDir := North
		if task == 0 {
			startDir = (startDir + 10) % 4
			res2 = 10
		}
		res2 += part2(g, startDir, debug)
	}
	if grid.TIMING_ACTIVE {
		fmt.Println("canMove: ", grid.Timing.CanMove)
		fmt.Println("contains: ", grid.Timing.Contains)
		fmt.Println("hash: ", grid.Timing.Hash)
		fmt.Println("insert: ", grid.Timing.Insert)
		fmt.Println("nextMove: ", grid.Timing.NextMove)
		fmt.Println("remove: ", grid.Timing.Remove)
	}

	return strconv.Itoa(res1), strconv.Itoa(res2)
}
