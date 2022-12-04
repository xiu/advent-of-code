package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func overlaps(pair string, target string) bool {
	floorPair, _ := strconv.Atoi(strings.Split(pair, "-")[0])
	ceilingPair, _ := strconv.Atoi(strings.Split(pair, "-")[1])

	floorTarget, _ := strconv.Atoi(strings.Split(target, "-")[0])
	ceilingTarget, _ := strconv.Atoi(strings.Split(target, "-")[1])

	if (floorPair >= floorTarget && floorPair <= ceilingTarget) || (ceilingPair >= floorTarget && ceilingPair <= ceilingTarget) {
		return true
	}

	return false
}

func isWithin(pair string, target string) bool {
	floorPair, _ := strconv.Atoi(strings.Split(pair, "-")[0])
	ceilingPair, _ := strconv.Atoi(strings.Split(pair, "-")[1])

	floorTarget, _ := strconv.Atoi(strings.Split(target, "-")[0])
	ceilingTarget, _ := strconv.Atoi(strings.Split(target, "-")[1])

	if floorPair >= floorTarget && ceilingPair <= ceilingTarget {
		return true
	}

	return false
}

func Parse(lines []string) (int, int) {
	fullyContained := 0
	rangeOverlap := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		left := strings.Split(line, ",")[0]
		right := strings.Split(line, ",")[1]

		if isWithin(left, right) || isWithin(right, left) {
			fullyContained += 1
		}

		if overlaps(left, right) || overlaps(right, left) {
			rangeOverlap += 1
		}
	}

	return fullyContained, rangeOverlap
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	count, countOverlap := Parse(lines)

	fmt.Printf("Solution #1: %d\n", count)
	fmt.Printf("Solution #2: %d\n", countOverlap)
}
