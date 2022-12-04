package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var mapping = map[string]string{
	"A": "X", // rock
	"B": "Y", // paper
	"C": "Z", // scissors
}

var scores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var wins = map[string]string{
	"A": "Z", // rock beats scissors
	"B": "X", // paper beats rock
	"C": "Y", // scissors beats paper
}

var loses = map[string]string{
	"A": "Y", // rock loses to paper
	"B": "Z", // paper loses to scissors
	"C": "X", // scissors loses to rock
}

func part1() int {
	total := 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var them, me string
		fmt.Sscanf(line, "%s %s", &them, &me)

		if wins[them] == me {
			total += scores[me] + 0
		} else if loses[them] == me {
			total += scores[me] + 6
		} else {
			total += scores[me] + 3
		}
	}

	return total
}

func part2() int {
	total := 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var them, outcome string
		fmt.Sscanf(line, "%s %s", &them, &outcome)

		if outcome == "X" {
			me := wins[them]
			total += scores[me] + 0
		} else if outcome == "Z" {
			me := loses[them]
			total += scores[me] + 6
		} else {
			me := mapping[them]
			total += scores[me] + 3
		}
	}

	return total
}

func main() {
	fmt.Println("Answer for Part 1:", part1())
	fmt.Println("Answer for Part 2:", part2())
}
