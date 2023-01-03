package day20

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mnlphlp/aoc22/util"
)

type Item struct {
	Val  int
	prev *Item
	next *Item
}
type File struct {
	items []*Item
	start *Item
}

func (f File) Get(is ...int) []int {
	ints := make([]int, len(is))
	zero := f.start
	for _, v := range f.items {
		if v.Val == 0 {
			zero = v
		}
	}
	for i, idx := range is {
		steps := idx % len(f.items)
		p := zero
		for j := 0; j < steps; j++ {
			p = p.next
		}
		ints[i] = p.Val
	}
	return ints
}

func (f File) String() string {
	ints := make([]int, len(f.items))
	cur := f.start
	for i := range f.items {
		ints[i] = cur.Val
		cur = cur.next
	}
	return fmt.Sprintf("%v", ints)
}

func (f *File) Move(i int, m int) {
	// avoid loops
	m = m % (len(f.items) - 1)
	if m == 0 {
		// no move
		return
	}
	cur := f.items[i]
	right := cur
	if cur == f.start {
		f.start = cur.next
	}
	if m < 0 {
		// move backwards
		for i := 0; i < util.Abs(m); i++ {
			right = right.prev
		}
	} else {
		right = right.next
		// move forwards
		for i := 0; i < m; i++ {
			right = right.next
		}
		if right == f.start {
			f.start = cur
		}
	}
	// move item to left of "right"
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	cur.prev = right.prev
	cur.next = right
	right.prev.next = cur
	right.prev = cur
}

func (f File) HasDuplicatePos() bool {
	next := make(map[*Item]bool)
	prev := make(map[*Item]bool)
	for _, v := range f.items {
		if next[v.next] || prev[v.prev] {
			return true
		}
		next[v.next] = true
		prev[v.prev] = true
	}
	return false
}

func (file *File) Reset() {
	file.start = file.items[0]
	for i, v := range file.items {
		if i == 0 {
			v.prev = file.items[len(file.items)-1]
		} else {
			v.prev = file.items[i-1]
		}
		if i == len(file.items)-1 {
			v.next = file.items[0]
		} else {
			v.next = file.items[i+1]
		}
	}
}

func parseInput(input string) File {
	file := File{
		items: make([]*Item, 0),
	}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		file.items = append(file.items, &Item{
			Val: util.ParseInt(line),
		})
	}
	file.Reset()
	return file
}

func mix(file File, debug bool, key int, iterations int) File {
	if key != 1 {
		for i := 0; i < len(file.items); i++ {
			old := util.Abs(file.items[i].Val)
			file.items[i].Val *= key
			if util.Abs(file.items[i].Val) < old {
				panic("Overflow")
			}
		}
	}
	for it := 0; it < iterations; it++ {
		for i, v := range file.items {
			file.Move(i, v.Val)
			if debug {
				fmt.Println(v)
				fmt.Println(file)
				if file.HasDuplicatePos() {
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
			file.Reset()
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
