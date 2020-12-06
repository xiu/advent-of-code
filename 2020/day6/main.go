package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func removeDuplicates(input string) (output string) {
	known := map[string]bool{}

	for _, r := range input {
		_, ok := known[string(r)]
		if !ok {
			known[string(r)] = true
			output = output + string(r)
		}
	}
	return
}

func countAnswers(input []string) (count int) {
	count = len(removeDuplicates(strings.Join(input, "")))
	return
}

func countAnswers2(input []string) (count int) {
	peopleCount := len(input)

	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if strings.Count(strings.Join(input, ""), string(r)) == peopleCount {
			count++
		}
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

	var currentGroup []string
	var answersCount int
	var answersCount2 int

	for _, line := range lines {
		if line == "" {
			answersCount += countAnswers(currentGroup)
			answersCount2 += countAnswers2(currentGroup)
			currentGroup = []string{}
			continue
		}
		currentGroup = append(currentGroup, line)
	}

	fmt.Printf("Solution #1: %d\n", answersCount)
	fmt.Printf("Solution #2: %d\n", answersCount2)
}
