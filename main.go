package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/mnlphlp/aoc22/day08"
	"gitlab.com/mnlphlp/aoc22/day09"
	"gitlab.com/mnlphlp/aoc22/day10"
	"gitlab.com/mnlphlp/aoc22/day11"
)

func notImplemented(day int) func(bool) {
	return func(b bool) {
		fmt.Printf("day %v is not implemented in go\n", day)
	}
}

var dayFuncs = [...]func(bool){
	notImplemented(1),
	notImplemented(2),
	notImplemented(3),
	notImplemented(4),
	notImplemented(5),
	notImplemented(6),
	notImplemented(7),
	day08.Solve,
	day09.Solve,
	day10.Solve,
	day11.Solve,
}

func main() {
	dayStr := flag.String("d", "", "day")
	daysString := flag.String("days", "", "days")
	test := flag.Bool("t", false, "test")
	flag.Parse()
	*dayStr = strings.Trim(*dayStr, "day.go")
	day, _ := strconv.Atoi(*dayStr)
	days := []int{}
	if day != 0 {
		days = append(days, day)
	}
	if *daysString != "" {
		for _, day := range strings.Split(*daysString, ",") {
			d, _ := strconv.Atoi(day)
			days = append(days, d)
		}
	}
	// if nothin is set, run all days
	if len(days) == 0 {
		for i := 1; i <= len(dayFuncs); i++ {
			days = append(days, i)
		}
	}

	fmt.Printf("calculating days: %v \n", days)

	for _, day := range days {
		fmt.Printf("\n##################\ncalculating day %d \n##################\n", day)
		dayFuncs[day-1](*test)
	}

}
