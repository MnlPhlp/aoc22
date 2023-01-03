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

/*
	func (f File) Move(i, m int) {
		if m == 0 {
			// no move
			return
		}
		oldPos := f[i].Pos
		newPos := (oldPos + m) % len(f)
		if m > 0 && newPos < oldPos {
			// wrapped around to the beginning
			newPos++
		} else if m < 0 {
			if newPos < 0 {
				// wrapped around to the end
				newPos = len(f) + (newPos % len(f)) - 1
			}
			if newPos == 0 {
				newPos = len(f) - 1
			}
		}
		f[i].Pos = newPos
		op := -1
		left, right := oldPos, newPos+1
		if newPos < oldPos {
			left, right = newPos-1, oldPos
			op = 1
		}
		for j := 0; j < len(f); j++ {
			if j == i {
				continue
			}
			if f[j].Pos > left && f[j].Pos < right {
				f[j].Pos += op
			}
		}

}
*/
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
	oldPos := f[i].Pos
	newPos := (oldPos + m) % len(f)
	if m > 0 && newPos < oldPos {
		// wrapped around to the beginning
		newPos++
	} else if m < 0 {
		if newPos < 0 {
			// wrapped around to the end
			newPos = len(f) + (newPos % len(f)) - 1
		}
		if newPos == 0 {
			newPos = len(f) - 1
		}
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
		} else if m > 0 && newPos < oldPos {
			// move right with wrapping
			if pos >= newPos && pos < oldPos {
				f[j].Pos++
			}
		} else if m < 0 && newPos < oldPos {
			// move left without wrapping
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

func mix(file File, debug bool) File {
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
	return file
}

func Solve(input string, debugFlag bool, task int) (string, string) {
	res1, res2 := 0, 0
	file := parseInput(input)
	if debugFlag {
		fmt.Printf("file: %v\n", file)
	}
	if task != 2 {
		mixed := mix(file, debugFlag)
		if debugFlag {
			fmt.Printf("mixed: %v\n", mixed)
		}
		res1 = util.Sum(mixed.Get(1000, 2000, 3000)...)
	}
	if task != 1 {
		res2 = 0
	}
	return strconv.Itoa(res1), strconv.Itoa(res2)
}
