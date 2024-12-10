package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Pos struct {
	x, y int
}

var (
	up    = Pos{x: 1}
	left  = Pos{y: -1}
	down  = Pos{x: -1}
	right = Pos{y: 1}
)

func addPos(p1, p2 Pos) Pos {
	return Pos{x: p1.x + p2.x, y: p1.y + p2.y}
}

func getTrailPos(p1 Pos, trailMap [][]int) int {
	return trailMap[p1.x][p1.y]
}

func HasPos(pos Pos, list []Pos) bool {
	for _, v := range list {
		if v.x == pos.x && v.y == pos.y {
			return true
		}
	}
	return false
}

func PosInBounds(pos Pos, width, height int) bool {
	return pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height
}
func main() {
	part1 := solve(input)
	fmt.Println("Result of part1", part1)

	part2 := solve2(input)
	fmt.Println("Result of part2", part2)
}

func solve(input string) int {
	res := 0

	parsed, numMap := ParseInput(input)

	for _, v := range numMap[0] {
		goals := make([]Pos, 0)
		CheckTrails(Pos{-1, 0}, v, parsed, &goals, false)
		res += len(goals)
	}

	return res
}

func solve2(input string) int {
	res := 0

	parsed, numMap := ParseInput(input)

	for _, v := range numMap[0] {
		goals := make([]Pos, 0)
		CheckTrails(Pos{-1, 0}, v, parsed, &goals, true)
		res += len(goals)
	}

	return res
}

func CheckTrails(lastPos, curPos Pos, trailMap [][]int, goals *[]Pos, unique bool) bool {
	if !PosInBounds(curPos, len(trailMap), len(trailMap[0])) {
		return false
	}

	trailPos := getTrailPos(curPos, trailMap)
	if lastPos.x != -1 && getTrailPos(lastPos, trailMap)+1 != trailPos {
		return false
	}

	if trailPos == 9 {
		if unique || !HasPos(curPos, *goals) {
			*goals = append(*goals, curPos)
		}
		return true
	}

	checked := false
	for _, v := range []Pos{up, down, left, right} {
		check := CheckTrails(curPos, addPos(curPos, v), trailMap, goals, unique)
		checked = checked || check
	}

	return checked
}

func ParseInput(input string) ([][]int, map[int][]Pos) {
	split := strings.Split(input, "\r\n")

	res := make([][]int, 0)
	numMap := make(map[int][]Pos, 0)

	for i, row := range split {
		fields := strings.Split(row, "")
		rowNumbers := make([]int, 0)
		for j, field := range fields {
			number, _ := strconv.Atoi(field)
			rowNumbers = append(rowNumbers, number)
			numMap[number] = append(numMap[number], Pos{i, j})
		}
		res = append(res, rowNumbers)
	}

	return res, numMap
}
