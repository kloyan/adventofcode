package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type point struct {
	x, y int
}

func main() {
	grid := map[point]rune{}
	var start, end point

	for x, r := range strings.Fields(input) {
		for y, c := range r {
			p := point{x, y}
			grid[p] = c

			if c == 'S' {
				start = p
				grid[start] = 'a'
			} else if c == 'E' {
				end = p
				grid[end] = 'z'
			}
		}
	}

	queue := []point{end}
	dist := map[point]int{}
	var newStart *point

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if grid[curr] == 'a' && newStart == nil {
			newStart = &curr
		}

		for _, d := range []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next := point{curr.x + d.x, curr.y + d.y}

			// point is outside the grid
			if _, ok := grid[next]; !ok {
				continue
			}

			// point was already visited
			if _, ok := dist[next]; ok {
				continue
			}

			if grid[curr] <= 1+grid[next] {
				dist[next] = dist[curr] + 1
				queue = append(queue, next)
			}
		}
	}

	fmt.Println("Answer for Part 1:", dist[start])
	fmt.Println("Answer for Part 2:", dist[*newStart])
}
