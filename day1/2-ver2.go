// Same algorithm as last version, but this one reads numbers from a file

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	sum := 2020
	nums := readNumsFromFile("nums.txt", "\n")
	n1, n2, n3, success := findThreeNumsThatAddToSum(sum, nums)
	if !success {
		fmt.Println("No combination of three numbers in the list added to the sum of", sum)
		return
	}
	prod := n1 * n2 * n3
	fmt.Println(n1, n2, n3, prod)
}

func findThreeNumsThatAddToSum(targetSum int, nums []int) (int, int, int, bool) {
	numsLen := len(nums)
  for i := 0; i < numsLen; i++ {
		iVal := nums[i]
    for j := i + 1; j < numsLen; j++ {
			jVal := nums[j]
			for k := j + 1; k < numsLen; k++ {
				kVal := nums[k]
				if sum := iVal + jVal + kVal; sum == targetSum {
					return iVal, jVal, kVal, true
				}
			}
		}
	}
	return 0, 0, 0, false
}

func readNumsFromFile(fileName, separator string) []int {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), separator)
	var nums = []int{}
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			break // just ignore lines that don't parse properly (empty, non-numbers, etc)
		}
		nums = append(nums, num)
	}
	return nums
}
