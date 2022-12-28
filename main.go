package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mnlphlp/aoc22/day01"
	"github.com/mnlphlp/aoc22/day02"
	"github.com/mnlphlp/aoc22/day03"
	"github.com/mnlphlp/aoc22/day04"
	"github.com/mnlphlp/aoc22/day05"
	"github.com/mnlphlp/aoc22/day06"
	"github.com/mnlphlp/aoc22/day07"
	"github.com/mnlphlp/aoc22/day08"
	"github.com/mnlphlp/aoc22/day09"
	"github.com/mnlphlp/aoc22/day10"
	"github.com/mnlphlp/aoc22/day11"
	"github.com/mnlphlp/aoc22/day12"
	"github.com/mnlphlp/aoc22/day13"
	"github.com/mnlphlp/aoc22/day14"
	"github.com/mnlphlp/aoc22/day15"
	"github.com/mnlphlp/aoc22/day16"
	"github.com/mnlphlp/aoc22/day17"
)

func notImplemented(day int) func(bool, int) (string, string) {
	return func(b bool, i int) (string, string) {
		fmt.Printf("day %v is not implemented in go\n", day)
		return "not", "implemented"
	}
}

func wrap(f func(bool) (string, string)) func(bool, int) (string, string) {
	return func(b bool, i int) (string, string) {
		return f(b)
	}
}

var dayFuncs = [...]func(bool, int) (string, string){
	wrap(day01.Solve),
	wrap(day02.Solve),
	wrap(day03.Solve),
	wrap(day04.Solve),
	wrap(day05.Solve),
	wrap(day06.Solve),
	wrap(day07.Solve),
	wrap(day08.Solve),
	wrap(day09.Solve),
	wrap(day10.Solve),
	wrap(day11.Solve),
	wrap(day12.Solve),
	wrap(day13.Solve),
	wrap(day14.Solve),
	wrap(day15.Solve),
	day16.Solve,
	day17.Solve,
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
	task := flag.Int("task", 0, "task (0=both, 1=task1, 2=task2)")
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

	results1 := make([]string, len(days))
	results2 := make([]string, len(days))
	times := make([]float32, len(days))

	start := time.Now()
	for i, day := range days {
		fmt.Printf("\n##################\ncalculating day %d \n##################\n", day)
		start := time.Now()
		res1, res2 := dayFuncs[day-1](*test, *task)
		times[i] = float32(time.Since(start).Microseconds()) / 1000
		results1[i] = res1
		results2[i] = res2
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
