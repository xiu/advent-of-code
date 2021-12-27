package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Part1(lines []string) (gamma uint64, epsilon uint64) {
	sums := make([]int, len(lines[0]))

	var gstr, estr string

	for _, value := range lines {
		if value == "" {
			continue
		}
		for key, char := range value {
			valint, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			sums[key] += valint
		}
	}

	for _, value := range sums {
		if value > len(lines)/2 {
			gstr += "1"
			estr += "0"
		} else {
			gstr += "0"
			estr += "1"
		}
	}

	gamma, err := strconv.ParseUint(gstr, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilon, err = strconv.ParseUint(estr, 2, 64)
	if err != nil {
		panic(err)
	}

	return
}

func filter(slice []string, criteria byte, currentKey int) (out []string) {
	zero := []string{}
	one := []string{}

	for _, value := range slice {
		if value == "" {
			continue
		}
		if value[currentKey] == '0' {
			zero = append(zero, value)
		} else {
			one = append(one, value)
		}
	}

	if len(zero) > len(one) {
		if criteria == '1' { // o2
			out = zero
		} else { // co2
			out = one
		}
	} else if len(one) > len(zero) {
		if criteria == '1' { // o2
			out = one
		} else { // co2
			out = zero
		}
	} else {
		if criteria == '1' { // o2
			out = one
		} else { // co2
			out = zero
		}
	}

	if len(out) == 1 {
		return out
	}

	return filter(out, criteria, currentKey+1)
}

func Part2(lines []string) (o2 uint64, co2 uint64) {
	o2, err := strconv.ParseUint(
		filter(lines, '1', 0)[0],
		2,
		64,
	)
	if err != nil {
		panic(err)
	}

	co2, err = strconv.ParseUint(
		filter(lines, '0', 0)[0],
		2,
		64,
	)
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// to string
	lines := strings.Split(string(data), "\n")

	gamma, epsilon := Part1(lines)
	o2, co2 := Part2(lines)

	fmt.Printf("Solution #1: %d\n", gamma*epsilon)
	fmt.Printf("Solution #2: %d\n", o2*co2)
}
