package day07

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseDirs(inputFile string) map[string]int {
	dirs := make(map[string]int)
	dirs["/"] = 0
	input, _ := os.ReadFile(inputFile)
	currentDir := "/"
	parentDirs := make([]string, 0)
	for _, l := range strings.Split(string(input), "\n") {
		if l == "$ cd /" {
			continue
		} else if strings.HasPrefix(l, "$ cd ..") {
			dirs[parentDirs[len(parentDirs)-1]] += dirs[currentDir]
			currentDir = parentDirs[len(parentDirs)-1]
			parentDirs = parentDirs[:len(parentDirs)-1]
		} else if strings.HasPrefix(l, "$ cd ") {
			parentDirs = append(parentDirs, currentDir)
			currentDir = currentDir + strings.Split(l, " ")[2] + "/"
			if _, ok := dirs[currentDir]; !ok {
				dirs[currentDir] = 0
			}
		} else if size, err := strconv.Atoi(strings.Split(l, " ")[0]); err == nil {
			dirs[currentDir] += size
		}
	}
	// get back to root
	for currentDir != "/" {
		last := len(parentDirs) - 1
		dirs[parentDirs[last]] += dirs[currentDir]
		currentDir = parentDirs[last]
		parentDirs = parentDirs[:last]
	}
	return dirs
}

func Solve(test bool) (string, string, time.Duration) {
	start := time.Now()

	inputFile := "day07/input.txt"
	if test {
		inputFile = "day07/testInput.txt"
	}
	dirs := parseDirs(inputFile)

	// Task 1
	sum := 0
	for _, val := range dirs {
		if val <= 100000 {
			sum += val
		}
	}
	fmt.Println("Total size of dirs <= 100000: ", sum)
	res1 := fmt.Sprintf("%d", sum)

	fsSize := 70000000
	fmt.Println("Total used space: ", dirs["/"])
	fmt.Println("Free space: ", fsSize-dirs["/"])
	neededSpace := 30000000 - (fsSize - dirs["/"])
	fmt.Println("Needed Space: ", neededSpace)
	// find smallest dir to delete
	minSize := fsSize
	smallestDir := ""
	for d, size := range dirs {
		if size < minSize && size >= neededSpace {
			minSize = size
			smallestDir = d
		}
	}
	fmt.Println("smallest dir to delete: ", smallestDir, minSize)
	res2 := fmt.Sprintf("%d", minSize)

	return res1, res2, time.Since(start)
}
