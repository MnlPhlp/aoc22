package day20

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mnlphlp/aoc22/util"
)

type Item struct {
	Val, Pos int
}
type File []Item

func (f File) Get(is ...int) []int {
	ints := make([]int, len(is))
	start := 0
	for _, v := range f {
		if v.Val == 0 {
			start = v.Pos
		}
	}
	for _, v := range f {
		for i, idx := range is {
			if v.Pos == (idx+start)%len(f) {
				ints[i] = v.Val
			}
		}
	}
	return ints
}

func (f File) String() string {
	ints := make([]int, len(f))
	for _, v := range f {
		ints[v.Pos] = v.Val
	}
	return fmt.Sprintf("%v", ints)
}

func (f File) Move(i int, m int) {
	if m == 0 {
		// no move
		return
	}
	m = m % (len(f) - 1)
	oldPos := f[i].Pos
	newPos := oldPos + m
	if newPos < 0 {
		// wrapped around to the end
		newPos = len(f) + (newPos % len(f)) - 1
	} else if newPos >= len(f) {
		// wrapped around to the beginning
		newPos = newPos % (len(f) - 1)
	} else if newPos == 0 {
		newPos = len(f) - 1
	} else if newPos == len(f)-1 {
		newPos = 0
	}
	if oldPos == newPos {
		return
	}
	for j := 0; j < len(f); j++ {
		if j == i {
			continue
		}
		pos := f[j].Pos
		if m > 0 && newPos > oldPos {
			// move right without wrapping
			if pos > oldPos && pos <= newPos {
				f[j].Pos--
			}
		} else if m < 0 && newPos < oldPos {
			// move left without wrapping
			if pos >= newPos && pos < oldPos {
				f[j].Pos++
			}
		} else if m > 0 && newPos < oldPos {
			// move right with wrapping
			if pos >= newPos && pos < oldPos {
				f[j].Pos++
			}
		} else if m < 0 && newPos > oldPos {
			// move left with wrapping
			if pos > oldPos && pos <= newPos {
				f[j].Pos--
			}
		}
	}
	f[i].Pos = newPos
}

func (f File) HasDuplicatePos() bool {
	seen := make(map[int]bool)
	for _, v := range f {
		if seen[v.Pos] {
			return true
		}
		seen[v.Pos] = true
	}
	return false
}

func parseInput(input string) File {
	file := make(File, 0)
	i := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		file = append(file, Item{
			Val: util.ParseInt(line),
			Pos: i,
		})
		i++
	}
	return file
}

func mix(file File, debug bool, key int, iterations int) File {
	if key != 1 {
		for i := 0; i < len(file); i++ {
			old := util.Abs(file[i].Val)
			file[i].Val *= key
			if util.Abs(file[i].Val) < old {
				panic("Overflow")
			}
		}
	}
	for it := 0; it < iterations; it++ {
		for i, v := range file {
			file.Move(i, v.Val)
			if debug {
				fmt.Println(v)
				fmt.Println(file)
				if file.HasDuplicatePos() {
					fmt.Println([]Item(file))
					panic("duplicate pos")
				}
			}
		}
	}
	return file
}

func Solve(input string, debugFlag bool, task int) (string, string) {
	res1, res2 := 0, 0
	file := parseInput(input)
	if debugFlag {
		fmt.Printf("file: %v\n", file)
		fmt.Printf("items: %v", []Item(file))
	}
	if task != 2 {
		mixed := mix(file, debugFlag, 1, 1)
		if debugFlag {
			fmt.Printf("mixed: %v\n", mixed)
		}
		res1 = util.Sum(mixed.Get(1000, 2000, 3000)...)
	}
	if task != 1 {
		if task != 2 {
			// reset file
			for i := range file {
				file[i].Pos = i
			}

		}
		mixed := mix(file, debugFlag, 811589153, 10)
		values := mixed.Get(1000, 2000, 3000)
		if debugFlag {
			fmt.Printf("mixed: %v\n", mixed)
			fmt.Printf("values: %v\n", values)
		}
		res2 = util.Sum(values...)
	}
	return strconv.Itoa(res1), strconv.Itoa(res2)
}
