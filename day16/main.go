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

type valve struct {
	idx       int
	neighbors []int
	rate      int
}

type key struct {
	idx  int
	time int
	mask int
}

func main() {
	re := regexp.MustCompile(`Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? ((?:[A-Z]+,?)+)`)
	indices := map[string]int{}
	split := strings.Split(strings.TrimSpace(input), "\n")

	start := 0
	for i, str := range split {
		var v string
		fmt.Sscanf(str, "Valve %s has flow rate", &v)
		indices[v] = i

		if v == "AA" {
			start = i
		}
	}

	valves := []valve{}
	for _, str := range split {
		m := re.FindStringSubmatch(strings.ReplaceAll(str, ", ", ","))
		valve := valve{
			idx:  indices[m[1]],
			rate: common.ToInt(m[2]),
		}
		for _, n := range strings.Split(m[3], ",") {
			valve.neighbors = append(valve.neighbors, indices[n])
		}
		valves = append(valves, valve)
	}

	cache := map[key]int{}
	ans := dfs(valves[start], 30, 0, cache, valves)

	fmt.Println("Answer for Part 1:", ans)
}

func dfs(curr valve, time int, mask int, cache map[key]int, valves []valve) int {
	if time < 1 {
		return 0
	}

	key := key{idx: curr.idx, time: time, mask: mask}
	if m, ok := cache[key]; ok {
		return m
	}

	ans := 0
	bit := 1 << curr.idx
	if (mask&bit == 0) && curr.rate > 0 {
		ans = (time-1)*curr.rate + dfs(curr, time-1, mask|bit, cache, valves)
	}

	for _, n := range curr.neighbors {
		ans = max(ans, dfs(valves[n], time-1, mask, cache, valves))
	}

	cache[key] = ans
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
