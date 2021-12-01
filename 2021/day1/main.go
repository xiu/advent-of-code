package main

import (
	"fmt"
	"io/ioutil"
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

	// solution 1
	var values []int
	var prev, inc int
	prev = 0
	inc = 0

	for _, value := range lines {
		if value == "" {
			continue
		}

		valint, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		if prev != 0 && valint > prev {
			inc++
		}
		prev = valint

		values = append(values, valint)
	}

	// solution 2
	var cur2, prev2, inc2 int
	prev2 = 0
	cur2 = 0
	inc2 = 0

	for key, _ := range values {
		if key < 2 {
			continue
		}
		prev2 = cur2

		cur2 = values[key] + values[key-1] + values[key-2]

		if prev2 == 0 {
			continue
		}

		if cur2 > prev2 {
			inc2++
		}
	}

	fmt.Printf("Solution #1: %d\n", inc)
	fmt.Printf("Solution #2: %d\n", inc2)
}
