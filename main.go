package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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
	"github.com/mnlphlp/aoc22/util"
)

func notImplemented(day int) func(string, bool, int) (string, string) {
	return func(str string, b bool, i int) (string, string) {
		fmt.Printf("day %v is not implemented in go\n", day)
		return "not", "implemented"
	}
}

func wrap(f func(bool) (string, string)) func(string, bool, int) (string, string) {
	return func(s string, b bool, i int) (string, string) {
		return f(b)
	}
}

var dayFuncs = [...]func(string, bool, int) (string, string){
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
	day17.Solve,
	notImplemented(18),
	notImplemented(19),
	notImplemented(20),
	notImplemented(21),
	notImplemented(22),
	notImplemented(23),
	notImplemented(24),
}

func capLength(str string, length int) string {
	if len(str) > length {
		return str[:length-3] + "..."
	}
	return str
}

func calcDay(day int, i int, results1 []string, results2 []string, times []time.Duration, test bool, task int, debug bool) {
	fmt.Printf("\n##################\ncalculating day %d \n##################\n", day)
	start := time.Now()
	input := util.ReadInput(day, test)
	res1, res2 := dayFuncs[day-1](input, test || debug, task)
	times[i] = time.Since(start)
	results1[i] = res1
	results2[i] = res2
}

func main() {
	dayStr := flag.String("d", "", "day")
	daysString := flag.String("days", "", "days")
	test := flag.Bool("t", false, "test")
	task := flag.Int("task", 0, "task (0=both, 1=task1, 2=task2)")
	updateReadme := flag.Bool("readme", false, "updateReadme")
	parallel := flag.Bool("p", false, "parallel")
	debug := flag.Bool("debug", false, "debug")
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
	times := make([]time.Duration, len(days))

	start := time.Now()
	if *parallel {
		wg := sync.WaitGroup{}
		for i, day := range days {
			wg.Add(1)
			go func(i int, day int) {
				defer wg.Done()
				calcDay(day, i, results1, results2, times, *test, *task, *debug)
			}(i, day)
		}
		wg.Wait()
	} else {
		for i, day := range days {
			calcDay(day, i, results1, results2, times, *test, *task, *debug)
		}
	}
	overall := time.Since(start)

	results := "## Results:\n"
	results += "day | result 1        | result 2        | time (ms) | % of overall time\n"
	results += "--: | :-------------: | :--------------:| --------: | :--------\n"
	for i, day := range days {
		results += fmt.Sprintf("%3d | %-15s | %-15s | %9.2f | %5.2f %%\n",
			day,
			capLength(results1[i], 15),
			capLength(results2[i], 15),
			float32(times[i].Microseconds())/1000,
			float32(times[i].Microseconds())/float32(overall.Microseconds())*100)
	}
	results += fmt.Sprintf("\nOverall Time: %v\n", overall)
	results += fmt.Sprintf("Summed Time: %v\n", util.Sum(times))
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
