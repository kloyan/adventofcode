package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const imageWidth = 40

var opCycles = map[string]int{
	"noop": 1,
	"addx": 2,
}

func main() {
	cmds := strings.Split(strings.TrimSpace(input), "\n")

	fmt.Println("Answer for Part 1:", calcSignalStrength(cmds))
	fmt.Print("Answer for Part 2:\n", renderImage(cmds))
}

func calcSignalStrength(cmds []string) int {
	var str, cycle int
	var reg int = 1

	for _, cmd := range cmds {
		var op string
		var arg int
		fmt.Sscanf(cmd, "%s %d", &op, &arg)

		for i := 0; i < opCycles[op]; i++ {
			cycle++
			if cycle == 20 || (cycle-20)%40 == 0 {
				str += cycle * reg
			}
		}

		reg += arg
	}

	return str
}

func renderImage(cmds []string) string {
	image := new(strings.Builder)
	row, col, x := 0, 0, 1

	for _, cmd := range cmds {
		var op string
		var arg int
		fmt.Sscanf(cmd, "%s %d", &op, &arg)

		for i := 0; i < opCycles[op]; i++ {
			if col >= x-1 && col <= x+1 {
				image.WriteRune('█')
			} else {
				image.WriteByte(' ')
			}

			col++
			if col == imageWidth {
				col = 0
				row++
				image.WriteByte('\n')
			}
		}

		x += arg
	}

	return image.String()
}
