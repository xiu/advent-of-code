package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func turn(azimuth string, direction string, degrees int) (string) {
	east := []string{"S", "W", "N", "E", "S", "W", "N"}
	south := []string{"W", "N", "E", "S", "W", "N", "E"}
	west := []string{"N", "E", "S", "W", "N", "E", "S"}
	north := []string{"E", "S", "W", "N", "E", "S", "W"}
	
	offset := degrees / 90

	if direction == "L" {
		offset = -offset
	}

	switch azimuth {
	case "N":
		return north[3+offset]
	case "S":
		return south[3+offset]
	case "E":
		return east[3+offset]
	case "W":
		return west[3+offset]
	}

	return ""
}

func turnWaypoint(x int, y int, direction string, degrees int) (int, int) {
	var optX, optY int
	
	if direction == "L" {
		degrees = -degrees
		optX = -1
		optY = 1
	} else {
		// R
		optX = 1
		optY = -1
	}

	offset := degrees / 90

	switch offset {
	case 1, -1:
		bufX := x
		x = y*optX
		y = bufX*optY
	case 2, -2:
		x = x*-1
		y = y*-1
	case 3, -3:
		if direction == "L" {
			direction = "R"
		} else {
			direction = "L"
		}
		x, y = turnWaypoint(x, y, direction, 90)
	}

	return x,y
}

// ProcessInput takes an input file and executes the challenges
func ProcessInput(input string) {
	fmt.Printf("---- Processing %s\n", input)
	data, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("Error", err)
	}

	// to string
	lines := strings.Split(string(data), "\n")

	// we start facing East
	azimuth := "E" 

	// +x towards north
	// -x towards south
	// +y towards east
	// -y towards west
	x := 0
	y := 0

	for _, line := range lines {
		action := string(line[0])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if action == "F" {
			action = azimuth
		}

		switch action {
		case "L":
			azimuth = turn(azimuth, action, value)
		case "R":
			azimuth = turn(azimuth, action, value)
		case "N":
			y += value
		case "S":
			y -= value
		case "E":
			x += value
		case "W":
			x -= value
		}

	}

	fmt.Printf("Solution #1: %d\n", abs(y) + abs(x))

	// reinit for part 2
	x = 0
	y = 0
	waypointX := 10
	waypointY := 1
	for _, line := range lines {
		action := string(line[0])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if action == "F" {
			x = x + waypointX*value
			y = y + waypointY*value
		}

		switch action {
		case "L":
			waypointX, waypointY = turnWaypoint(waypointX, waypointY, action, value)
		case "R":
			waypointX, waypointY = turnWaypoint(waypointX, waypointY, action, value)
		case "N":
			waypointY += value
		case "S":
			waypointY -= value
		case "E":
			waypointX += value
		case "W":
			waypointX -= value
		}
	}

	fmt.Printf("Solution #2: %d\n", abs(y) + abs(x))

	fmt.Printf("---- Processing %s done\n", input)

}

func main() {
	ProcessInput("example")
	ProcessInput("input")
}
