package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func calcScore(item rune) int {
	if item > 96 {
		return int(item) - 96
	}

	return int(item) - 38
}

func getErrorSum() int {
	sum := 0

	for _, items := range strings.Fields(input) {
		m := make(map[rune]bool)

		for _, item := range items[:len(items)/2] {
			m[item] = true
		}

		for _, item := range items[len(items)/2:] {
			if m[item] {
				sum += calcScore(item)
				break
			}
		}
	}

	return sum
}

func getBadgeSum() int {
	bags := strings.Fields(input)
	sum := 0

	for i := 0; i < len(bags)-1; i = i + 3 {
		m := make(map[rune]int)

		for j := 0; j < 3; j++ {
			items := bags[i+j]

			for _, item := range items {
				if m[item] == j {
					m[item]++
				}

				if j == 2 && m[item] == 3 {
					sum += calcScore(item)
					break
				}
			}
		}
	}

	return sum
}

func main() {
	fmt.Println("Answer for Part 1:", getErrorSum())
	fmt.Println("Answer for Part 2:", getBadgeSum())
}
