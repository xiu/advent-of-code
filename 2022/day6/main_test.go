package main

import (
	"testing"
)

var example = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		var want = 7

		got := Part1(example)
		if got != want {
			t.Errorf("Part1 = %d; want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		var want = 19

		got := Part2(example)
		if got != want {
			t.Errorf("Part2 = %d; want %d", got, want)
		}
	})
}
