package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	packets := parsePackets(input)
	fmt.Println("Answer for Part 1:", sumIndices(packets))
	fmt.Println("Answer for Part 2:", getDecoderKey(packets))
}

func parsePackets(input string) []any {
	packets := []any{}
	split := strings.Split(strings.TrimSpace(input), "\n\n")
	for _, pair := range split {
		fields := strings.Fields(pair)
		var left, right []any

		// too lazy to write a parser :(
		json.Unmarshal([]byte(fields[0]), &left)
		json.Unmarshal([]byte(fields[1]), &right)

		packets = append(packets, left, right)
	}

	return packets
}

func getDecoderKey(packets []any) int {
	var div1, div2 []any
	json.Unmarshal([]byte("[[2]]"), &div1)
	json.Unmarshal([]byte("[[6]]"), &div2)

	packets = append(packets, div1, div2)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	res := 0
	for i, p := range packets {
		if compare(p, div1) == 0 {
			res += i + 1
		} else if compare(p, div2) == 0 {
			res *= i + 1
		}
	}

	return res
}

func sumIndices(packets []any) int {
	sum := 0
	for i := 0; i < len(packets); i += 2 {
		if compare(packets[i], packets[i+1]) < 0 {
			sum += i/2 + 1
		}
	}

	return sum
}

func compare(left any, right any) int {
	if isNum(left) && isNum(right) {
		return compareNums(left, right)
	}

	if isNum(left) {
		left = []any{left}
	}

	if isNum(right) {
		right = []any{right}
	}

	return compareLists(left, right)
}

func isNum(v any) bool {
	_, ok := v.(float64)
	return ok
}

func compareNums(a, b any) int {
	return int(a.(float64) - b.(float64))
}

func compareLists(a, b any) int {
	al := a.([]any)
	bl := b.([]any)

	for i := range al {
		if len(al) == i || len(bl) == i {
			return len(al) - len(bl)
		}

		res := compare(al[i], bl[i])
		if res != 0 {
			return res
		}
	}

	return len(al) - len(bl)
}
