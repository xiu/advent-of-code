package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const preambleSize = 25

var preamble *IntQueue

// IntQueue implements a queue of Integer
type IntQueue struct {
	stack []int
}

// New creates a new Queue
func (s *IntQueue) New() *IntQueue {
	s.stack = []int{}
	return s
}

// Dequeue returns the first element in the queue
func (s *IntQueue) Dequeue() int {
	elem := s.stack[len(s.stack)-1]
	s.stack = s.stack[1:len(s.stack)]
	return elem
}

// Enqueue adds an element to the back of the queue
func (s *IntQueue) Enqueue(value int) {
	s.stack = append(s.stack, value)
}

// Len returns the length of the queue
func (s *IntQueue) Len() int {
	return len(s.stack)
}

func match(value int) bool {
	for outterk, item := range preamble.stack {
		for innerk, item2 := range preamble.stack {
			if innerk == outterk {
				continue
			}

			if item+item2 == value {
				return true
			}
		}
	}

	return false
}

func part2(input []string, target int) (solution int) {
	for key := range input[2:] {
		items := []int{}
		total := 0

		for i := key; i >= 0; i-- {
			value2, err := strconv.Atoi(input[i])
			if err != nil {
				panic(err)
			}

			items = append(items, value2)
			total += value2

			if total == target {
				sort.Ints(items)
				solution = items[0] + items[len(items)-1]
				return
			}
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

	preamble = new(IntQueue)
	preamble.stack = []int{}

	for _, line := range lines[0:preambleSize] {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		preamble.Enqueue(value)
	}

	for _, line := range lines[preambleSize:] {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if !match(value) {
			fmt.Printf("Solution #1: %d\n", value)
			fmt.Printf("Solution #2: %d\n", part2(lines, value))
			return
		}

		preamble.Dequeue()
		preamble.Enqueue(value)
	}
}
