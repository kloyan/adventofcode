package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/kloyan/adventofcode/common"
)

//go:embed input.txt
var input string

func getTopCalories() int {
	s := bufio.NewScanner(strings.NewReader(input))
	largest := 0
	current := 0

	for s.Scan() {
		t := s.Text()
		if t != "" {
			current += common.ToInt(t)
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
	s := bufio.NewScanner(strings.NewReader(input))
	topk := make([]int, k)
	current := 0

	for s.Scan() {
		t := s.Text()
		if t != "" {
			current += common.ToInt(t)
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
	fmt.Println("Answer for Part 1", getTopCalories())
	fmt.Println("Answer for Part 2:", getSumTopKCalories(3))
}
