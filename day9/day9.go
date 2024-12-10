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
			if v.Id != -1 {
				checksum += index * v.Id
			}
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

func partTwo(data []Segment) int {
	toMove := len(data) - 1
	for toMove >= 0 {
		// find the next data segment
		for data[toMove].Id == -1 || data[toMove].Size == 0 {
			toMove--
		}
		dest := 0

		for dest < toMove {
			if data[dest].Id == -1 && data[dest].Size >= data[toMove].Size {

				segToMove := Segment{Id: data[toMove].Id, Size: data[toMove].Size}
				newEmptySpace := Segment{Size: data[dest].Size - data[toMove].Size, Id: -1}
				newData := make([]Segment, 0)
				newData = append(newData, data[:dest]...)
				newData = append(newData, segToMove, newEmptySpace)
				data[toMove] = Segment{Id: -1, Size: data[toMove].Size}
				newData = append(newData, data[dest+1:]...)
				data = newData
				toMove++
				break
			}
			dest++
		}
		toMove--
	}
	return calcChecksum(data)
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
	asSegTwo := parseInput(data)
	fmt.Println(partTwo(asSegTwo))
}
