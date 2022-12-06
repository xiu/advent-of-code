package main

import (
	"fmt"
	"os"
	"strings"
)

func checkUniqueness(list []rune) bool {
	for _, r := range list {
		if strings.Count(string(list), string(r)) > 1 {
			return false
		}
	}
	return true
}

func Part1(lines []string) (int) {
	var queue []rune

	for _, line := range lines {
		if line == "" {
			continue
		}

		for i, r := range line {
			if len(queue) > 3 {
				queue = queue[1:]
			}
			queue = append(queue, r)

			if len(queue) == 4 && checkUniqueness(queue) {
				return i+1
			}
		}
	}

	return 0
}

func Part2(lines []string) (int) {
	var queue []rune

	for _, line := range lines {
		if line == "" {
			continue
		}

		for i, r := range line {
			if len(queue) > 13 {
				queue = queue[1:]
			}
			queue = append(queue, r)

			if len(queue) == 14 && checkUniqueness(queue) {
				return i+1
			}
		}
	}

	return 0
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	part1 := Part1(lines)
	part2 := Part2(lines)

	fmt.Printf("Solution #1: %d\n", part1)
	fmt.Printf("Solution #2: %d\n", part2)
}
