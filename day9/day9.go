package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func vizData(data []Segment) {
	for _, v := range data {
		if v.Id == -1 {
			for i := 0; i < v.Size; i++ {
				fmt.Print(".")
			}
		} else {
			for i := 0; i < v.Size; i++ {
				fmt.Print(v.Id)
			}
		}
	}
	fmt.Println("")
}

func calcChecksum(data []Segment) int {
	index := 0
	checksum := 0
	for _, v := range data {
		for i := 0; i < v.Size; i++ {
			checksum += index * v.Id
			index += 1
		}
	}
	return checksum
}

func partOne(data []Segment) int {
	lPointer := 0
	rPointer := len(data) - 1
	out := make([]Segment, 0)

	for lPointer < rPointer {
		// find next freespace
		for data[lPointer].Id != -1 || data[lPointer].Size == 0 {
			out = append(out, data[lPointer])
			lPointer++
		}

		// consume the free space
		for data[lPointer].Size > 0 {
			// find the next data segment
			for data[rPointer].Id == -1 || data[rPointer].Size == 0 {
				rPointer--
			}
			if lPointer >= rPointer {
				break
			}

			// consume entire memory if we can
			if data[lPointer].Size >= data[rPointer].Size {
				out = append(out, data[rPointer])
				data[lPointer].Size -= data[rPointer].Size
				data[rPointer].Size = 0
			} else {
				// else consume partial memory
				partial := Segment{Size: data[lPointer].Size, Id: data[rPointer].Id}
				out = append(out, partial)
				data[rPointer].Size -= data[lPointer].Size
				data[lPointer].Size = 0
			}
		}
	}

	return calcChecksum(out)
}

type Segment struct {
	Id   int
	Size int
}

func parseInput(data []byte) []Segment {
	asChars := strings.Split(string(data), "")
	out := []Segment{}
	for i := 0; i < len(asChars); i++ {
		asInt, _ := strconv.Atoi(asChars[i])
		isFreeSpace := i%2 == 1
		var toAppend Segment
		if isFreeSpace {
			toAppend = Segment{Id: -1, Size: asInt}
		} else {
			toAppend = Segment{Id: i / 2, Size: asInt}
		}
		out = append(out, toAppend)
	}
	return out
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	asSeg := parseInput(data)
	fmt.Println(partOne(asSeg))
}
