package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(lnum []int, rnum []int) int {
	sort.Ints(lnum)
	sort.Ints(rnum)

	total := 0
	for i := 0; i < len(lnum); i++ {
		diff := lnum[i] - rnum[i]
		if diff < 0 {
			total += -diff
		} else {
			total += diff
		}
	}
	return total
}

func parse(data []string) ([]int, []int) {
	lnum := make([]int, len(data))
	rnum := make([]int, len(data))

	for _, v := range data {

		if v == "" {
			continue
		}
		s := strings.Fields(v)
		l, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])
		lnum = append(lnum, l)
		rnum = append(rnum, r)
	}
	return lnum, rnum
}

func part_two(l []int, r []int) int {
	m := make(map[int]int)
	for _, v := range r {
		curr := m[v]
		m[v] = curr + 1
	}
	score := 0
	for _, v := range l {
		score += v * m[v]
	}
	return score
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)
	as_string := string(data)
	as_list := strings.Split(as_string, "\n")
	lparsed, rparsed := parse(as_list)
	res1 := part_one(lparsed, rparsed)
	lparsed, rparsed = parse(as_list)
	res2 := part_two(lparsed, rparsed)
	print(res1, "    ", res2)
}
