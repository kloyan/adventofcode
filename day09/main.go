package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type coord struct {
	x, y int
}

func main() {
	moves := strings.Split(strings.TrimSpace(input), "\n")

	fmt.Println("Answer for Part 1:", moveRope(2, moves))
	fmt.Println("Answer for Part 2:", moveRope(10, moves))
}

func moveRope(n int, moves []string) int {
	rope := make([]coord, n)
	visited := make(map[coord]bool)

	for _, move := range moves {
		var direction byte
		var steps int
		fmt.Sscanf(move, "%c %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			moveHead(rope, direction)
			moveTail(rope)
			visited[rope[n-1]] = true
		}
	}

	return len(visited)
}

func moveHead(rope []coord, direction byte) {
	switch direction {
	case 'L':
		rope[0].x--
	case 'R':
		rope[0].x++
	case 'U':
		rope[0].y++
	case 'D':
		rope[0].y--
	}
}

func moveTail(rope []coord) {
	for i := 1; i < len(rope); i++ {
		dx := rope[i-1].x - rope[i].x
		dy := rope[i-1].y - rope[i].y

		if adjacent(dx, dy) {
			return
		}

		if dx > 0 {
			rope[i].x++
		} else if dx < 0 {
			rope[i].x--
		}

		if dy > 0 {
			rope[i].y++
		} else if dy < 0 {
			rope[i].y--
		}
	}
}

func adjacent(dx, dy int) bool {
	return math.Abs(float64(dx)) <= 1 && math.Abs(float64(dy)) <= 1
}
