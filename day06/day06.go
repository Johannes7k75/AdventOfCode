package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input_test.txt
var input string

type Guard struct {
	x int
	y int
	//    1
	// -2   2
	//   -1
	dir int
}

func main() {
	part1 := solve(input)
	fmt.Println("Result of part1", part1)

	part2 := solve2(input)
	fmt.Println("Result of part2", part2)
}

func solve(input string) int {
	// TODO: Fix that its of by one in input but input_test is correct

	res := 0

	parsed, guardPos := ParseInput(input)
	walkedFields := WalkPath(parsed, guardPos)

	res = len(walkedFields)

	return res
}

func WalkPath(field [][]bool, guard Guard) [][2]int {
	var visited [][2]int

	for {
		x, y := MovementByGuard(guard)

		if guard.x+x >= len(field) || guard.y+y >= len(field[0]) {
			break
		}

		if field[guard.x+x][guard.y+y] {
			guard = TurnRight(guard)
		} else {
			guard.x += x
			guard.y += y
			if !IsVisited(visited, [2]int{guard.x, guard.y}) {
				visited = append(visited, [2]int{guard.x, guard.y})
			}
		}
	}

	return visited
}

func IsVisited(visited [][2]int, pos [2]int) bool {
	for _, v := range visited {
		if v[0] == pos[0] && v[1] == pos[1] {
			return true
		}
	}
	return false
}

func IsVisitedWithDir(visited [][3]int, pos [3]int) bool {
	for _, v := range visited {
		if v[0] == pos[0] && v[1] == pos[1] && v[2] == pos[2] {
			return true
		}
	}
	return false
}

func TurnRight(guard Guard) Guard {
	newGuard := Guard{guard.x, guard.y, guard.dir}
	switch newGuard.dir {
	case 1:
		newGuard.dir = 2
	case -1:
		newGuard.dir = -2
	case 2:
		newGuard.dir = -1
	case -2:
		newGuard.dir = 1
	}

	return newGuard
}

func ParseInput(input string) ([][]bool, Guard) {
	splitInput := strings.Split(input, "\r\n")
	res := CreateMatrix(len(splitInput), len(splitInput[0]))

	guardPos := Guard{0, 0, 1}

	for x, v := range splitInput {
		for y := 0; y < len(v); y++ {
			res[x][y] = v[y] == '#'
			if v[y] == '^' {
				guardPos = Guard{x, y, 1}
			}
		}
	}

	return res, guardPos
}

func CreateMatrix(rows int, cols int) [][]bool {
	multiSlice := make([][]bool, rows)
	for i := range multiSlice {
		multiSlice[i] = make([]bool, cols)
	}
	return multiSlice
}

func solve2(input string) int {
	res := 0

	parsed, guardPos := ParseInput(input)
	rows, cols := len(parsed), len(parsed[0])

	walkedFields := WalkPath(parsed, guardPos)

	for _, pathField := range walkedFields {
		x := pathField[0]
		y := pathField[1]

		parsed[x][y] = true

		guard := Guard{guardPos.x, guardPos.y, guardPos.dir}

		visitedWithDir := make(map[[3]int]bool)

		for {
			moveX, moveY := MovementByGuard(guard)

			newX, newY := guard.x+moveX, guard.y+moveY
			if newX < 0 || newX >= rows || newY < 0 || newY >= cols {
				break
			}

			if parsed[newX][newY] {
				guard = TurnRight(guard)
			} else {
				key := [3]int{guard.x, guard.y, guard.dir}
				if visitedWithDir[key] {
					res++
					break
				}

				visitedWithDir[key] = true

				guard.x, guard.y = newX, newY
			}
		}

		parsed[x][y] = false
	}

	// for x := 0; x < rows; x++ {
	// 	for y := 0; y < cols; y++ {
	// 		if parsed[x][y] {
	// 			continue
	// 		}

	// 		parsed[x][y] = true

	// 		guard := Guard{guardPos.x, guardPos.y, guardPos.dir}

	// 		visitedWithDir := make(map[[3]int]bool)

	// 		for {
	// 			moveX, moveY := MovementByGuard(guard)

	// 			newX, newY := guard.x+moveX, guard.y+moveY
	// 			if newX < 0 || newX >= rows || newY < 0 || newY >= cols {
	// 				break
	// 			}

	// 			if parsed[newX][newY] {
	// 				guard = TurnRight(guard)
	// 			} else {
	// 				key := [3]int{guard.x, guard.y, guard.dir}
	// 				if visitedWithDir[key] {
	// 					res++
	// 					break
	// 				}

	// 				visitedWithDir[key] = true

	// 				guard.x, guard.y = newX, newY
	// 			}
	// 		}

	// 		parsed[x][y] = false
	// 	}
	// }

	return res
}

func PrintMap(m [][]bool) {
	for x := range m {
		for y := range m[x] {
			if m[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func MovementByGuard(guard Guard) (int, int) {
	x := 0
	y := 0

	switch guard.dir {
	case 1:
		x = -1
	case -1:
		x = 1
	case 2:
		y = 1
	case -2:
		y = -1
	}

	return x, y
}
