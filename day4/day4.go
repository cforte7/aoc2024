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

func checkX(data [][]string, row int, col int) int {
	if row == 0 || col == 0 || row == len(data)-1 || col == len(data[0])-1 {
		return 0
	}
	if data[row][col] != "A" {
		return 0
	}
	pairs := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	for i := range pairs {
		one_one := []int{row + pairs[i][0], col + pairs[i][1]}
		one_two := []int{row + pairs[(i+1)%4][0], col + pairs[(i+1)%4][1]}
		two_one := []int{row + pairs[(i+2)%4][0], col + pairs[(i+2)%4][1]}
		two_two := []int{row + pairs[(i+3)%4][0], col + pairs[(i+3)%4][1]}

		c11 := data[one_one[0]][one_one[1]]
		c12 := data[one_two[0]][one_two[1]]
		c21 := data[two_one[0]][two_one[1]]
		c22 := data[two_two[0]][two_two[1]]
		if c11 == c12 && c21 == c22 && (c11 == "M" && c22 == "S" || c11 == "S" && c22 == "M") {
			return 1
		}
	}
	return 0
}

func partTwo(data [][]string, startRow int, startCol int) int {
	if startRow == 0 || startCol == 0 || startRow == len(data)-1 || startCol == len(data[0])-1 {
		return 0
	}

	for row := 1; row < len(data)-1; row++ {
		for col := 1; col < len(data[row])-1; col++ {
			if data[row][col] != "A" {
				continue
			}
		}
	}
	return 0
}

func partOne(data [][]string) (int, int) {
	total := 0
	total2 := 0
	// dirs := [8][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			total2 += checkX(data, row, col)
			// for _, dir := range dirs {
			// 	total += checkDirection(data, row, col, dir)
			// }
		}
	}
	return total, total2
}

func main() {
	asLines := helpers.Txt_to_lines("input.txt")
	asChars := helpers.StringsToChars(asLines)
	fmt.Println(partOne(asChars))
}

// 2305 is too high
