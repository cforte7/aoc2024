package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/cforte7/aoc2024/helpers"
)

func part1(data string) int {
	r, _ := regexp.Compile(`mul\((-?\d+),(-?\d+)\)`)
	vals := r.FindAllString(data, -1)
	total := 0
	for _, v := range vals {
		// fmt.Println(v)
		s := strings.Split(v, ",")
		val1, _ := strconv.Atoi(strings.Trim(s[0], "mul("))
		val2, _ := strconv.Atoi(strings.Trim(s[1], ")"))
		total += val1 * val2
	}

	return total
}

func main() {
	data := helpers.ReadFile("input.txt")
	fmt.Println(part1(data))
}
