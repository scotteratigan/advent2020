package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	maxId := 0
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "")

		codePos := 0
		lowerPosition := ""
		binaryPositionCheck := func(i int) bool {
			result := false
			if split[codePos] == lowerPosition {
				result = true
			}
			codePos += 1
			return result
		}

		lowerPosition = "F"
		row := sort.Search(127, binaryPositionCheck)
		lowerPosition = "L"
		column := sort.Search(7, binaryPositionCheck)
		id := (row * 8) + column
		if id > maxId {
			maxId = id
		}
	}

	fmt.Println(maxId)
}
