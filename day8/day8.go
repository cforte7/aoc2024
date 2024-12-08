package main

import (
	"fmt"
	"os"

	"github.com/cforte7/aoc2024/helpers"
)

type SignalMap struct {
	Antennas  map[string]map[[2]int]bool
	Antinodes map[[2]int]bool
}

func insideBounds(rowMax, colMax, row, col int) bool {
	return row >= 0 && col >= 0 && row <= rowMax && col <= colMax
}

func partOne(data [][]string) (int, int) {
	rowMax := len(data) - 1
	colMax := len(data[0]) - 1
	signalMap := make(map[string]map[[2]int]bool)
	antinodes := make(map[[2]int]bool)
	antinodes2 := make(map[[2]int]bool)
	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data); col++ {
			antennaType := data[row][col]
			if antennaType == "." {
				continue
			}

			antennaLocs := signalMap[antennaType]
			// FTT with an antenna type, init dict, no anitpoles to make
			if antennaLocs == nil {
				antennaLocs = make(map[[2]int]bool)
				antennaLocs[[2]int{row, col}] = true
				signalMap[antennaType] = antennaLocs
				continue
			}

			for k := range antennaLocs {
				rowDiff := k[0] - row
				colDiff := k[1] - col
				antinodeOne := [2]int{k[0] + rowDiff, k[1] + colDiff}
				antinodeTwo := [2]int{row - rowDiff, col - colDiff}
				// part 1
				if insideBounds(rowMax, colMax, antinodeOne[0], antinodeOne[1]) {
					antinodes[antinodeOne] = true
				}
				if insideBounds(rowMax, colMax, antinodeTwo[0], antinodeTwo[1]) {
					antinodes[antinodeTwo] = true
				}

				// part 2
				antinodes2[[2]int{row, col}] = true
				antinodes2[[2]int{k[0], k[1]}] = true
				for insideBounds(rowMax, colMax, antinodeOne[0], antinodeOne[1]) {
					antinodes2[antinodeOne] = true
					antinodeOne = [2]int{antinodeOne[0] + rowDiff, antinodeOne[1] + colDiff}
				}
				for insideBounds(rowMax, colMax, antinodeTwo[0], antinodeTwo[1]) {
					antinodes2[antinodeTwo] = true
					antinodeTwo = [2]int{antinodeTwo[0] - rowDiff, antinodeTwo[1] - colDiff}
				}
			}

			antennaLocs[[2]int{row, col}] = true
			signalMap[antennaType] = antennaLocs

		}
	}

	partOneAns := len(antinodes)
	partTwoAns := len(antinodes2)

	return partOneAns, partTwoAns
}

func main() {
	path := os.Args[1]
	data := helpers.Txt_to_lines(path)
	asChars := helpers.StringsToChars(data)
	fmt.Println(partOne(asChars))
}

// 1042 too low
