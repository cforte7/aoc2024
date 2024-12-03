package helpers

import (
	"os"
	"strconv"
	"strings"
)

func Txt_to_lines(path string) []string {
	data, _ := os.ReadFile(path)
	as_string := string(data)
	as_list := strings.Split(as_string, "\n")
	out := make([]string, 0)
	for _, v := range as_list {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func Spaces_to_ints(data []string) [][]int {
	// turn something like ["5 2", "1  2   3 4"] into [[5,2],[1,2,3,4]]
	out := make([][]int, 0)
	for _, v := range data {

		s := strings.Fields(v)
		l := make([]int, 0)
		for _, vi := range s {
			num, _ := strconv.Atoi(vi)
			l = append(l, num)
		}
		out = append(out, l)
	}
	return out
}

func ReadFile(path string) string {
	data, _ := os.ReadFile(path)
	as_string := string(data)
	return as_string
}
