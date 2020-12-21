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

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

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
	fmt.Println("Valid passports:", validPassports, "/", len(passports))
}

// Start with 202 valid...

func passportIsValid(p map[string]string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if !validYear(p, "byr", 1920, 2002) {
		return false
	}
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if !validYear(p, "iyr", 2010, 2020) {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !validYear(p, "eyr", 2020, 2030) {
		return false
	}

	if !heightValid(p) {
		return false
	}

	if !hairColorValid(p) {
		return false
	}

	if !eyeColorValid(p) {
		return false
	}

	if !idNumValid(p) {
		return false
	}
	return true
}

func idNumValid(p map[string]string) bool {
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	id, ok := p["pid"]
	if !ok {
		return false
	}
	idReg := regexp.MustCompile(`^\d{9}$`)
	idMatch := idReg.Match([]byte(id))
	return idMatch
}

func eyeColorValid(p map[string]string) bool {
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	eyeColor := p["ecl"]
	if eyeColor != "amb" && eyeColor != "blu" && eyeColor != "brn" && eyeColor != "gry" && eyeColor != "grn" && eyeColor != "hzl" && eyeColor != "oth" {
		return false
	}
	return true
}

func hairColorValid(p map[string]string) bool {
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hairColor, ok := p["hcl"]
	if !ok {
		return false
	}
	hcReg := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	hcMatch := hcReg.Match([]byte(hairColor))
	return hcMatch
}

func heightValid(p map[string]string) bool {
	// hgt (Height) - a number followed by either cm or in.
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	heightStr, ok := p["hgt"]
	if !ok {
		return false
	}
	hReg := regexp.MustCompile(`^(\d+)([a-z]+)`)
	hMatch := hReg.FindStringSubmatch(heightStr)
	if len(hMatch) < 2 {
		return false
	}
	heightVal := hMatch[1]
	heightUnit := hMatch[2]
	// Check unit is correct:
	if heightUnit != "in" && heightUnit != "cm" {
		return false
	}
	heightInt, err := strconv.Atoi(heightVal)
	if err != nil {
		return false
	}
	// Check value is in range:
	if heightUnit == "in" {
		if !isInRange(heightInt, 59, 76) {
			return false
		}
	} else if heightUnit == "cm" {
		if !isInRange(heightInt, 150, 193) {
			return false
		}
	}
	return true
}

func isInRange(i, min, max int) bool {
	if i < min || i > max {
		return false
	}
	return true
}

func validYear(p map[string]string, passportField string, min, max int) bool {
	yearStr, ok := p[passportField]
	if !ok {
		return false
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}
	if !isInRange(year, min, max) {
		return false
	}
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
	// alternately, I could do a regex split `\n\s*\n`, but this is fine
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
