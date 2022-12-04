package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func countFullyContainedPairs() int {
	count := 0
	for _, line := range strings.Fields(input) {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &x1, &y1, &x2, &y2)

		if (x1 >= x2 && y1 <= y2) || (x1 <= x2 && y1 >= y2) {
			count++
		}
	}

	return count
}

func countOverlappingPairs() int {
	count := 0
	for _, line := range strings.Fields(input) {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &x1, &y1, &x2, &y2)

		if x1 <= y2 && y1 >= x2 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println("Answer to Part 1:", countFullyContainedPairs())
	fmt.Println("Answer to Part 2:", countOverlappingPairs())
}
