// Each policy actually describes two positions in the password,
// where 1 means the first character, 2 means the second character,
// and so on. (Be careful; Toboggan Corporate Policies have no
// concept of "index zero"!) Exactly one of these positions must
// contain the given letter. Other occurrences of the letter are
// irrelevant for the purposes of policy enforcement.
// password.txt contains the passwords

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// todo: revisit with goroutines
func main() {
	validPassCount := 0
	lines := readFileIntoStrArr("passwords.txt", "\n")
	fullRegex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	for _, line := range lines {
		if isValidPassword(line, fullRegex) {
			validPassCount++
		}
	}
	fmt.Println("Valid passwords:", validPassCount)

}

func isValidPassword(line string, reg *regexp.Regexp) bool {
	matchArr := reg.FindStringSubmatch(line)
	firstIndex := toInt(matchArr[1]) - 1 // -1 because input is 1-indexed but strings are 0-indexed
	secondIndex := toInt(matchArr[2]) - 1
	char := matchArr[3]
	pw := matchArr[4]
	matches := 0
	if string(pw[firstIndex]) == char || string(pw[secondIndex]) == char {
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
