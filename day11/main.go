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

const monkeyTemplate = `Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`

type Monkey struct {
	Items []int
	Op    func(int) int
	Test  func(int) int
}

func main() {
	monkeys, lcm := parseMonkeys(input)

	fmt.Println("Answer for Part 1:", solve(monkeys, 20, func(w int) int { return (w % lcm) / 3 }))
	fmt.Println("Answer for Part 2:", solve(monkeys, 10_000, func(w int) int { return w % lcm }))
}

func solve(monkeys []Monkey, rounds int, decreaseWorry func(int) int) int {
	cpy := append([]Monkey{}, monkeys...)
	inspected := make([]int, len(cpy))

	for i := 0; i < rounds; i++ {
		for j, monkey := range cpy {
			for _, w := range monkey.Items {
				worry := decreaseWorry(monkey.Op(w))
				target := monkey.Test(worry)
				cpy[target].Items = append(cpy[target].Items, worry)
				inspected[j]++
			}

			cpy[j].Items = nil
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	return inspected[0] * inspected[1]
}

func parseMonkeys(input string) ([]Monkey, int) {
	monkeys := make([]Monkey, 0)
	lcm := 1

	for _, str := range strings.Split(strings.TrimSpace(input), "\n\n") {
		var i, test, operand, trueTarget, falseTarget int
		var items, symbol string

		replacer := strings.NewReplacer(", ", ",", "* old", "^ 2")
		fmt.Sscanf(replacer.Replace(str), monkeyTemplate,
			&i, &items, &symbol, &operand, &test, &trueTarget, &falseTarget)

		monkey := Monkey{
			Items: parseItems(items),
			Op: func(worry int) int {
				switch symbol {
				case "+":
					return worry + operand
				case "*":
					return worry * operand
				default:
					return worry * worry
				}
			},
			Test: func(worry int) int {
				if worry%test == 0 {
					return trueTarget
				}

				return falseTarget
			},
		}

		monkeys = append(monkeys, monkey)
		lcm *= test
	}

	return monkeys, lcm
}

func parseItems(str string) []int {
	items := make([]int, 0)

	for _, num := range strings.Split(str, ",") {
		items = append(items, common.ToInt(num))
	}

	return items
}
