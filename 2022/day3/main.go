package main

import (
	"fmt"
	"os"
	"strings"
)

func getPriority(r rune) int {
	if int(r) <= 90 { // uppercase
		return int(r) - 38
	} else { // lowercase
		return int(r) - 96
	}
}

func removeDuplicates(line string) (clean string) {
	hit := make(map[rune]bool)

	for _, r := range line {
		if _, value := hit[r]; !value {
			hit[r] = true
			clean = fmt.Sprintf("%s%s", clean, string(r))
		}
	}

	return clean
}

func findBadge(lines []string, i int) string {
	group := fmt.Sprintf("%s%s%s",
		removeDuplicates(lines[i]),
		removeDuplicates(lines[i-1]),
		removeDuplicates(lines[i-2]),
	)

	for _, r := range group {
		if strings.Count(group, string(r)) == 3 {
			return string(r)
		}
	}
	return ""
}

func Parse(lines []string) (int, int) {
	uniqueValues := ""
	badges := ""
	sum := 0
	sumBadges := 0

	for i, value := range lines {
		if value == "" {
			continue
		}
		if (i+1)%3 == 0 {
			badges = fmt.Sprintf("%s%s", badges, findBadge(lines, i))
		}
		left := value[:len(value)/2]
		right := value[len(value)/2:]

		localUniqueValues := ""

		for _, r := range left {
			if strings.Contains(right, string(r)) {
				if !strings.Contains(localUniqueValues, string(r)) {
					localUniqueValues = fmt.Sprintf("%s%s", localUniqueValues, string(r))
				}
			}
		}

		uniqueValues = fmt.Sprintf("%s%s", uniqueValues, localUniqueValues)
	}

	for _, r := range uniqueValues {
		sum += getPriority(r)
	}

	for _, r := range badges {
		sumBadges += getPriority(r)
	}

	return sum, sumBadges
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	sum, sumBadges := Parse(lines)

	fmt.Printf("Solution #1: %d\n", sum)
	fmt.Printf("Solution #2: %d\n", sumBadges)
}
