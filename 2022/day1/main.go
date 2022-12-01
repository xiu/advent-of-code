package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sort"
	"strings"
)

func Parse(lines []string) (calories []int) {
	current := 0
	for _, value := range lines {
		if value == "" {
			calories = append(calories, current)
			current = 0
			continue
		}
		val, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		current += val
	}

	sort.Ints(calories)
	return calories
}


func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	calories := Parse(lines)

	fmt.Printf("Solution #1: %d\n", calories[len(calories)-1])
	fmt.Printf("Solution #2: %d\n", calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3])
}
