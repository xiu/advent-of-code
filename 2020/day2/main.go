package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// to string
	lines := strings.Split(string(data), "\n")

	// to int
	var validcount int = 0
	var validcount2 int = 0

	for _, line := range lines {
		r := regexp.MustCompile(`^(?P<floor>\d*)-(?P<ceiling>\d*) (?P<letter>\w{1}): (?P<password>.*)`)
		result := r.FindStringSubmatch(line)

		count := strings.Count(result[4], result[3])

		floor, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		ceiling, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}

		if count <= ceiling && count >= floor {
			validcount++
		}

		var hits int = 0
		passwordAsRune := []rune(result[4])
		if passwordAsRune[floor-1] == rune(result[3][0]) {
			hits++
		}

		if passwordAsRune[ceiling-1] == rune(result[3][0]) {
			fmt.Printf("%c %c\n", passwordAsRune[ceiling-1], rune(result[3][0]))
			hits++
		}

		if hits == 1 {
			validcount2++
		}
	}

	fmt.Printf("RESULT 1: %d valid passwords\n", validcount)
	fmt.Printf("RESULT 2: %d valid passwords\n", validcount2)

}
