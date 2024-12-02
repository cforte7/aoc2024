package main

import (
	"fmt"

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

func desc_comp(val1 int, val2 int) bool {
	diff := val1 - val2
	is_asc := val2 >= val1
	return is_asc || diff < 1 || diff > 3
}

func desc2(row []int) bool {
	ever_damp := false
	for i := 0; i < len(row)-1; i++ {
		bad := desc_comp(row[i], row[i+1])
		if !bad {
			continue
		}

		if bad && ever_damp {
			return false
		}
		if bad {
			ever_damp = true
			if i == 0 {
				continue
			}
			if desc_comp(row[i-1], row[i+1]) {
				return false
			}
		}
	}
	return true
}

func asc_comp(val1 int, val2 int) bool {
	diff := val2 - val1
	is_desc := val2 <= val1
	return is_desc || diff < 1 || diff > 3
}

func asc2(row []int) bool {
	ever_damp := false
	for i := 0; i < len(row)-1; i++ {
		bad := asc_comp(row[i], row[i+1])
		if !bad {
			continue
		}

		if bad && ever_damp {
			return false
		}
		if bad {
			ever_damp = true
			if i == 0 {
				continue
			}
			if asc_comp(row[i-1], row[i+1]) {
				return false
			}
		}
	}
	return true
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

func day_two(data [][]int) int {
	total := 0
	for _, v := range data {
		if asc2(v) || desc2(v) {
			total += 1
			continue
		}
		fmt.Println("bad: ", v)
	}
	return total
}

func main() {
	data := helpers.Txt_to_lines("input.txt")
	as_ints := helpers.Spaces_to_ints(data)
	// print(len(as_ints))
	// print(day_one(as_ints))
	print(day_two(as_ints))
}
