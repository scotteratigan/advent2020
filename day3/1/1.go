package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	course := readFileIntoStrArr("course.txt", "\n")
	courseWidth := len(course[0])
	trees := 0
	x := 0
	for y := 1; y < len(course); y++ {
		x = (x + 3) % courseWidth
		fmt.Println(course[y])
		position := string(course[y][x])
		fmt.Println(x, position)
		if position == "#" {
			trees++
		}
	}
	fmt.Println("trees:", trees)
}

func readFileIntoStrArr(fileName, separator string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), separator)
	return lines
}
