package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gitlab.com/mnlphlp/aoc22/day01"
	"gitlab.com/mnlphlp/aoc22/day02"
	"gitlab.com/mnlphlp/aoc22/day03"
	"gitlab.com/mnlphlp/aoc22/day04"
	"gitlab.com/mnlphlp/aoc22/day05"
	"gitlab.com/mnlphlp/aoc22/day06"
	"gitlab.com/mnlphlp/aoc22/day07"
	"gitlab.com/mnlphlp/aoc22/day08"
	"gitlab.com/mnlphlp/aoc22/day09"
	"gitlab.com/mnlphlp/aoc22/day10"
	"gitlab.com/mnlphlp/aoc22/day11"
	"gitlab.com/mnlphlp/aoc22/day12"
	"gitlab.com/mnlphlp/aoc22/day13"
	"gitlab.com/mnlphlp/aoc22/day14"
)

func notImplemented(day int) func(bool) (string, string, time.Duration) {
	return func(b bool) (string, string, time.Duration) {
		start := time.Now()
		fmt.Printf("day %v is not implemented in go\n", day)
		return "       not", "done in go", time.Since(start)
	}
}

var dayFuncs = [...]func(bool) (string, string, time.Duration){
	day01.Solve,
	day02.Solve,
	day03.Solve,
	day04.Solve,
	day05.Solve,
	day06.Solve,
	day07.Solve,
	day08.Solve,
	day09.Solve,
	day10.Solve,
	day11.Solve,
	day12.Solve,
	day13.Solve,
	day14.Solve,
	notImplemented(15),
	notImplemented(16),
	notImplemented(17),
	notImplemented(18),
	notImplemented(19),
	notImplemented(20),
	notImplemented(21),
	notImplemented(22),
	notImplemented(23),
	notImplemented(24),
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

	results1 := make([]string, 0, len(days))
	results2 := make([]string, 0, len(days))
	times := make([]float32, 0, len(days))

	start := time.Now()
	for _, day := range days {
		fmt.Printf("\n##################\ncalculating day %d \n##################\n", day)
		res1, res2, time := dayFuncs[day-1](*test)
		results1 = append(results1, res1)
		results2 = append(results2, res2)
		times = append(times, float32(time.Microseconds())/1000)
	}
	overall := float32(time.Since(start).Microseconds()) / 1000

	fmt.Println("\n\n###########\n# Results #\n###########")
	for i, day := range days {
		fmt.Printf("day %2d:  %-15s  %-15s  (%6.2f ms / %5.2f %%)\n", day, results1[i], results2[i], times[i], times[i]/overall*100)
	}
	fmt.Printf("Overall Time: %.2f s\n", overall/1000)

}
