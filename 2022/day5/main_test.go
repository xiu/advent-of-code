package main

import (
	"testing"
)

var example = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		var want = "CMZ"

		got, _ := Parse(example)
		if got != want {
			t.Errorf("Part1 = %s; want %s", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		var want = "MCD"

		_, got := Parse(example)
		if got != want {
			t.Errorf("Part2 = %s; want %s", got, want)
		}
	})
}
