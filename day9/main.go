/*
--- Day 9: Encoding Error ---
With your neighbor happily enjoying their video game, you turn your attention to an open data port on the little screen in the seat in front of you.

Though the port is non-standard, you manage to connect it to your computer through the clever use of several paperclips. Upon connection, the port outputs a series of numbers (your puzzle input).

The data appears to be encrypted with the eXchange-Masking Addition System (XMAS) which, conveniently for you, is an old cypher with an important weakness.

XMAS starts by transmitting a preamble of 25 numbers. After that, each number you receive should be the sum of any two of the 25 immediately previous numbers. The two numbers will have different values, and there might be more than one such pair.

For example, suppose your preamble consists of the numbers 1 through 25 in a random order. To be valid, the next number must be the sum of two of those numbers:

26 would be a valid next number, as it could be 1 plus 25 (or many other pairs, like 2 and 24).
49 would be a valid next number, as it is the sum of 24 and 25.
100 would not be valid; no two of the previous 25 numbers sum to 100.
50 would also not be valid; although 25 appears in the previous 25 numbers, the two numbers in the pair must be different.
Suppose the 26th number is 45, and the first number (no longer an option, as it is more than 25 numbers ago) was 20. Now, for the next number to be valid, there needs to be some pair of numbers among 1-19, 21-25, or 45 that add up to it:

26 would still be a valid next number, as 1 and 25 are still within the previous 25 numbers.
65 would not be valid, as no two of the available numbers sum to it.
64 and 66 would both be valid, as they are the result of 19+45 and 21+45 respectively.
Here is a larger example which only considers the previous 5 numbers (and has a preamble of length 5):

35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
(example.txt)
In this example, after the 5-number preamble, almost every number is the sum of two of the previous 5 numbers; the only number that does not follow this rule is 127.

The first step of attacking the weakness in the XMAS data is to find the first number in the list (after the preamble) which is not the sum of two of the 25 numbers before it. What is the first number that does not have this property?
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// nums := readFileIntoIntArr("example.txt")
	// invalid := findInvalidEntry(nums, 5)
	nums := readFileIntoIntArr("input.txt")
	invalid := findInvalidEntry(nums, 25)
	fmt.Println("#1 - Invalid:", invalid)
	// start and end are the beginning and end index of the slice of nums that add up to the invalid entry:
	start, end := getSequenceThatSumsToTotal(nums, invalid)
	min := minIntSlice(nums[start : end+1])
	max := maxIntSlice(nums[start : end+1])
	secretSum := min + max
	// prompt asks for the sum of the smallest and largest in this range
	fmt.Println("#2 - Secret Sum:", secretSum)
}

func minIntSlice(nums []int) int {
	var min int
	for i, num := range nums {
		if num < min || i == 0 {
			min = num
		}
	}
	return min
}

func maxIntSlice(nums []int) int {
	var max int
	for i, num := range nums {
		if num > max || i == 0 {
			max = num
		}
	}
	return max
}

func getSequenceThatSumsToTotal(nums []int, targetSum int) (int, int) {
	// example of a "sliding window" problem
	i1 := 0
	i2 := 1
	for {
		numSlice := nums[i1 : i2+1]
		sum := getSumOfSlice(numSlice)
		if sum == targetSum {
			return i1, i2
		}
		if sum < targetSum {
			// if sum is too small, try adding an additional number
			i2++
		} else {
			// if sum is too big, start over with start position 1 higher, and reset the end index to 1 greater than start index
			i1++
			i2 = i1 + 1
		}
		if i2 == len(nums) {
			// in this case, we're at the end, so this isn't happening
			return -1, -1
		}
	}
}

func getSumOfSlice(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func findInvalidEntry(nums []int, preamble int) int {
	for i := preamble; i < len(nums); i++ {
		currentNum := nums[i]
		availableToSum := nums[i-preamble : i]
		possible := sumIsPossible(availableToSum, currentNum)
		if !possible {
			return currentNum
		}
	}
	return -1
}

func sumIsPossible(nums []int, targetSum int) bool {
	len := len(nums)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == targetSum {
				return true
			}
		}
	}
	return false
}

func readFileIntoIntArr(fileName string) []int {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	strings := strings.Split(string(data), "\n")
	// allocate a slice the same length as the number of strings we have so we don't have to keep appending
	ints := make([]int, len(strings))
	for i, strNum := range strings {
		iVal, err := strconv.Atoi(strNum)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = iVal
	}
	return ints
}
