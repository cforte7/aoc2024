package main

import (
	"fmt"
	"os"

	"github.com/cforte7/aoc2024/helpers"
)

type Position struct {
	row int
	col int
}

func newPos(r, c int) Position {
	return Position{row: r, col: c}
}

func getRegion(data [][]string, start Position, value string) ([]Position, int) {
	out := make([]Position, 0)
	seen := make(map[Position]bool)
	toVisit := make([]Position, 0)
	outside := make(map[[2]Position]bool)
	toVisit = append(toVisit, start)
	var curr Position
	seen[start] = true
	for len(toVisit) > 0 {
		curr = toVisit[0]
		toVisit = toVisit[1:]
		out = append(out, curr)
		for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			toCheck := newPos(curr.row+v[0], curr.col+v[1])
			_, seenToCheck := seen[toCheck]
			if helpers.IsInBounds(toCheck.row, toCheck.col, data) && value == data[toCheck.row][toCheck.col] && !seenToCheck {
				toVisit = append(toVisit, toCheck)
				seen[toCheck] = true
			} else if !helpers.IsInBounds(toCheck.row, toCheck.col, data) || value != data[toCheck.row][toCheck.col] {
				outside[[2]Position{curr, toCheck}] = true
			}
		}
	}

	return out, len(outside) * len(out)
}

func partOne(data [][]string) int {
	regions := make([][]Position, 0)
	seen := make(map[Position]bool)
	score := 0
	for row := range data {
		for col := range data[0] {
			pos := newPos(row, col)
			_, visted := seen[pos]
			if visted {
				continue
			}
			newRegion, newScore := getRegion(data, pos, data[pos.row][pos.col])
			regions = append(regions, newRegion)
			score += newScore
			for _, newPos := range newRegion {
				seen[newPos] = true
			}
		}
	}
	// this is wrong, its just to get the compiler to stfu
	return score
}

func main() {
	path := os.Args[1]
	data := helpers.Txt_to_lines(path)
	asChars := helpers.StringsToChars(data)
	fmt.Println(partOne(asChars))
}
