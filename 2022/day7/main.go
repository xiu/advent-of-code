package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name string
	Directories []*Directory
	Files []*File
	Parent *Directory
}

func (d *Directory) Out() *Directory {
	return d.Parent
}

func (d *Directory) Size() (total int) {
	for _, dir := range d.Directories {
		total += dir.Size()
	}

	for _, file := range d.Files {
		total += file.Size
	}

	return total
}

func (d *Directory) FindDirectory(name string) (dir *Directory) {
	for _, dir := range d.Directories {
		if dir.Name == name {
			return dir
		}
	}
	return nil
}

func (d *Directory) SizeSum(le int) (total int) {
	for _, dir := range d.Directories {
		if dir.Size() <= le {
			total += dir.Size()
		}
		total += dir.SizeSum(le)
	}
	return
}

func (d *Directory) FreeSpace(ge int) (res int) {
	for _, dir := range d.Directories {
		size := dir.Size()
		subSize := dir.FreeSpace(ge)
		if size >= ge && (size < res || res == 0) {
			res = size
		}
		if subSize != 0 && subSize < res {
			res = subSize
		}
	}
	return res
}

func Parse(lines []string) (part1 int, part2 int) {
	var root = &Directory{
		Name: "/",
	}
	var current = root

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line[0] == byte('$') {
			// Just filter cd, we don't mind ls
			if line[2:4] == "cd" {
				dir := line[5:]
				if dir == "/" {
					continue
				}

				if dir == ".." {
					current = current.Parent
					continue
				}

				if found := current.FindDirectory(dir); found == nil {
					newDirectory := &Directory{
						Name: dir,
						Parent: current,
					}

					current.Directories = append(current.Directories, newDirectory)

					current = newDirectory
				} else {
					current = found
				}
			}
		} else {
			split := strings.Split(line, " ")

			if val, err := strconv.Atoi(split[0]); err != nil {
				// we are on a directory
				current.Directories = append(current.Directories, &Directory{
					Name: split[1],
					Parent: current,
				})
			} else {
				// file
				current.Files = append(current.Files, &File{
					Name: split[1],
					Size: val,
				})
			}
		}
	}

	part1 = root.SizeSum(100000)

	// part2
	objective := 30000000 - (70000000 - root.Size())
	part2 = root.FreeSpace(objective)

	return part1, part2
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	part1, part2 := Parse(lines)

	fmt.Printf("Solution #1: %d\n", part1)
	fmt.Printf("Solution #2: %d\n", part2)
}
