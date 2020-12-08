package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func executeCommands(input []string, flipped map[int]bool) (pos int) {
	pos = 0
	acc := 0
	visited := map[int]bool{}
	currentFlip := len(flipped)

	for {
		command, value := parseCommand(input[pos])

		_, ok := visited[pos]
		if ok {
			if flipped == nil {
				// solution #1
				return acc
			}
			return executeCommands(input, flipped)
		}
		visited[pos] = true

		switch command {
		case "nop":
			if flipped != nil {
				_, ok := flipped[pos]
				if !ok && len(flipped) == currentFlip {
					// we didn't try to flip that one and
					// we didn't flip in this iteration
					flipped[pos] = true
					pos = pos + value
					break
				}
			}
			pos++
		case "acc":
			acc = acc + value
			pos++
		case "jmp":
			if flipped != nil {
				_, ok := flipped[pos]
				if !ok && len(flipped) == currentFlip {
					// we didn't try to flip that one and
					// we didn't flip in this iteration
					flipped[pos] = true
					pos++
					break
				}
			}
			pos = pos + value
		}

		if pos == len(input) {
			// last instruction of the file, we terminate
			return acc
		}
	}
}

func parseCommand(input string) (command string, value int) {
	command = strings.Split(input, " ")[0]
	value, err := strconv.Atoi(strings.Split(input, " ")[1])
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

	fmt.Printf("Solution #1: %d\n", executeCommands(lines, nil))
	fmt.Printf("Solution #2: %d\n", executeCommands(lines, map[int]bool{}))

}
