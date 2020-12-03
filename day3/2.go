package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	course := readFileIntoStrArr("course.txt", "\n")
	treeProduct := treesOnPath(1, 1, course) * treesOnPath(3, 1, course) * treesOnPath(5, 1, course) * treesOnPath(7, 1, course) * treesOnPath(1, 2, course)
	fmt.Println(treeProduct)
}

func treesOnPath(right, down int, course []string) int {
	courseLength := len(course)
	courseWidth := len(course[0])
	treeCount := 0
	x := 0
	for y := down; y < courseLength; y += down {
		x = (x + right) % courseWidth
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
