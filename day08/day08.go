package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1 := solve(input)
	fmt.Println("Result of part1", part1)

	part2 := solve2(input)
	fmt.Println("Result of part2", part2)
}

func solve(input string) int {
	res := 0

	antennaMap, width, height := ParseInput(input)

	var newPos []Pos

	for _, value := range antennaMap {
		for i, iv := range value {
			for j, jv := range value {
				if j == i {
					continue
				}

				min := Pos{x: Min(iv.x, jv.x), y: Min(iv.y, jv.y)}
				max := Pos{x: Max(iv.x, jv.x), y: Max(iv.y, jv.y)}

				offset := Pos{x: max.x - min.x, y: max.y - min.y}

				var minPos, maxPos Pos

				dx := iv.x - jv.x
				dy := iv.y - jv.y

				minPos = Pos{
					x: iv.x - Sign(dx)*offset.x,
					y: iv.y - Sign(dy)*offset.y,
				}

				maxPos = Pos{
					x: iv.x + Sign(dx)*offset.x,
					y: iv.y + Sign(dy)*offset.y,
				}

				if InBounds(maxPos.x, 0, width) && InBounds(maxPos.y, 0, height) && !HasPos(append(newPos, value...), maxPos) {
					res += 1
					newPos = append(newPos, maxPos)
				}
				if InBounds(minPos.x, 0, width) && InBounds(minPos.y, 0, height) && !HasPos(append(newPos, value...), minPos) {
					res += 1
					newPos = append(newPos, minPos)
				}

			}
		}
	}

	return res
}

func InBounds(x, min, max int) bool {
	return x >= min && x < max
}

func HasPos(posList []Pos, pos Pos) bool {
	for _, p := range posList {
		if p.x == pos.x && p.y == pos.y {
			return true
		}
	}
	return false
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Pos struct {
	x, y int
}

func ParseInput(input string) (map[string][]Pos, int, int) {
	res := make(map[string][]Pos, 0)
	width, height := 0, 0

	split := strings.Split(input, "\r\n")
	height = len(split)

	for x, row := range split {
		fields := strings.Split(row, "")
		width = len(fields)
		for y, field := range fields {
			if field != "." {
				channel := field
				res[channel] = append(res[channel], Pos{x, y})
			}
		}
	}

	return res, width, height
}

func solve2(input string) int {
	res := 0

	antennaMap, width, height := ParseInput(input)

	var newPos []Pos

	for _, value := range antennaMap {
		for i, iv := range value {
			for j, jv := range value {
				if j == i {
					continue
				}

				min := Pos{x: Min(iv.x, jv.x), y: Min(iv.y, jv.y)}
				max := Pos{x: Max(iv.x, jv.x), y: Max(iv.y, jv.y)}

				offset := Pos{x: max.x - min.x, y: max.y - min.y}

				dx := iv.x - jv.x
				dy := iv.y - jv.y

				newPos = AddNewAntinode(width, height, offset, iv, jv, Pos{x: dx, y: dy}, newPos, true)
				newPos = AddNewAntinode(width, height, offset, iv, jv, Pos{x: dx, y: dy}, newPos, false)
			}
		}
	}

	for range newPos {
		res += 1
	}

	return res
}

func Sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

func AddNewAntinode(width, height int, offset, last, cur, vec Pos, newPos []Pos, isMin bool) []Pos {
	var pos Pos

	if isMin {
		pos.x = cur.x - Sign(vec.x)*offset.x
		pos.y = cur.y - Sign(vec.y)*offset.y
	} else {
		pos.x = cur.x + Sign(vec.x)*offset.x
		pos.y = cur.y + Sign(vec.y)*offset.y
	}

	if !InBounds(pos.x, 0, width) || !InBounds(pos.y, 0, height) {
		return newPos
	}

	if !HasPos(newPos, pos) {
		newPos = append(newPos, pos)
	}

	return AddNewAntinode(width, height, offset, cur, pos, vec, newPos, isMin)
}
