package main

import (
	"github.com/cforte7/aoc2024/helpers"
)

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

func asc_is_valid(val1 int, val2 int) bool {
	diff := val2 - val1
	is_asc := val2 > val1
	return is_asc && (diff >= 1 && diff <= 3)
}

func desc_is_valid(val1 int, val2 int) bool {
	diff := val1 - val2
	is_desc := val2 < val1
	return is_desc && (diff >= 1 && diff <= 3)
}

func row_val_asc(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if !asc_is_valid(row[i], row[i+1]) {
			return false
		}
	}
	return true
}

func row_val_desc(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if !desc_is_valid(row[i], row[i+1]) {
			return false
		}
	}
	return true
}

func check_row(row []int) bool {
	if row_val_asc(row) || row_val_desc(row) {
		return true
	}
	rowCopy := make([]int, len(row))
	copy(rowCopy, row)
	for index := 0; index < len(row); index++ {
		rowCopy := make([]int, len(row))
		copy(rowCopy, row)
		subset_row := append(rowCopy[:index], rowCopy[index+1:]...)
		if row_val_asc(subset_row) || row_val_desc(subset_row) {
			return true
		}
	}
	return false
}

func day_two(data [][]int) int {
	total := 0

	for _, v := range data {
		if check_row(v) {
			total += 1
		}
	}
	return total
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
	print(day_one(as_ints), "\n")
	print(day_two(as_ints), "\n")
}
