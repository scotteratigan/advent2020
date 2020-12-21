package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	inputStr := readFileIntoStr("input.txt")
	inputArr := splitStrByEmptyLines(inputStr)
	newLineRegExp := regexp.MustCompile(`\r?\n`)
	sumAllUnanimousAnswers := 0

	for _, answerGroup := range inputArr {
		answersMap := make([]map[byte]bool, 0)
		answers := newLineRegExp.Split(answerGroup, -1)
		for _, answer := range answers {
			answerSet := createMapOfAnswers(answer)
			answersMap = append(answersMap, answerSet)
		}
		unanimous := answersPresentInAll(answersMap)
		fmt.Println("unanimous:", unanimous)
		sumAllUnanimousAnswers += unanimous
	}
	fmt.Println("sumAllUnanimousAnswers:", sumAllUnanimousAnswers)
}

// numbers here are wrong, getting lost/confused in this logic here
// need to test smaller pieces

func answersPresentInAll(mapArr []map[byte]bool) int {
	unanimousCount := 0
	// iterate through any answer set, and count values that are present in all others
	for key := range mapArr[0] {
		// fmt.Println("key:", key)
		unanimous := true
		for i := 0; i < len(mapArr); i++ {
			_, ok := mapArr[i][key]
			if !ok {
				unanimous = false
				break
			}
		}
		if unanimous {
			unanimousCount++
		}
	}
	return unanimousCount
}

func createMapOfAnswers(answers string) map[byte]bool {
	uniqueAns := map[byte]bool{}
	for _, letter := range answers {
		uniqueAns[byte(letter)] = true
	}
	return uniqueAns
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
	split := emptyLineRegExp.Split(input, -1)
	return split
}
