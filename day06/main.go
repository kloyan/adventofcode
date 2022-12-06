package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Answer for Part 1:", findMarker(input, 4))
	fmt.Println("Answer for Part 2:", findMarker(input, 14))
}

func findMarker(input string, size int) int {
	for i := size; i < len(input); i++ {
		str := input[i-size : i]
		if isUnique(str) {
			return i
		}
	}

	panic(fmt.Errorf("missing marker"))
}

func isUnique(str string) bool {
	mask := 0
	for i := range str {
		bit := 1 << (str[i] - byte('a'))
		if (mask & bit) > 0 {
			return false
		}

		mask |= bit
	}

	return true
}
