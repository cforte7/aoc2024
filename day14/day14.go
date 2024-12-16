package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/cforte7/aoc2024/helpers"
)

const (
	RowCount = 103
	ColCount = 101
)

type XY struct {
	X int
	Y int
}

type Robot struct {
	Pos XY
	Vel XY
}

func (r *Robot) updatePos(n int) {
	newX := (r.Pos.X + r.Vel.X*n) % ColCount
	newY := (r.Pos.Y + r.Vel.Y*n) % RowCount

	if newX < 0 {
		newX += ColCount
	}
	if newY < 0 {
		newY += RowCount
	}
	r.Pos = XY{newX, newY}
}

func (r Robot) getQuadrant() int {
	midRow := RowCount / 2
	midCol := ColCount / 2
	x := r.Pos.X
	y := r.Pos.Y
	var out int
	if r.Pos.X == midCol || r.Pos.Y == midRow {
		return 0
	}

	if x < midCol {
		if y < midRow {
			out = 4
		} else {
			out = 3
		}
	} else if y < midRow {
		out = 1
	} else {
		out = 2
	}
	return out
}

func parseInput(data []string) []Robot {
	out := make([]Robot, 0)
	re := regexp.MustCompile(`-?\d+`)
	for _, v := range data {
		res := re.FindAllString(v, -1)
		px, _ := strconv.Atoi(res[0])
		py, _ := strconv.Atoi(res[1])
		vx, _ := strconv.Atoi(res[2])
		vy, _ := strconv.Atoi(res[3])
		out = append(out, Robot{XY{px, py}, XY{vx, vy}})
	}
	return out
}

func DisplayRobots(robots []Robot) {
	display := make([][]string, RowCount)
	for i := range display {
		newRow := make([]string, ColCount)
		for j := range newRow {
			newRow[j] = "."
		}
		display[i] = newRow
	}
	for _, v := range robots {
		display[v.Pos.Y][v.Pos.X] = "X"
	}

	helpers.Print2DArray(display)
}

func partOne(robots []Robot) int {
	quads := make(map[int]int)
	quads[1] = 0
	quads[2] = 0
	quads[3] = 0
	quads[4] = 0
	display := make([][]int, RowCount)
	for i := range display {
		newRow := make([]int, ColCount)
		display[i] = newRow
	}

	for _, v := range robots {
		v.updatePos(100)
		q := v.getQuadrant()
		display[v.Pos.Y][v.Pos.X] += 1
		quads[q] += 1
	}

	return quads[1] * quads[2] * quads[3] * quads[4]
}

func calcDist(robots []Robot) float64 {
	dist := 0.0
	n := float64(len(robots))
	for i := range robots {
		x1 := robots[i].Pos.X
		y1 := robots[i].Pos.Y
		for j := i + 1; j < len(robots); j++ {

			x2 := robots[j].Pos.X
			y2 := robots[j].Pos.Y

			dx := x2 - x1
			dy := y2 - y1
			d := float64(dx*dx + dy*dy)
			dist += math.Sqrt(d)
		}
	}
	return 2.0 * dist / (n * (n - 1.0))
}

func partTwo(robots []Robot, count int) {
	fmt.Println(robots[0])
	for i := 0; i < count; i++ {
		for j := range robots {
			robots[j].updatePos(1)
		}

		dist := calcDist(robots)
		if dist < 40.0 {
			fmt.Println(i+1, calcDist(robots))
			DisplayRobots(robots)
			time.Sleep(time.Millisecond * 500)

		}
	}
}

func main() {
	path := os.Args[1]
	count := os.Args[2]
	countInt, _ := strconv.Atoi(count)
	data := helpers.Txt_to_lines(path)
	asRobots := parseInput(data)
	fmt.Println(partOne(asRobots))
	asRobots2 := parseInput(data)
	partTwo(asRobots2, countInt)
}
