package main

import (
	"testing"
)

var example = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		var want = 21

		got, _ := Parse(example)
		if got != want {
			t.Errorf("Part1 = %d; want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		var want = 8

		_, got := Parse(example)
		if got != want {
			t.Errorf("Part2 = %d; want %d", got, want)
		}
	})
}
