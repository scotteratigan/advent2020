// Re-implementation of 2.go using concurrency
// Not seeing a real speed improvement, I should probably break up the data into chunks to process in relatively fewer goroutines

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
	fullRegex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	for _, line := range lines {
		wg.Add(1)
		go func(l string) {
			if isValidPassword(l, fullRegex) {
				atomic.AddUint64(&validPassCount, 1)
			}
			wg.Done()
		}(line)
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
	if pw[firstIndex] == char || pw[secondIndex] == char {
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
