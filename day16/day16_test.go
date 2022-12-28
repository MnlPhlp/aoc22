package day16

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var res1, res2 string

func solve() {
	if res1 == "" {
		os.Chdir("..")
		res1, res2 = Solve(false, 0)
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
