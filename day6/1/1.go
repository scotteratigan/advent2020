package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	inputStr := readFileIntoStr("input.txt")
	inputArr := splitStrByEmptyLines(inputStr)
	fmt.Println(inputArr)
	sum := 0
	for _, answers := range inputArr {
		uniqueCount := getUniqueCountInGroup(answers)
		fmt.Println(answers, uniqueCount)
		sum += uniqueCount
	}
	fmt.Println("Sum:", sum)
}

func getUniqueCountInGroup(answers string) int {
	uniqueAns := map[byte]bool{}
	for _, letter := range answers {
		uniqueAns[byte(letter)] = true
	}
	return len(uniqueAns)
}

func readFileIntoStr(fileName string) string {
	str, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func splitStrByEmptyLines(input string) []string {
	emptyLineRegExp := regexp.MustCompile(`\n\r?\n`)
	newLineRegExp := regexp.MustCompile(`\r?\n`)
	split := emptyLineRegExp.Split(input, -1)
	var newSplit []string

	// remove newlines in split strings:
	for _, str := range split {
		newStr := newLineRegExp.ReplaceAll([]byte(str), []byte(""))
		newSplit = append(newSplit, string(newStr))
	}
	return newSplit
}
