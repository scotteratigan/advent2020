// BFFFBBFRRR
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	lines := readFileIntoStrArr("input.txt", "\n")
	maxId := -1
	for _, l := range lines {
		id := calcSeatID(l)
		fmt.Println(id)
		if id > maxId {
			maxId = id
		}
	}
	fmt.Println("Max:", maxId)
}

func calcSeatID(input string) int {
	r := calcRow(input)
	c := calcCol(input)
	return r*8 + c
}

func calcCol(input string) int {
	min := 0
	max := 7
	mid := 0
	for _, c := range input {
		mid = (max + min) / 2
		switch c {
		case rune('L'):
			max = mid
		case rune('R'):
			if (max+min)%2 > 0 {
				mid++
			}
			min = mid
		}
	}
	return mid
}

func calcRow(input string) int {
	min := 0
	max := 127
	mid := 0
	for _, c := range input {
		mid = (max + min) / 2
		switch c {
		case rune('F'):
			max = mid
		case rune('B'):
			if (max+min)%2 > 0 {
				mid++
			}
			min = mid
		}
	}
	return mid
}

func readFileIntoStrArr(fileName, separator string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), separator)
	return lines
}
