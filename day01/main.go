package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/kloyan/adventofcode/common"
)

//go:embed input.txt
var input string

func getTopCalories() int {
	largest := 0
	current := 0

	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			current += common.ToInt(line)
			continue
		}

		if current > largest {
			largest = current
		}

		current = 0
	}

	return largest
}

func getSumTopKCalories(k int) int {
	topk := make([]int, k)
	current := 0

	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			current += common.ToInt(line)
			continue
		}

		if current > topk[0] {
			topk[0] = current
			sort.Ints(topk)
		}

		current = 0
	}

	sum := 0
	for _, calories := range topk {
		sum += calories
	}

	return sum
}

func main() {
	fmt.Println("Answer for Part 1:", getTopCalories())
	fmt.Println("Answer for Part 2:", getSumTopKCalories(3))
}
