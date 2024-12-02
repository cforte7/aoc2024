package main

import "github.com/cforte7/aoc2024/helpers"

func desc(row []int) int {
	for i := 0; i < len(row)-1; i++ {
		if row[i+1] >= row[i] {
			return 0
		}
		diff := row[i] - row[i+1]
		if diff < 1 || diff > 3 {
			return 0
		}
	}
	return 1
}

func asc(row []int) int {
	for i := 0; i < len(row)-1; i++ {
		if row[i+1] <= row[i] {
			return 0
		}
		diff := row[i+1] - row[i]
		if diff < 1 || diff > 3 {
			return 0
		}
	}
	return 1
}

func day_one(data [][]int) int {
	total := 0
	for _, v := range data {
		if v[1] > v[0] {
			total += asc(v)
		} else {
			total += desc(v)
		}
	}
	return total
}

func main() {
	data := helpers.Txt_to_lines("input.txt")
	as_ints := helpers.Spaces_to_ints(data)
	print(day_one(as_ints))
}
