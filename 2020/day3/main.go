package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getTreeCount(matrix [][]string, stepDown int, stepRight int) (trees int) {
	var x int = 0
	var y int = 0
	trees = 0

	for {
		if y > len(matrix)-1 {
			break
		}

		if x > len(matrix[y])-1 {
			x = x - len(matrix[y])
		}

		if matrix[y][x] == "#" {
			trees++
		}

		x = x + stepRight
		y = y + stepDown
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

	var matrix [][]string
	var lineId int = 0

	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
		lineId++
	}

	fmt.Printf("Solution #1: %d\n", getTreeCount(matrix, 1, 3))

	fmt.Printf("Solution #2: %d\n",
		getTreeCount(matrix, 1, 1)*
			getTreeCount(matrix, 1, 3)*
			getTreeCount(matrix, 1, 5)*
			getTreeCount(matrix, 1, 7)*
			getTreeCount(matrix, 2, 1))
}
