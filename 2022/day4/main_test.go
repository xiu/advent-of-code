package main

import (
	"testing"
)

var example = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		var want = 2

		got, _ := Parse(example)
		if got != want {
			t.Errorf("Part1 = %d; want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		var want = 4

		_, got := Parse(example)
		if got != want {
			t.Errorf("Part2 = %d; want %d", got, want)
		}
	})
}
