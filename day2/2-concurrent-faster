// Re-implementation of 2.go using concurrency
// Definitely much faster broken up into batches, doesn't really matter unless dataset is large
// This implementation is about 3x as fast (tested with 5 million records) compared to the single-thread approach.
// The ideal number of batches would depend on the available cores I imagine.

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	var validPassCount uint64
	var wg sync.WaitGroup
	lines := readFileIntoStrArr("p2.txt", "\n")
	// lines := readFileIntoStrArr("passwords.txt", "\n")
	numLines := len(lines)
	numSlices := 8
	perSlice := numLines / (numSlices - 1) // minus 1 because of the remainder
	fullRegex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	for batch := 0; batch < numSlices; batch++ {
		minIndex := batch * perSlice
		maxIndex := min(((batch + 1) * perSlice), (numLines)) // not subtracting 1 because the second arg of a slice range is exclusive
		// Create a slice of the full dataset for each goroutine to process:
		slice := lines[minIndex:maxIndex]
		wg.Add(1)
		go func(sl []string) {
			for _, line := range sl {
				if isValidPassword(line, fullRegex) {
					atomic.AddUint64(&validPassCount, 1)
				}
			}
			wg.Done()
		}(slice)
	}
	wg.Wait()
	fmt.Println("Valid passwords:", validPassCount)
}

func isValidPassword(line string, reg *regexp.Regexp) bool {
	matchArr := reg.FindStringSubmatch(line)
	firstIndex := toInt(matchArr[1]) - 1 // -1 because input is 1-indexed but strings are 0-indexed
	secondIndex := toInt(matchArr[2]) - 1
	char := []byte(matchArr[3])[0] // convert to byte so that it can be compared to the string index
	pw := matchArr[4]
	matches := 0
	if pw[firstIndex] == char {
		matches++
	}
	if pw[secondIndex] == char {
		matches++
	}
	return matches == 1 // only 1 char should match, not 0 or 2
}

func toInt(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic("Unable to convert a number, aborting.")
	}
	return num
}

func readFileIntoStrArr(fileName, separator string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), separator)
	return lines
}

func min(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	}
	return n2
}
