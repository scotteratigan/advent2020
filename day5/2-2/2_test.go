package main

import (
	"testing"
)

// Test functions must start with the word "Test" in a file named "xxx_test.go"
// then simply 'go test' at the command line to run

func TestConvertInt(t *testing.T) {
	var tests = []struct {
		input    string
		trueChar byte
		expected int
	}{
		{"ABBA", byte('B'), 6},
		{"BFFFBBF", byte('B'), 70},
		{"FFFBBBF", byte('B'), 14},
		{"BBFFBBF", byte('B'), 102},
		{"RRR", byte('R'), 7},
		{"RLL", byte('R'), 4},
	}

	for _, tc := range tests {
		output := ConvertInt(tc.input, tc.trueChar)
		if output != tc.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", tc.input, tc.expected, output)
		}
	}
}

func TestGetSeat(t *testing.T) {
	s := GetSeat("BBFFBBFRLL")
	if s.row != 102 {
		t.Error("Row is wrong, expected 102 but received", s.row)
	}
	if s.col != 4 {
		t.Error("Col is wrong, expected 4 but received", s.col)
	}
}
