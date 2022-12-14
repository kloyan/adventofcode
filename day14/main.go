package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/kloyan/adventofcode/common"
)

//go:embed input.txt
var input string
var src = Point{500, 0}

type Point struct {
	x, y int
}

func (p Point) Add(x, y int) Point {
	return Point{p.x + x, p.y + y}
}

func main() {
	cave, height := parseCave(input)
	fmt.Println("Answer for Part 1:", reachVoid(copyCave(cave), height))
	fmt.Println("Answer for Part 2:", blockSource(cave, height+2))
}

func reachVoid(cave map[Point]bool, height int) int {
	sand := 0

	for {
		// Increment the height by one to create an artifical floor
		// Reaching it counts as reaching the void
		q := flow(cave, src, height+1)
		if q.y == height {
			return sand
		}

		sand++
		cave[q] = true
	}
}

func blockSource(cave map[Point]bool, height int) int {
	sand := 0

	for {
		q := flow(cave, src, height)
		if q == src {
			return sand + 1
		}

		sand++
		cave[q] = true
	}
}

func flow(cave map[Point]bool, p Point, height int) Point {
	q := p.Add(0, 1)
	if !cave[q] && q.y != height {
		return flow(cave, q, height)
	}

	lq := p.Add(-1, 1)
	if !cave[lq] && lq.y != height {
		return flow(cave, lq, height)
	}

	rq := p.Add(1, 1)
	if !cave[rq] && rq.y != height {
		return flow(cave, rq, height)
	}

	return p
}

func parseCave(input string) (map[Point]bool, int) {
	cave := make(map[Point]bool)
	height := 0

	for _, str := range strings.Split(strings.TrimSpace(input), "\n") {
		coords := strings.Split(str, " -> ")
		for i := 1; i < len(coords); i++ {
			from, to := parsePoint(coords[i-1]), parsePoint(coords[i])

			if from.y > height {
				height = from.y
			}

			if from.x == to.x {
				for _, y := range between(from.y, to.y) {
					cave[Point{from.x, y}] = true
				}
			} else {
				for _, x := range between(from.x, to.x) {
					cave[Point{x, from.y}] = true
				}
			}
		}
	}

	return cave, height
}

func parsePoint(str string) Point {
	from := strings.Split(str, ",")
	return Point{common.ToInt(from[0]), common.ToInt(from[1])}
}

func between(a, b int) []int {
	if a > b {
		a, b = b, a
	}

	n := []int{}
	for i := a; i <= b; i++ {
		n = append(n, i)
	}

	return n
}

func copyCave(cave map[Point]bool) map[Point]bool {
	cpy := make(map[Point]bool)
	for k, v := range cave {
		cpy[k] = v
	}

	return cpy
}
