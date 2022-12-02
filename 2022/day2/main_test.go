package main

import (
	"testing"
)

var example = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		var want = 15

		got, _ := Parse(example)
		if got != want {
			t.Errorf("Part1 = %d; want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		var want = 12

		_, got := Parse(example)
		if got != want {
			t.Errorf("Part2 = %d; want %d", got, want)
		}
	})
}