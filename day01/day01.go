package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	distance := solve(input)
	fmt.Println("The Total Distance is ", distance)

	sinilarityScore := solve2(input)
	fmt.Println("The Sinilarity Score is ", sinilarityScore)
}

func SplitIntoCols(input string) ([]int, []int) {
	rows := strings.Split(input, "\r\n")

	var firstCol []int
	var secondCol []int

	for _, v := range rows {
		cols := strings.Split(v, "   ")
		firstNum, err1 := strconv.Atoi(cols[0])
		secondNum, err2 := strconv.Atoi(cols[1])
		if err1 != nil {
			fmt.Println("Num not parseable ", err1)
			continue
		}
		if err2 != nil {
			fmt.Println("Num not parseable ", err2)
			continue
		}

		firstCol = append(firstCol, firstNum)
		secondCol = append(secondCol, secondNum)
	}

	return firstCol, secondCol
}

func SortCol(col []int) {
	sort.Slice(col, func(i, j int) bool {
		return col[i] < col[j]
	})
}

func solve(input string) int {
	firstCol, secondCol := SplitIntoCols(input)

	SortCol(firstCol)
	SortCol(secondCol)

	var totalDistance int = 0
	for i, num := range firstCol {
		totalDistance += CalculateDistance(num, secondCol[i])
	}

	return totalDistance
}

func CalculateDistance(num1 int, num2 int) int {
	if num1 < num2 {
		return num2 - num1
	}
	return num1 - num2
}

func solve2(input string) int {
	firstCol, secondCol := SplitIntoCols(input)

	SortCol(firstCol)
	SortCol(secondCol)

	var totalSimilarityScore int = 0
	for _, num := range firstCol {
		count := 0

		for _, num2 := range secondCol {
			if num == num2 {
				count += 1
			}
		}
		totalSimilarityScore += (num * count)
	}

	return totalSimilarityScore
}
