package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type BagChildren struct {
	Quantity int
	Bag      *Bag
}

type Bag struct {
	Name     string
	Children []BagChildren
	Parents  []*Bag
}

var bagList []*Bag

func findBag(name string) *Bag {
	for _, bag := range bagList {
		if bag.Name == name {
			return bag
		}
	}

	// if we don't find one, create it
	bag := &Bag{
		Name: name,
	}
	bagList = append(bagList, bag)
	return bag
}

func countParents(bag *Bag, totalParents map[string]bool) (count int) {
	if totalParents == nil {
		totalParents = map[string]bool{}
	}

	for _, parent := range bag.Parents {
		_, ok := totalParents[parent.Name]
		if !ok {
			totalParents[parent.Name] = true
			count++
			count = count + countParents(parent, totalParents)
		}
	}

	return count
}

func countChildren(bag *Bag, depth int) (count int) {
	if depth > 0 {
		// adding itself
		count++
	}

	if len(bag.Children) == 0 {
		// just return itself
		return
	}

	for _, child := range bag.Children {
		cc := countChildren(child.Bag, depth+1)

		if cc > 0 {
			count = count + child.Quantity*cc
		} else {
			count = count + child.Quantity
		}
	}

	return count
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// to string
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		r, err := regexp.Compile(`([0-9]*)+ ([a-z]* [a-z]*) bag`)
		if err != nil {
			panic(err)
		}

		bags := r.FindAllStringSubmatch(line, -1)

		if bags[0][2] == "no other" {
			continue
		}

		currentBagName := strings.Join(
			strings.Split(line, " ")[0:2],
			" ",
		)

		currentBag := findBag(currentBagName)

		for _, bag := range bags {
			currentQuantity, err := strconv.Atoi(bag[1])
			if err != nil {
				panic(err)
			}
			currentChild := findBag(bag[2])

			currentBag.Children = append(currentBag.Children, BagChildren{currentQuantity, currentChild})
			currentChild.Parents = append(currentChild.Parents, currentBag)
		}
	}

	fmt.Printf("Solution #1: %d\n", countParents(findBag("shiny gold"), nil))
	fmt.Printf("Solution #2: %d\n", countChildren(findBag("shiny gold"), 0))

}
