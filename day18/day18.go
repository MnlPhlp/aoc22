package day18

import (
	"fmt"
	"strings"

	"github.com/mnlphlp/aoc22/util"
)

func parseInput(input string) map[util.Pos3]struct{} {
	droplet := make(map[util.Pos3]struct{}, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		pos := util.Pos3{}
		fmt.Sscanf(line, "%d,%d,%d", &pos.X, &pos.Y, &pos.Z)
		droplet[pos] = struct{}{}
	}
	return droplet
}

// all 6 directions
var directions = []util.Pos3{util.Pos3{1, 0, 0}, util.Pos3{-1, 0, 0}, util.Pos3{0, 1, 0}, util.Pos3{0, -1, 0}, util.Pos3{0, 0, 1}, util.Pos3{0, 0, -1}}

func surfaceArea(droplet map[util.Pos3]struct{}) int {
	surface := 0
	for pos := range droplet {
		// check how many faces are not covered by other cubes
		for _, d := range directions {
			if _, ok := droplet[pos.Add(d)]; !ok {
				surface++
			}
		}
	}
	return surface
}

func pocketSurfaceArea(droplet map[util.Pos3]struct{}) int {
	// find the smallest and largest x, y and z coordinates
	min, max := util.Pos3{X: 0, Y: 0, Z: 0}, util.Pos3{X: 0, Y: 0, Z: 0}
	for pos := range droplet {
		min = util.MinPos(min, pos)
		max = util.MaxPos(max, pos)
	}
	// find closed pockets in the droplet
	pockets := make(map[util.Pos3]struct{}, 0)
	for x := min.X + 1; x < max.X; x++ {
		for y := min.Y + 1; y < max.Y; y++ {
			for z := min.Z + 1; z < max.Z; z++ {
				pos := util.Pos3{X: x, Y: y, Z: z}
				if _, solid := droplet[pos]; solid {
					continue
				}
				if _, inPocket := pockets[pos]; inPocket {
					skipDistance := 0
					for inPocket {
						skipDistance++
						_, inPocket = pockets[pos.Add(util.Pos3{X: 0, Y: 0, Z: skipDistance})]
					}
					z += skipDistance - 1
					continue
				}
				// open spot found, might be a pocket
				// check if the opening is connected to the outside
				// If not add to pockets
				pocketTmp := make(map[util.Pos3]struct{})
				open := make(map[util.Pos3]struct{}, 0)
				open[pos] = struct{}{}
				for len(open) > 0 {
					// get first element
					var p util.Pos3
					for p = range open {
						delete(open, p)
						pocketTmp[p] = struct{}{}
						break
					}
					// check if it is connected to the outside
					if p.X <= min.X || p.X >= max.X || p.Y <= min.Y || p.Y >= max.Y || p.Z <= min.Z || p.Z >= max.Z {
						pocketTmp = nil
						break
					}
					// add all neighbors to open
					for _, d := range directions {
						n := p.Add(d)
						if _, solid := droplet[n]; !solid {
							if _, inPocket := pocketTmp[n]; !inPocket {
								open[n] = struct{}{}
							}
						}
					}
				}
				for p := range pocketTmp {
					pockets[p] = struct{}{}
				}

			}
		}
	}
	// calculate surface area of pockets
	return surfaceArea(pockets)
}

func Solve(input string, debug bool, task int) (string, string) {
	res1, res2 := 0, 0
	droplet := parseInput(input)
	if debug {
		fmt.Println("Droplet: ", droplet)
	}
	if task != 2 {
		res1 = surfaceArea(droplet)
		fmt.Println("Result 1: ", res1)
	}
	if task != 1 {
		if res1 == 0 {
			res1 = surfaceArea(droplet)
		}
		pocketArea := pocketSurfaceArea(droplet)
		res2 = res1 - pocketArea
		fmt.Println("Result 2: ", res2)
	}

	return fmt.Sprint(res1), fmt.Sprint(res2)
}
