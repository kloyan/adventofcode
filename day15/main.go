package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const targetRow int = 2000000

func main() {
	covered := make(map[int]bool)
	for _, str := range strings.Split(strings.TrimSpace(input), "\n") {
		var sx, sy, bx, by int
		fmt.Sscanf(str, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)

		dist := abs(sx-bx) + abs(sy-by) - abs(targetRow-sy)
		if dist < 0 {
			continue
		}

		for x := sx - dist; x <= sx+dist; x++ {
			if x != bx || targetRow != by {
				covered[x] = true
			}
		}
	}

	fmt.Println("Answer for Part 1:", len(covered))
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
