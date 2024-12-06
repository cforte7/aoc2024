package main

import (
	"fmt"
	"strings"

	"github.com/cforte7/aoc2024/helpers"
)

func parseInput() ([][]int, [][]int) {
	data := helpers.ReadFile("input.txt")
	inTwo := strings.Split(data, "\n\n")

	rules := helpers.BufferToLines(inTwo[0])
	pages := helpers.BufferToLines(inTwo[1])

	parsedRules := helpers.LinesToInts(rules, "|")
	parsedPages := helpers.LinesToInts(pages, ",")
	return parsedRules, parsedPages
}

func ruleDict(rules [][]int) map[int](map[int]bool) {
	shapedRules := make(map[int](map[int]bool))

	for _, v := range rules {
		if shapedRules[v[0]] == nil {
			shapedRules[v[0]] = map[int]bool{}
		}

		shapedRules[v[0]][v[1]] = true
	}
	return shapedRules
}

func getMiddle(page []int) int {
	return page[len(page)/2]
}

func isNext(pageSubset []int, numRules map[int]bool) bool {
	for _, v := range pageSubset {
		if !numRules[v] {
			return false
		}
	}
	return true
}

func evalPage(toAdd []int, rules map[int]map[int]bool) (int, []int) {
	for i, v := range toAdd {
		numRules := rules[v]
		if i == len(toAdd)-1 || isNext(toAdd[i+1:], numRules) {
			out := make([]int, 0)
			out = append(toAdd[:i], toAdd[i+1:]...)
			return v, out
		}
	}
	return 0, toAdd
}

func partTwo(page []int, rules map[int]map[int]bool) int {
	toAdd := make([]int, len(page))
	copy(toAdd, page)
	out := make([]int, 0)
	for len(toAdd) > 1 {
		newVal, newList := evalPage(toAdd, rules)
		toAdd = newList
		out = append(out, newVal)
	}
	out = append(out, toAdd[0])
	return getMiddle(out)
}

func checkPage(page []int, shapedRules map[int]map[int]bool) int {
	for i := range page {

		numRules := shapedRules[page[i]]
		if i == len(page)-1 {
			continue
		}

		toCheck := page[i+1:]
		for _, c := range toCheck {
			if !numRules[c] {
				return 0
			}
		}
	}
	return getMiddle(page)
}

func partOne(rules [][]int, pages [][]int) (int, int) {
	shapedRules := ruleDict(rules)
	total := 0
	total2 := 0
	for _, page := range pages {
		partOneVal := checkPage(page, shapedRules)
		total += partOneVal
		if partOneVal == 0 {
			total2 += partTwo(page, shapedRules)
		}
	}
	return total, total2
}

func main() {
	rules, pages := parseInput()
	fmt.Println(partOne(rules, pages))
}
