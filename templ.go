package main

import (
	"fmt"
	"os"

	"github.com/cforte7/aoc2024/helpers"
)

func main() {
	path := os.Args[1]
	data := helpers.Txt_to_lines(path)
	fmt.Println(data)
}
