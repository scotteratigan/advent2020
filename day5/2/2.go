// BFFFBBFRRR
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const numRows = 128
const numColumns = 8

func main() {
	lines := readFileIntoStrArr("input.txt", "\n")
	seats := [numRows][numColumns]bool{}
	for _, l := range lines {
		r := calcRow(l)
		c := calcCol(l)
		seats[r][c] = true
	}

	// now find the missing seat:
	for r := 1; r < numRows; r++ {
		fmt.Println("Checking row:", r)
		row := seats[r][:]
		prevRow := seats[r-1][:]
		emptySeatIndex := emptySeatInRow(row)
		if rowIsOccupied(prevRow) && emptySeatIndex >= 0 {
			fmt.Println("Empty seat in row:", r, emptySeatIndex)
			fmt.Println("Seat ID:", (r*8 + emptySeatIndex))
			break
		}
	}
}

func rowIsOccupied(row []bool) bool {
	for _, seatOccupied := range row {
		if seatOccupied {
			return true
		}
	}
	return false
}

func emptySeatInRow(row []bool) int {
	for seatIndex, seatOccupied := range row {
		if !seatOccupied {
			return seatIndex
		}
	}
	return -1
}

func calcCol(input string) int {
	min := 0
	max := numColumns - 1
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
	max := numRows - 1
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

// BFFFBBF - 70
// 1000110
// from end:
// 2 + 4 + 64 = 70
