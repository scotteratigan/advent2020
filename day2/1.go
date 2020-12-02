// Each line gives the password policy and then the password.
// The password policy indicates the lowest and highest number of
// times a given letter must appear for the password to be valid.
// For example, 1-3 a means that the password must contain a at
// least 1 time and at most 3 times.
// Example: `2-4 r: prrmspx` is valid because it contains between 2 and 4 'r' chars
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
			validPassCount += 1
		}
	}
	fmt.Println("Valid passwords:", validPassCount)

}

func isValidPassword(line string, reg *regexp.Regexp) (bool) {
	matchArr := reg.FindStringSubmatch(line)
	min := toInt(matchArr[1])
	max := toInt(matchArr[2])
	char := matchArr[3]
	pw := matchArr[4]
	count := strings.Count(pw, char)
	return count >= min && count <= max
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
