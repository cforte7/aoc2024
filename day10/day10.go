package main

import (
	"fmt"
	"os"

	"github.com/cforte7/aoc2024/helpers"
)

func findPeaks(i, j int, peaks [][2]int, data [][]int) [][2]int {
	if data[i][j] == 9 {
		return append(peaks, [2]int{i, j})
	}

	// up
	if helpers.IsInBounds(i+1, j, data) && data[i+1][j] == data[i][j]+1 {
		peaks = findPeaks(i+1, j, peaks, data)
	}

	if helpers.IsInBounds(i-1, j, data) && data[i-1][j] == data[i][j]+1 {
		peaks = findPeaks(i-1, j, peaks, data)
	}
	if helpers.IsInBounds(i, j+1, data) && data[i][j+1] == data[i][j]+1 {
		peaks = findPeaks(i, j+1, peaks, data)
	}
	if helpers.IsInBounds(i, j-1, data) && data[i][j-1] == data[i][j]+1 {
		peaks = findPeaks(i, j-1, peaks, data)
	}
	return peaks
}

func partOne(data [][]int) (int, int) {
	score := 0
	score2 := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == 0 {

				uniquePeaks := make(map[[2]int]bool, 0)
				peaks := findPeaks(i, j, make([][2]int, 0), data)
				for _, v := range peaks {
					uniquePeaks[v] = true
				}
				score += len(uniquePeaks)
				score2 += len(peaks)
			}
		}
	}
	return score, score2
}

func main() {
	path := os.Args[1]
	data := helpers.Txt_to_lines(path)
	asInts := helpers.LinesToInts(data, "")
	fmt.Println(partOne(asInts))
}
