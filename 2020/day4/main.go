package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var requiredFields = []map[string]string{
	map[string]string{"name": "byr", "rule": "^(19[2-9][0-9]|200[0-2])$"},
	map[string]string{"name": "iyr", "rule": "^(201[0-9]|2020)$"},
	map[string]string{"name": "eyr", "rule": "^(202[0-9]|2030)$"},
	map[string]string{"name": "hgt", "rule": "^((1[5-8][0-9]|19[0-3])cm|((59|6[0-9]|7[0-6])in))$"},
	map[string]string{"name": "hcl", "rule": "^#([0-9]|[a-f]){6}$"},
	map[string]string{"name": "ecl", "rule": "^(amb|blu|brn|gry|grn|hzl|oth)$"},
	map[string]string{"name": "pid", "rule": "^[0-9]{9}$"},
	// "cid",
}
var passportList []map[string]string

func validatePassport(passport map[string]string) bool {
	for _, field := range requiredFields {
		_, ok := passport[field["name"]]
		if !ok {
			return false
		}
	}
	return true
}

func validatePassportExtended(passport map[string]string) bool {
	for _, field := range requiredFields {
		_, ok := passport[field["name"]]
		if !ok {
			return false
		}
		r := regexp.MustCompile(field["rule"])
		if !r.MatchString(passport[field["name"]]) {
			fmt.Printf("%s invalid:\t%s\n", field["name"], passport[field["name"]])
			return false
		}
		fmt.Printf("%s valid:\t%s\n", field["name"], passport[field["name"]])

	}
	return true
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// to string
	lines := strings.Split(string(data), "\n")

	currentPassport := 0
	var validPassports int
	var validPassports2 int
	passportList = append(passportList, make(map[string]string))

	for _, line := range lines {
		if line == "" {
			// new passport
			currentPassport++
			passportList = append(passportList, make(map[string]string))
			continue
		}

		items := strings.Split(line, " ")

		for _, item := range items {
			parsed := strings.Split(item, ":")
			passportList[currentPassport][parsed[0]] = parsed[1]
		}
	}

	for _, passport := range passportList {
		valid := validatePassport(passport)
		valid2 := validatePassportExtended(passport)
		if valid {
			validPassports++
		}
		if valid2 {
			validPassports2++
		}
	}
	fmt.Printf("Solution #1: %d\n", validPassports)
	fmt.Printf("Solution #2: %d\n", validPassports2)
}
