package main

import (
	"fmt"
	"strings"

	"github.com/cforte7/aoc2024/helpers"
)

func checkDirection(data [][]string, startRow int, startCol int, dir [2]int) int {
	endRow := startRow + dir[0]*3
	endCol := startCol + dir[1]*3
	maxRow := len(data) - 1
	maxCol := len(data[0]) - 1

	if endRow > maxRow || endCol > maxCol || endRow < 0 || endCol < 0 {
		return 0
	}
	toCheck := make([]string, 0)

	rowInds := make([]int, 0)
	colInds := make([]int, 0)
	for rowAug := 0; rowAug <= 3; rowAug++ {
		rowInds = append(rowInds, startRow+rowAug*dir[0])
	}
	for colAug := 0; colAug <= 3; colAug++ {
		colInds = append(colInds, startCol+colAug*dir[1])
	}

	for i := range rowInds {
		rowTarget := rowInds[i]
		colTarget := colInds[i]
		toCheck = append(toCheck, data[rowTarget][colTarget])
	}
	joinedCheck := strings.Join(toCheck, "")
	if joinedCheck == "XMAS" {
		return 1
	}
	return 0
}

func partOne(data [][]string) int {
	total := 0
	dirs := [8][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			for _, dir := range dirs {
				total += checkDirection(data, row, col, dir)
			}
		}
	}
	return total
}

func main() {
	asLines := helpers.Txt_to_lines("input.txt")
	asChars := helpers.StringsToChars(asLines)
	fmt.Println(partOne(asChars))
}
