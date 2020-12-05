package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func getSeat(input string) (row int, column int, seat int) {
	upperRow := 127
	lowerRow := 0

	upperColumn := 7
	lowerColumn := 0

	for _, c := range input {
		midRow := int((upperRow - lowerRow) / 2)
		midColumn := int((upperColumn - lowerColumn) / 2)
		switch c {
		case 'F':
			if upperRow-lowerRow == 1 {
				row = lowerRow
			} else {
				upperRow = lowerRow + midRow
			}
		case 'B':
			if upperRow-lowerRow == 1 {
				row = upperRow
			} else {
				lowerRow = lowerRow + midRow + 1
			}
		case 'L':
			if upperColumn-lowerColumn == 1 {
				column = lowerColumn
			} else {
				upperColumn = lowerColumn + midColumn
			}
		case 'R':
			if upperColumn-lowerColumn == 1 {
				column = upperColumn
			} else {
				lowerColumn = lowerColumn + midColumn + 1
			}
		}
	}

	seat = row*8 + column

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

	var seats []int

	for _, line := range lines {
		_, _, seat := getSeat(line)

		seats = append(seats, seat)
	}

	sort.Ints(seats)

	fmt.Printf("Solution #1: %d\n", seats[len(seats)-1])

	for key, seat := range seats {
		if key == 0 || key == len(seats)-1 {
			continue
		}

		if seats[key+1] == seat+2 {
			fmt.Printf("Solution #2: %d\n", seat+1)
		}
	}
}
