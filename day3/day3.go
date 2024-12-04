package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/cforte7/aoc2024/helpers"
)

func parse_and_mult(v string) int {
	s := strings.Split(v, ",")
	val1, _ := strconv.Atoi(strings.Trim(s[0], "mul("))
	val2, _ := strconv.Atoi(strings.Trim(s[1], ")"))
	return val1 * val2
}

func part1(data string) int {
	r, _ := regexp.Compile(`mul\((-?\d+),(-?\d+)\)`)
	vals := r.FindAllString(data, -1)
	total := 0
	for _, v := range vals {
		total += parse_and_mult(v)
	}
	return total
}

func part2(data string) int {
	r, _ := regexp.Compile(`mul\((-?\d+),(-?\d+)\)|do\(\)|don't\(\)`)
	vals := r.FindAllString(data, -1)

	total := 0
	should_mult := true
	for _, v := range vals {
		if v == "do()" {
			should_mult = true
		} else if v == "don't()" {
			should_mult = false
		} else if should_mult {
			total += parse_and_mult(v)
		}
	}
	return total
}

func main() {
	data := helpers.ReadFile("input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}
