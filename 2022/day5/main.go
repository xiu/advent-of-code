package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getResult(stacks [][]string) (res string) {
	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}
		res = fmt.Sprintf("%s%s", res, string(stack[0]))
	}
	return res
}

func Parse(lines []string) (part1 string, part2 string) {
	stacks := make([][]string, 9)
	stacks2 := make([][]string, 9)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if !strings.Contains(line, "move") {
			if !strings.Contains(line, "[") {
				continue
			}

			// each item is separated by 4
			j := 0
			for i, item := range line {
				if i%4 == 1 {
					if item != ' ' {
						stacks[j] = append(stacks[j], string(item))
						copy(stacks2, stacks)
					}
					j++
				}
			}
		} else {
			qty, _ := strconv.Atoi(strings.Split(line, " ")[1])
			src, _ := strconv.Atoi(strings.Split(line, " ")[3])
			dst, _ := strconv.Atoi(strings.Split(line, " ")[5])

			// part1
			for i := 0; i < qty; i++ {
				stacks[dst-1] = append([]string{stacks[src-1][0]}, stacks[dst-1]...)
				stacks[src-1] = stacks[src-1][1:]
			}

			// part2
			// we need a copy as Go doesn't forcefully realloc on append
			// see: https://stackoverflow.com/questions/33783281/golang-appending-slices-with-or-w-o-allocation
			orig := make([]string, len(stacks2[dst-1]))
			copy(orig, stacks2[dst-1])
			stacks2[dst-1] = append([]string{}, stacks2[src-1][0:qty]...)
			stacks2[dst-1] = append(stacks2[dst-1], orig...)
			stacks2[src-1] = stacks2[src-1][qty:]
		}
	}

	return getResult(stacks), getResult(stacks2)
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	count, countOverlap := Parse(lines)

	fmt.Printf("Solution #1: %s\n", count)
	fmt.Printf("Solution #2: %s\n", countOverlap)
}
