package main

import (
	"flag"
	"fmt"
	"os"
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
	"gitlab.com/mnlphlp/aoc22/day15"
	"gitlab.com/mnlphlp/aoc22/day16"
)

func notImplemented(day int) func(bool) (string, string) {
	return func(b bool) (string, string) {
		fmt.Printf("day %v is not implemented in go\n", day)
		return "not", "implemented"
	}
}

var dayFuncs = [...]func(bool) (string, string){
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
	day15.Solve,
	day16.Solve,
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
	updateReadme := flag.Bool("readme", false, "updateReadme")
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
		start := time.Now()
		res1, res2 := dayFuncs[day-1](*test)
		times = append(times, float32(time.Since(start).Microseconds())/1000)
		results1 = append(results1, res1)
		results2 = append(results2, res2)
	}
	overall := float32(time.Since(start).Microseconds()) / 1000

	results := "## Results:\n"
	results += "day | result 1        | result 2        | time (ms) | % of overall time\n"
	results += "--: | :-------------: | :--------------:| --------: | :--------\n"
	for i, day := range days {
		results += fmt.Sprintf("%3d | %-15s | %-15s | %9.2f | %5.2f %%\n", day, results1[i], results2[i], times[i], times[i]/overall*100)
	}
	results += fmt.Sprintf("\nOverall Time: %.2f s\n", overall/1000)
	if *updateReadme {
		content, _ := os.ReadFile("README.md")
		startIndex := strings.Index(string(content), "## Results:\n")
		endIndex := strings.Index(string(content), "Overall Time:")
		start := []byte{}
		if startIndex >= 0 {
			start = content[:startIndex]
		}
		end := []byte{}
		if endIndex >= 0 {
			endIndex += strings.Index(string(content[endIndex:]), "\n")
			end = content[endIndex:]
		}
		os.WriteFile("README.md", append(start, append([]byte(results), end...)...), 0644)
	}
	fmt.Printf("\n%s", results)
}
