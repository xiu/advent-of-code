package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// to string
	lines := strings.Split(string(data), "\n")

	var depth, pos int
	var aim, depth2, pos2 int

	for _, value := range lines {
		if value == "" {
			continue
		}

		words := strings.Fields(value)

		order := words[0]
		arg, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}

		switch order {
		case "forward":
			pos += arg
			pos2 += arg
			depth2 += aim * arg
		case "down":
			depth += arg
			aim += arg
		case "up":
			depth -= arg
			aim -= arg
		}
	}

	fmt.Printf("Solution #1: %d\n", depth*pos)
	fmt.Printf("Solution #2: %d\n", depth2*pos2)
}
