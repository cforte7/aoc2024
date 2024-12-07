package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cforte7/aoc2024/helpers"
)

type Problem struct {
	Args     []int
	Solution int
}

func CanBeSolved(solu, tot int, args []int) bool {
	if len(args) == 0 {
		return solu == tot
	}

	toApply := args[0]
	restArgs := args[1:]

	multTot := tot * toApply
	addTot := tot + toApply

	return CanBeSolved(solu, multTot, restArgs) || CanBeSolved(solu, addTot, restArgs)
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
		startVal := v.Args[0]
		restArgs := v.Args[1:]
		if CanBeSolved(v.Solution, startVal, restArgs) {
			total += v.Solution
		}
	}
	return total
}

func main() {
	path := os.Args[1]
	data := helpers.Txt_to_lines(path)
	parsedData := parseData(data)
	ans1 := partOne(parsedData)
	fmt.Println(ans1)
}
