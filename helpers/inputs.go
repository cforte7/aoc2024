package helpers

import (
	"os"
	"strconv"
	"strings"
)

func BufferToLines(data string) []string {
	as_list := strings.Split(data, "\n")
	out := make([]string, 0)
	for _, v := range as_list {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func Txt_to_lines(path string) []string {
	data, _ := os.ReadFile(path)
	as_string := string(data)
	return BufferToLines(as_string)
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

func LinesToInts(data []string, sep string) [][]int {
	out := make([][]int, 0)
	for _, v := range data {

		s := strings.Split(v, sep)
		l := make([]int, 0)
		for _, vi := range s {
			num, _ := strconv.Atoi(vi)
			l = append(l, num)
		}
		out = append(out, l)
	}
	return out
}

func StringsToChars(data []string) [][]string {
	out := make([][]string, 0)
	for _, v := range data {
		as_char := strings.Split(v, "")
		out = append(out, as_char)
	}
	return out
}

func ReadFile(path string) string {
	data, _ := os.ReadFile(path)
	as_string := string(data)
	return as_string
}
