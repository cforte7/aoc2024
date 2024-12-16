package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cforte7/aoc2024/helpers"
)

func findStart(data [][]string) helpers.Pos {
	for i, row := range data {
		for j, val := range row {
			if val == "@" {
				return helpers.Pos{Row: i, Col: j}
			}
		}
	}
	panic("didnt find start")
}

func combinePos(a, b helpers.Pos) helpers.Pos {
	return helpers.Pos{Row: a.Row + b.Row, Col: a.Col + b.Col}
}

func executeMove(data *[][]string, move helpers.Pos) {
	currPos := findStart(*data)
	maybeNextPos := combinePos(currPos, move)

	maybeNextChar := (*data)[maybeNextPos.Row][maybeNextPos.Col]
	if maybeNextChar == "." {
		(*data)[maybeNextPos.Row][maybeNextPos.Col] = "@"
		(*data)[currPos.Row][currPos.Col] = "."
		return
	} else if maybeNextChar == "#" {
		return
	}
	terminalPos := maybeNextPos
	for (*data)[terminalPos.Row][terminalPos.Col] == "O" {
		terminalPos = combinePos(terminalPos, move)
	}
	if (*data)[terminalPos.Row][terminalPos.Col] == "." {
		(*data)[terminalPos.Row][terminalPos.Col] = "O"
		(*data)[currPos.Row][currPos.Col] = "."
		(*data)[maybeNextPos.Row][maybeNextPos.Col] = "@"
	}
}

func partOne(data *[][]string, cmds []helpers.Pos) int {
	score := 0
	for _, v := range cmds {
		executeMove(data, v)
	}

	for i, row := range *data {
		for j, val := range row {
			if val == "O" {
				score += i*100 + j
			}
		}
	}

	return score
}

func parseCmds(cmds []string) []helpers.Pos {
	out := make([]helpers.Pos, 0)
	for _, v := range cmds {
		if v == "\n" {
			continue
		}
		var val helpers.Pos
		switch v {
		case "<":
			val = helpers.Pos{Row: 0, Col: -1}
		case "^":
			val = helpers.Pos{Row: -1, Col: 0}
		case ">":
			val = helpers.Pos{Row: 0, Col: 1}
		case "v":
			val = helpers.Pos{Row: 1, Col: 0}
		default:
			panic("invalid move")
		}
		out = append(out, val)
	}
	return out
}

func main() {
	path := os.Args[1]
	asString, _ := os.ReadFile(path)

	asSplit := strings.Split(string(asString), "\n\n")
	map_ := asSplit[0]
	commands := asSplit[1]
	splitCommands := strings.Split(commands, "")
	asChars := helpers.StringsToChars(helpers.BufferToLines(map_))
	parsedCmds := parseCmds(splitCommands)
	fmt.Println(partOne(&asChars, parsedCmds))
}
