package main

import (
	"fmt"
	"os"
	"strings"
)

func isVisible(lines []string, x int, y int) (bool) {
	value := lines[y][x]
	total := 4
	
	// north-south
	for i := 0; i < y; i++ {
		if lines[i][x] >= value {
			total -= 1
			break
		}
	}

	for i := y+1; i < len(lines[0]); i++ {
		if lines[i][x] >= value {
			total -= 1
			break
		}
	}

	// west-east
	for i := 0; i < x; i++ {
		if lines[y][i] >= value {
			total -= 1
			break
		}
	}

	for i := x+1; i < len(lines[0]); i++ {
		if lines[y][i] >= value {
			total -= 1
			break
		}
	}
	
	return total > 0
}

func countTrees(lines []string, x int, y int) (score int) {
	value := lines[y][x]
	score = 1
	count := 0
	
	// north-south
	for i := y-1; i >= 0; i-- {
		count += 1
		if lines[i][x] >= value {
			break
		}
	}
	score *= count
	count = 0

	for i := y+1; i < len(lines[0]); i++ {
		count += 1
		if lines[i][x] >= value {
			break
		}
	}
	score *= count
	count = 0

	// west-east
	for i := x-1; i >= 0; i-- {
		count += 1
		if lines[y][i] >= value {
			break
		}
	}
	score *= count
	count = 0

	for i := x+1; i < len(lines[0]); i++ {
		count += 1
		if lines[y][i] >= value {
			break
		}
	}
	score *= count
	
	return score
}

func Parse(lines []string) (part1 int, part2 int) {
	maxX := len(lines[0])
	maxY := len(lines)

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if isVisible(lines, x, y) {
				part1 += 1
			}

			if c := countTrees(lines, x, y); c > part2 {
				part2 = c
			}
		}
	}

	return part1, part2
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	part1, part2 := Parse(lines)

	fmt.Printf("Solution #1: %d\n", part1)
	fmt.Printf("Solution #2: %d\n", part2)
}
