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

func checkPage(page []int, shapedRules map[int]map[int]bool) int {
	for i := range page {

		numRules := shapedRules[page[i]]
		if i == len(page)-1 {
			continue
		}

		toCheck := page[i+1:]
		// fmt.Println(numRules)
		// fmt.Println(toCheck)
		for _, c := range toCheck {
			if !numRules[c] {
				// fmt.Println(c, " not found in rules")
				return 0
			}
		}
	}
	return getMiddle(page)
}

func partOne(rules [][]int, pages [][]int) int {
	shapedRules := ruleDict(rules)
	total := 0
	for _, page := range pages {
		total += checkPage(page, shapedRules)
	}
	return total
}

func main() {
	rules, pages := parseInput()
	fmt.Println(partOne(rules, pages))
}
