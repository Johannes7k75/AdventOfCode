package main

import (
	_ "embed"
	"fmt"
	"math"
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

	var stringsToCheck = make([]string, 0)
	splitInput := strings.Split(input, "\r\n")

	rows := len(splitInput)
	cols := len(splitInput[0])

	verticalStrings := CreateMatrix(cols, rows)
	leftDiagonalChars := CreateMatrix(cols+rows, rows+cols)
	rightDiagonalChars := CreateMatrix(cols+rows, rows+cols)

	for i, v := range splitInput {

		splitChars := strings.Split(v, "")

		stringsToCheck = append(stringsToCheck, v)

		for j, char := range splitChars {
			verticalStrings[j][i] = char

			rightDiagonalIndex := j - i
			if rightDiagonalIndex < 0 {
				rightDiagonalIndex = len(v) + int(math.Abs(float64(rightDiagonalIndex)))
			}
			rightDiagonalChars[rightDiagonalIndex][i] = char

			leftDiagonalIndex := j + i
			leftDiagonalChars[leftDiagonalIndex][i] = char
		}
	}

	stringToCount := "xmas"
	countHor := CountString(stringsToCheck, stringToCount)
	countVer := CountString(JoinString(verticalStrings), stringToCount)
	countLeftDiagonal := CountString(JoinString(leftDiagonalChars), stringToCount)
	countRightDiagonal := CountString(JoinString(rightDiagonalChars), stringToCount)
	res += countHor
	res += countVer
	res += countLeftDiagonal
	res += countRightDiagonal

	return res
}

func CountString(array []string, str string) int {
	count := 0
	for _, v := range array {
		count += strings.Count(strings.ToUpper(v), strings.ToUpper(str))
		count += strings.Count(Reverse(strings.ToUpper(v)), strings.ToUpper(str))
	}
	return count
}

func JoinString(array [][]string) []string {
	out := make([]string, len(array[0]))
	for _, v := range array {
		out = append(out, strings.Join(v, ""))
	}
	return out
}

func CreateMatrix(rows int, cols int) [][]string {
	multiSlice := make([][]string, rows)
	for i := range multiSlice {
		multiSlice[i] = make([]string, cols)
	}
	return multiSlice
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func solve2(input string) int {
	res := 0
	splitInput := strings.Split(input, "\r\n")
	rows := len(splitInput)
	cols := len(splitInput[0])

	charMatrix := CreateMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			charMatrix[i][j] = strings.Split(splitInput[i], "")[j]
		}
	}

	for i, v := range charMatrix {
		for j, char := range v {
			if strings.ToLower(char) != "a" {
				continue
			}

			if i < 1 || i >= len(charMatrix)-1 {
				continue
			}
			if j < 1 || j >= len(v)-1 {
				continue
			}

			ulChar := strings.ToLower(charMatrix[i-1][j-1])
			urChar := strings.ToLower(charMatrix[i-1][j+1])
			blChar := strings.ToLower(charMatrix[i+1][j-1])
			brChar := strings.ToLower(charMatrix[i+1][j+1])

			if ((ulChar == "m" && brChar == "s") || (ulChar == "s" && brChar == "m")) && ((blChar == "m" && urChar == "s") || (blChar == "s" && urChar == "m")) {
				res += 1
			}

		}
	}

	return res
}

//   j 0 1 2 3
// i
// 0   x x x x
// 1   x x x x
// 2   x x x x
// 3   x x x x
// right Diagonal Index j-i
// left Diagonal Index j+i
