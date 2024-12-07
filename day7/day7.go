package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cforte7/aoc2024/helpers"
)

type Problem struct {
	Args     []int
	Solution int
}

func (p *Problem) CanBeSolved() bool {
	return false
}

func (p *Problem) getScore() int {
	if p.CanBeSolved() {
		return p.Solution
	}
	return 0
}

func parseData(data []string) []Problem {
	out := make([]Problem, 0)
	for _, v := range data {
		splitOne := strings.Split(v, ":")
		s := helpers.ListAtoi(splitOne[1])
		num, _ := strconv.Atoi(splitOne[0])
		out = append(out, Problem{Solution: num, Args: s})
	}
	return out
}

func partOne(data []Problem) int {
	total := 0
	for _, v := range data {
		total += v.getScore()
	}
	return total
}

func main() {
	data := helpers.Txt_to_lines("test.txt")
	parsedData := parseData(data)
	ans1 := partOne(parsedData)
	fmt.Println(ans1)
}
