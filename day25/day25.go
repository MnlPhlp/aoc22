package day25

import (
	"fmt"
	"strings"
)

var snafuChar = map[int]string{
	-2: "=",
	-1: "-",
	0:  "0",
	1:  "1",
	2:  "2",
}
var snafuNum = map[string]int{
	"=": -2,
	"-": -1,
	"0": 0,
	"1": 1,
	"2": 2,
}

func snafuAdd(a, b string) string {
	// reversed...
	split1, split2 := strings.Split(a, ""), strings.Split(b, "")
	var reversed1, reversed2 []string
	for i := len(split1) - 1; i >= 0; i-- {
		reversed1 = append(reversed1, split1[i])
	}
	for i := len(split2) - 1; i >= 0; i-- {
		reversed2 = append(reversed2, split2[i])
	}

	longer, shorter := reversed1, reversed2
	if len(longer) < len(shorter) {
		longer, shorter = shorter, longer
	}

	ans := make([]int, len(longer)+1)
	for i := 0; i < len(longer); i++ {
		sum := snafuNum[longer[i]]
		if i < len(shorter) {
			sum += snafuNum[shorter[i]]
		}
		ans[i] += sum
		if ans[i] > 2 {
			ans[i] -= 5
			ans[i+1]++
		} else if ans[i] < -2 {
			ans[i] += 5
			ans[i+1]--
		}
	}

	for ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}

	snafu := ""
	for _, a := range ans {
		snafu = snafuChar[a] + snafu
	}
	return snafu

}

func part1(numSnafu []string, debug bool) string {
	s := ""
	for _, s1 := range numSnafu {
		s = snafuAdd(s, s1)
	}
	return s
}

func Solve(input string, debug bool, task int) (string, string) {
	ret1 := ""
	numbersSnafu := strings.Split(input, "\n")
	if task != 1 {
		ret1 = part1(numbersSnafu, debug)
	}
	fmt.Println("Result: ", ret1)

	return ret1, ""
}
