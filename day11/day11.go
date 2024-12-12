package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cforte7/aoc2024/helpers"
)

func updateStone(stone int) []int {
	asStr := strconv.Itoa(stone)
	if stone == 0 {
		return []int{1}
	} else if len(asStr)%2 == 0 {
		lnum, _ := strconv.Atoi(asStr[:len(asStr)/2])
		rnum, _ := strconv.Atoi(asStr[len(asStr)/2:])
		return []int{lnum, rnum}
	}

	return []int{stone * 2024}
}

var cache = make(map[[2]int]int)

func recurUpdateStone(stone int, count int) int {
	if count == 0 {
		return 1
	}

	cachedVal, ok := cache[[2]int{stone, count}]
	if ok {
		return cachedVal
	}

	newCount := count - 1

	asStr := strconv.Itoa(stone)
	if stone == 0 {
		out := recurUpdateStone(1, newCount)
		cache[[2]int{stone, count}] = out
		return out
	} else if len(asStr)%2 == 0 {
		lnum, _ := strconv.Atoi(asStr[:len(asStr)/2])
		rnum, _ := strconv.Atoi(asStr[len(asStr)/2:])
		out := recurUpdateStone(lnum, newCount) + recurUpdateStone(rnum, newCount)
		cache[[2]int{stone, count}] = out
		return out
	}

	out := recurUpdateStone(stone*2024, newCount)
	cache[[2]int{stone, count}] = out
	return out
}

func partOne(data []int) int {
	for range 25 {
		new := make([]int, 0)
		for i := range data {
			new = append(new, updateStone(data[i])...)
		}
		data = new
	}
	return len(data)
}

func partTwo(data []int) int {
	score := 0
	for _, v := range data {
		score += recurUpdateStone(v, 75)
	}
	return score
}

func main() {
	path := os.Args[1]
	data := helpers.ReadFile(path)
	asInts := helpers.ListAtoi(data)
	asInts2 := helpers.ListAtoi(data)

	fmt.Println(partOne(asInts))
	fmt.Println(partTwo(asInts2))
}
