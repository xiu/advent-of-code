package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func executeRound(matrix [][]string, occupiedSeatLimit int, directCheck bool) (output [][]string, changes int, occupiedSeats int) {
	// for each coordinates, check the 8 adjacent coordinates
	// if a seat is empty and there are no occupied adjacent seats to it, switch to occupied
	// if a seat is occupied and there are 4 or more seats occupied adjacent to it, switch to empty
	// this is executed _simultaneously_ on every seat

	var (
		toSwitch [][]int
	)

	// for each row, stop at each element and check around
	for x, row := range matrix {
		for y, value := range row {
			if value == "." {
				continue
			}

			if value == "#" {
				occupiedSeats++
			}

			occupiedAdjacentCount := 0

			// check adjacent coordinates
			for checkX := x - 1; checkX <= (x + 1); checkX++ {
				if checkX < 0 || checkX > len(matrix)-1 {
					continue
				}
				for checkY := y - 1; checkY <= (y + 1); checkY++ {
					if checkY < 0 ||
						checkY > len(matrix[checkX])-1 ||
						checkY == y && checkX == x {
						continue
					}

					if directCheck {
						if matrix[checkX][checkY] == "#" {
							occupiedAdjacentCount++
						}
					} else {
						offsetX := checkX - x
						offsetY := checkY - y
						x2 := checkX
						y2 := checkY
						i := 0
						for {
							if x2 < 0 || x2 > len(matrix)-1 ||
								y2 < 0 || y2 > len(matrix[x2])-1 {
								break
							}

							if matrix[x2][y2] == "#" {
								occupiedAdjacentCount++
								break
							} else if matrix[x2][y2] == "L" {
								break
							}

							i++
							x2 = checkX + i*offsetX
							y2 = checkY + i*offsetY
						}
					}
				}
			}

			if (matrix[x][y] == "L" && occupiedAdjacentCount == 0) ||
				(matrix[x][y] == "#" && occupiedAdjacentCount >= occupiedSeatLimit) {
				toSwitch = append(toSwitch, []int{x, y})
			}
		}
	}

	// now switch
	for _, order := range toSwitch {
		x := order[0]
		y := order[1]

		switch matrix[x][y] {
		case "L":
			matrix[x][y] = "#"
			occupiedSeats++
			break
		case "#":
			matrix[x][y] = "L"
			occupiedSeats--
			break
		}
	}

	return matrix, len(toSwitch), occupiedSeats
}

// ProcessInput takes an input file and executes the challenges
func ProcessInput(input string) {
	var (
		matrix  [][]string
		matrix2 [][]string
		changes int
	)

	fmt.Printf("---- Processing %s\n", input)
	data, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("Error", err)
	}

	// to string
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
		matrix2 = append(matrix2, strings.Split(line, ""))
	}

	occupiedSeats := 0

	for {
		matrix, changes, occupiedSeats = executeRound(matrix, 4, true)

		if changes == 0 {
			break
		}
	}
	fmt.Printf("Solution #1: %d\n", occupiedSeats)

	occupiedSeats = 0
	for {
		matrix2, changes, occupiedSeats = executeRound(matrix2, 5, false)

		if changes == 0 {
			break
		}
	}
	fmt.Printf("Solution #2: %d\n", occupiedSeats)
	fmt.Printf("---- Processing %s done\n", input)

}

func main() {
	ProcessInput("example")
	ProcessInput("input")
}
