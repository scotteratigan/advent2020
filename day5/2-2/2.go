package main

import (
	"fmt"
	"math"
)

// Seat contains coords in row, col
type Seat struct {
	row int
	col int
}

func main() {
	i := "BBFFBBFRLL"
	s := GetSeat(i)
	fmt.Println(s)
}

// GetSeat converts string to Seat{row, col}
func GetSeat(str string) Seat {
	// Row info is first 7 chars, Col is next 3 chars
	// Example: BFFFBBFRRR
	rText := str[:7]
	cText := str[7:]
	seat := Seat{
		row: ConvertInt(rText, byte('B')),
		col: ConvertInt(cText, byte('R')),
	}
	return seat
}

// ConvertInt converts a pseudo-binary string array to an int value
func ConvertInt(str string, trueCharVal byte) int {
	// trueCharVal is the value to consider true
	// places are right to left, as usual
	// "1011", "1" => 1 + 2 + 0 + 8 = 11
	// "ABBA", "B" => 0 + 2 + 4 + 0 = 6
	n := 0
	power := 0.0
	for i := len(str) - 1; i >= 0; i-- {
		c := str[i]
		if trueCharVal == c {
			n += int(math.Pow(2, power))
		}
		power++
	}
	return n
}
