package day19

import (
	"testing"

	"github.com/mnlphlp/aoc22/util"
	"github.com/stretchr/testify/assert"
)

var blueprints = parseInput(util.ReadInputUnittest(19, false))

func TestTask1(t *testing.T) {
	qualities := getQualityNumbers(blueprints, 24, true)
	assert.Equal(t, 1192, util.Sum(qualities))
}

func TestTask2(t *testing.T) {
	qualities := getMaxGeodes(blueprints[:3], 32, true)
	assert.Equal(t, 14725, util.Mul(qualities))
}
