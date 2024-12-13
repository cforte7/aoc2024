package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type XY struct {
	X float64
	Y float64
}

type Machine struct {
	A XY
	B XY
	P XY
}

func parseLine(line string) XY {
	s1 := strings.Split(line, ": ")
	s2 := strings.Split(s1[1], ", ")
	x := strings.TrimLeft(s2[0], "X+")
	x = strings.TrimLeft(x, "=")
	y := strings.TrimLeft(s2[1], "Y+")
	y = strings.TrimLeft(y, "=")
	yInt, _ := strconv.ParseFloat(y, 32)
	xInt, _ := strconv.ParseFloat(x, 32)
	return XY{xInt, yInt}
}

func parseInput(path string) []Machine {
	data, _ := os.ReadFile(path)
	asString := string(data)
	asMachines := strings.Split(asString, "\n\n")
	out := make([]Machine, 0)

	for _, v := range asMachines {
		asLines := strings.Split(v, "\n")
		a := asLines[0]
		b := asLines[1]
		prize := asLines[2]
		mach := Machine{parseLine(a), parseLine(b), parseLine(prize)}
		out = append(out, mach)
	}
	return out
}

func cramer(m Machine, addr float64) float64 {
	detCoef := m.A.X*m.B.Y - m.A.Y*m.B.X

	detA := (m.P.X+addr)*m.B.Y - (m.P.Y+addr)*m.B.X
	detB := m.A.X*(m.P.Y+addr) - m.A.Y*(m.P.X+addr)
	A := detA / detCoef
	B := detB / detCoef
	_, divA := math.Modf(A)
	_, divB := math.Modf(B)
	if divB == 0 && divA == 0 {
		return 3*A + B
	}
	return 0
}

func partOne(inputs []Machine) float64 {
	score := 0.0
	for _, v := range inputs {
		score += cramer(v, 0.0)
	}
	return score
}

func partTwo(inputs []Machine) float64 {
	score := 0.0
	for _, v := range inputs {
		score += cramer(v, 10000000000000)
	}
	return score
}

func main() {
	path := os.Args[1]
	machines := parseInput(path)
	one := partOne(machines)
	fmt.Println(one)
	two := partTwo(machines)
	fmt.Println(strconv.FormatFloat(two, 'f', -1, 64))
}
