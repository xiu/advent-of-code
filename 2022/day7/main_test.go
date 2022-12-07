package main

import (
	"testing"
)

var example = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		var want = 95437

		got, _ := Parse(example)
		if got != want {
			t.Errorf("Part1 = %d; want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		var want = 24933642

		_, got := Parse(example)
		if got != want {
			t.Errorf("Part2 = %d; want %d", got, want)
		}
	})
}
