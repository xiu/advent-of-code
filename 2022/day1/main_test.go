package main

import (
	"testing"
)

var example = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
	"",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		var want = 24000

		calories := Parse(example)
		got := calories[len(calories)-1]
		if got != want {
			t.Errorf("Part1 = %d; want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		var want = 45000

		calories := Parse(example)
		got := calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]
		if got != want {
			t.Errorf("Part2 = %d; want %d", got, want)
		}
	})
}