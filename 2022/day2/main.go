package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Parse(lines []string) (score1 int, score2 int) {
	for _, value := range lines {
		if value == "" { continue }
		valueSplit := strings.Split(value, " ")
		opp := valueSplit[0]
		self := valueSplit[1]
		
		switch (opp) {
		case "A": // rock
			if self == "X" { // rock / lose
				score1 += 4
				score2 += 3
			}
			if self == "Y" { // paper / draw
				score1 += 8
				score2 += 4
			}
			if self == "Z" { // scissor / win
				score1 += 3
				score2 += 8
			}
		case "B": // paper
			if self == "X" { // rock / lose
				score1 += 1
				score2 += 1
			}
			if self == "Y" { // paper / draw
				score1 += 5
				score2 += 5
			}
			if self == "Z" { // scissor / win
				score1 += 9
				score2 += 9
			}
		case "C": // scissor
			if self == "X" { // rock / lose
				score1 += 7
				score2 += 2
			}
			if self == "Y" { // paper / draw
				score1 += 2
				score2 += 6
			}
			if self == "Z" { // scissor / win
				score1 += 6
				score2 += 7
			}
		}
	}

	return score1, score2
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	score, score2 := Parse(lines)

	fmt.Printf("Solution #1: %d\n", score)
	fmt.Printf("Solution #2: %d\n", score2)
}
