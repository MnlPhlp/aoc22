package day16

import (
	"os"
	"testing"

	"github.com/mnlphlp/aoc22/util"
	"github.com/stretchr/testify/assert"
)

var res1, res2 string

func solve() {
	if res1 == "" {
		os.Chdir("..")
		input := util.ReadInput(16, false)
		res1, res2 = Solve(input, false, 0)
	}
}

func TestDay16Task1(t *testing.T) {
	solve()
	assert.Equal(t, "1741", res1)
}

func TestDay16Task2(t *testing.T) {
	solve()
	assert.Equal(t, "2316", res2)
}
