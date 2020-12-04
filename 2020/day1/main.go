package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// to string
	lines := strings.Split(string(data), "\n")

	// to int
	var values []int

	for _, value := range lines {
		valint, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		values = append(values, valint)
	}

	// guess first star
	for _, value := range values {
		for _, value2 := range values {
			if value+value2 == 2020 {
				fmt.Printf("Solution first star: %d\n", value*value2)
			}
		}
	}

	// second star
	for _, value := range values {
		for _, value2 := range values {
			for _, value3 := range values {
				if value+value2+value3 == 2020 {
					fmt.Printf("Solution second star: %d\n", value*value2*value3)
				}
			}
		}
	}
}
