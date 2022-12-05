package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/kloyan/adventofcode/common"
)

//go:embed input.txt
var input string

func main() {
	parts := strings.Split(input, "\n\n")
	crates := strings.Split(parts[0], "\n")
	cmd := strings.Split(strings.TrimSpace(parts[1]), "\n")

	stacks := populateStacks(crates)
	part1(stacks, cmd)
	fmt.Println("Answer for Part 1:", getTopCrates(stacks))

	stacks = populateStacks(crates)
	part2(stacks, cmd)
	fmt.Println("Answer for Part 2:", getTopCrates(stacks))
}

func populateStacks(crates []string) []*common.Stack {
	stacks := make([]*common.Stack, 0)

	for j := 1; j < len(crates[0]); j += 4 {
		s := common.Stack{}
		for i := len(crates) - 2; i >= 0; i-- {
			c := crates[i][j]
			// ignore whitespaces
			if c != 32 {
				s.Push(c)
			}
		}

		stacks = append(stacks, &s)
	}

	return stacks
}

func part1(stacks []*common.Stack, cmd []string) {
	for _, cmd := range cmd {
		var num, from, to int
		fmt.Sscanf(cmd, "move %d from %d to %d", &num, &from, &to)

		for i := 0; i < num; i++ {
			el, err := stacks[from-1].Pop()
			if err != nil {
				panic(err)
			}

			stacks[to-1].Push(el)
		}
	}
}

func part2(stacks []*common.Stack, cmd []string) {
	tmp := common.Stack{}
	for _, cmd := range cmd {
		var num, from, to int
		fmt.Sscanf(cmd, "move %d from %d to %d", &num, &from, &to)

		for i := 0; i < num; i++ {
			el, err := stacks[from-1].Pop()
			if err != nil {
				panic(err)
			}

			tmp.Push(el)
		}

		for !tmp.IsEmpty() {
			el, err := tmp.Pop()
			if err != nil {
				panic(err)
			}

			stacks[to-1].Push(el)
		}
	}
}

func getTopCrates(stacks []*common.Stack) string {
	sb := new(strings.Builder)
	for _, s := range stacks {
		el, err := s.Pop()
		if err != nil {
			panic(err)
		}

		b := el.(byte)
		sb.WriteByte(b)
	}

	return sb.String()
}
