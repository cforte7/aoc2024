package main

import (
	"fmt"

	"github.com/cforte7/aoc2024/helpers"
)

func findStart(data [][]string) [2]int {
	for r, row := range data {
		for c, col := range row {
			if col == "^" {
				return [2]int{r, c}
			}
		}
	}
	panic("no start found!")
}

func leavingMap(pos, dir [2]int, data [][]string) bool {
	maxRow := len(data) - 1
	maxCol := len(data[0]) - 1
	nextPos := getNextPost(pos, dir)

	if nextPos[0] == -1 || nextPos[1] == -1 || nextPos[0] == maxRow+1 || nextPos[1] == maxCol+1 {
		return true
	}
	return false
}

func canContinue(pos, dir [2]int, data [][]string) bool {
	nextRow := pos[0] + dir[0]
	nextCol := pos[1] + dir[1]

	return data[nextRow][nextCol] != "#"
}

func getNextPost(pos, dir [2]int) [2]int {
	return [2]int{pos[0] + dir[0], pos[1] + dir[1]}
}

func rotateDir(dir [2]int) [2]int {
	return [2]int{dir[1], -dir[0]}
}

func dayOne(data [][]string) int {
	pos := findStart(data)
	dir := [2]int{-1, 0}
	visited := make(map[[2]int]bool)
	for !leavingMap(pos, dir, data) {
		visited[[2]int{pos[0], pos[1]}] = true
		if !canContinue(pos, dir, data) {
			dir = rotateDir(dir)
		}
		pos = getNextPost(pos, dir)
	}

	visited[[2]int{pos[0], pos[1]}] = true
	return len(visited)
}

func main() {
	data := helpers.Txt_to_lines("input.txt")
	asChars := helpers.StringsToChars(data)
	fmt.Println(dayOne(asChars))
}
