package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/kloyan/adventofcode/common"
)

//go:embed input.txt
var input string

type key struct {
	idx          int
	time         int
	mask         int
	withElephant bool
}

func main() {
	re := regexp.MustCompile(`Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? ((?:[A-Z]+,?)+)`)
	valves := map[string]int{}
	tunnels := map[string][]string{}

	for _, str := range strings.Split(strings.TrimSpace(input), "\n") {
		m := re.FindStringSubmatch(strings.ReplaceAll(str, ", ", ","))
		valves[m[1]] = common.ToInt(m[2])
		tunnels[m[1]] = strings.Split(m[3], ",")
	}

	indices := map[string]int{}
	for v := range valves {
		indices[v] = len(indices)
	}

	cache := map[key]int{}
	part1 := dfs("AA", 30, 0, false, cache, indices, valves, tunnels)
	fmt.Println("Answer for Part 1:", part1)

	part2 := dfs("AA", 26, 0, true, cache, indices, valves, tunnels)
	fmt.Println("Answer for Part 2:", part2)
}

func dfs(curr string, time int, mask int, withElephant bool, cache map[key]int, indices map[string]int, valves map[string]int, tunnels map[string][]string) int {
	if time == 0 {
		if withElephant {
			return dfs("AA", 26, mask, false, cache, indices, valves, tunnels)
		}
		return 0
	}

	key := key{idx: indices[curr], time: time, mask: mask, withElephant: withElephant}
	if c, ok := cache[key]; ok {
		return c
	}

	val := 0
	bit := 1 << indices[curr]
	if mask&bit == 0 && valves[curr] > 0 {
		val = (time-1)*valves[curr] + dfs(curr, time-1, mask|bit, withElephant, cache, indices, valves, tunnels)
	}

	for _, n := range tunnels[curr] {
		val = max(val, dfs(n, time-1, mask, withElephant, cache, indices, valves, tunnels))
	}

	cache[key] = val
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
