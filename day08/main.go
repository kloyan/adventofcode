package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rows := strings.Fields(input)

	fmt.Println("Answer for Part 1:", calcVisibleTrees(rows))
	fmt.Println("Answer for Part 2:", getHighestScenicScore(rows))
}

func calcVisibleTrees(rows []string) int {
	count := 2*(len(rows)+len(rows[0])) - 4

	for row := range rows {
		if row == 0 || row == len(rows)-1 {
			continue
		}

		for col := range rows[row] {
			if col == 0 || col == len(rows[row])-1 {
				continue
			}

			tree := rows[row][col]

			if visibleInRow(rows[row][:col], tree) {
				count++
			} else if visibleInRow(rows[row][col+1:], tree) {
				count++
			} else if visibleInCol(rows[:row], col, tree) {
				count++
			} else if visibleInCol(rows[row+1:], col, tree) {
				count++
			}
		}
	}

	return count
}

func visibleInRow(row string, tree byte) bool {
	for i := range row {
		if row[i] >= tree {
			return false
		}
	}

	return true
}

func visibleInCol(rows []string, col int, tree byte) bool {
	for i := range rows {
		if rows[i][col] >= tree {
			return false
		}
	}

	return true
}

func getHighestScenicScore(rows []string) int {
	highest := 0
	for row := range rows {
		if row == 0 || row == len(rows)-1 {
			continue
		}

		for col := range rows[row] {
			if col == 0 || col == len(rows[row])-1 {
				continue
			}

			score := 1
			tree := rows[row][col]

			score *= countTreesLeft(rows[row][:col], tree)
			score *= countTreesRight(rows[row][col+1:], tree)
			score *= countTreesTop(rows[:row], col, tree)
			score *= countTreesBottom(rows[row+1:], col, tree)

			if score > highest {
				highest = score
			}
		}
	}

	return highest
}

func countTreesLeft(row string, tree byte) int {
	trees := 0
	for i := len(row) - 1; i >= 0; i-- {
		trees++
		if row[i] >= tree {
			break
		}
	}

	return trees
}

func countTreesRight(row string, tree byte) int {
	trees := 0
	for i := 0; i < len(row); i++ {
		trees++
		if row[i] >= tree {
			break
		}
	}

	return trees
}

func countTreesTop(rows []string, col int, tree byte) int {
	trees := 0
	for i := len(rows) - 1; i >= 0; i-- {
		trees++
		if rows[i][col] >= tree {
			break
		}
	}

	return trees
}

func countTreesBottom(rows []string, col int, tree byte) int {
	trees := 0
	for i := 0; i < len(rows); i++ {
		trees++
		if rows[i][col] >= tree {
			break
		}
	}

	return trees
}
