package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

type slope struct {
	right int
	down  int
}

func main() {
	course := readFileIntoStrArr("course.txt", "\n")
	// An array of problems and answers of the same length. Considered map but a struct index seemed like the wrong idea.
	var problemSet = [5]slope{
		slope{right: 1, down: 1},
		slope{right: 3, down: 1},
		slope{right: 5, down: 1},
		slope{right: 7, down: 1},
		slope{right: 1, down: 2},
	}
	var answerSet = [5]int{}
	var wg sync.WaitGroup
	// No need to worry about mutex exclusion if each goroutine has it's own dedicated index to write to:
	for i := range problemSet {
		wg.Add(1)
		go func(i int) {
			answerSet[i] = treesOnPath(problemSet[i], course)
			fmt.Println(problemSet[i], answerSet[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	prod := 1
	for _, sum := range answerSet {
		prod *= sum
	}
	fmt.Println(prod)
}

func treesOnPath(sl slope, course []string) int {
	courseLength := len(course)
	courseWidth := len(course[0])
	treeCount := 0
	x := 0
	for y := sl.down; y < courseLength; y += sl.down {
		x = (x + sl.right) % courseWidth
		if string(course[y][x]) == "#" {
			treeCount++
		}
	}
	return treeCount
}

func readFileIntoStrArr(fileName, separator string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), separator)
	return lines
}
