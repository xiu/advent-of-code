package main

import (
	"testing"
)

var example = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		var want = uint64(198)

		gamma, epsilon := Part1(example)
		got := gamma*epsilon
		if got != want {
			t.Errorf("Part1 = %d; want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		var want = uint64(230)

		o2, co2 := Part2(example)
		got := o2*co2
		if got != want {
			t.Errorf("Part2 = %d; want %d", got, want)
		}
	})
}