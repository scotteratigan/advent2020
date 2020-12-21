// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID) // optional field

// Passports are separated by blank lines:
//
// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
// byr:1937 iyr:2017 cid:147 hgt:183cm
//

// Going to read this into a map because I suspect part 2 will involve checking values

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// type passport struct {
// 	byr string
// 	iyr string
// 	eyr string
// 	hgt string
// 	hcl string
// 	ecl string
// 	pid string
// 	cid string
// }
// Not using struct because I can't use: p[key] = value
// invalid operation - struct does not support indexing
// so I'd have to have a switch statement... maybe that's ok

func main() {
	fmt.Println("Running passport validation...")
	input := readFileIntoStr("input.txt")
	passportStrings := splitStrByEmptyLines(input)
	passports := make([]map[string]string, 0)
	for _, str := range passportStrings {
		p := makePassport(str)
		passports = append(passports, p)
	}
	validPassports := 0
	for _, p := range passports {
		if passportIsValid(p) {
			validPassports++
		}
	}
	fmt.Println("Valid passports:", validPassports)
}

func passportIsValid(p map[string]string) bool {
	if _, ok := p["byr"]; !ok {
		return false
	}
	if _, ok := p["iyr"]; !ok {
		return false
	}
	if _, ok := p["eyr"]; !ok {
		return false
	}
	if _, ok := p["hgt"]; !ok {
		return false
	}
	if _, ok := p["hcl"]; !ok {
		return false
	}
	if _, ok := p["ecl"]; !ok {
		return false
	}
	if _, ok := p["pid"]; !ok {
		return false
	}
	// if _, ok := p["cid"]; !ok {
	// 	return false
	// }
	return true
}

func makePassport(input string) map[string]string {
	p := make(map[string]string)
	for _, kvPair := range strings.Split(input, " ") {
		entry := strings.Split(kvPair, ":")
		key := entry[0]
		value := entry[1]
		p[key] = value
	}
	return p
}

func readFileIntoStr(fileName string) string {
	str, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func splitStrByEmptyLines(input string) []string {
	// alternately, I could do a regex split
	result := []string{}
	buff := ""
	for _, l := range strings.Split(input, "\n") {
		if len(l) > 0 {
			buff = buff + l + " "
		} else {
			result = append(result, buff[:len(buff)-1])
			buff = ""
		}
	}
	return result
}
