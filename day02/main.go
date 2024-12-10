package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	safeReports := solve(input)
	fmt.Println("The amount of safe reports is ", safeReports)

	safeReportsDampener := solve2(input)
	fmt.Println("The amount of safe reports is with Dampener", safeReportsDampener)
}

func SplitIntoRows(input string) [][]int {
	splitRows := strings.Split(input, "\r\n")

	var rows [][]int = make([][]int, len(splitRows))
	for i, splitRow := range splitRows {
		numbers := strings.Split(splitRow, " ")
		rows[i] = make([]int, len(numbers))

		for j, v := range numbers {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Could not parse num")
				continue
			}

			rows[i][j] = num
		}
	}

	return rows
}

func solve(input string) int {
	rows := SplitIntoRows(input)

	var safeReports int = 0
	for _, row := range rows {

		safeReport := CheckIRowIsSafe(row)

		if safeReport {
			safeReports++
		}
	}

	return safeReports
}

func solve2(input string) int {
	rows := SplitIntoRows(input)

	var safeReports int = 0
	for _, row := range rows {

		if CheckIRowIsSafe(row) || CheckWithDampeners(row) {
			safeReports++
			continue
		}
	}

	return safeReports
}

func CheckIRowIsSafe(row []int) bool {
	var increasing bool
	var safeReport bool = true

	increasing = row[0] < row[1]
	for i, num := range row {
		if i == 0 {
			continue
		}

		difference := num - row[i-1]

		if difference == 0 {
			safeReport = false
			continue
		}

		if increasing && difference < 0 || difference > 3 {
			safeReport = false
			continue
		}

		if !increasing && difference > 0 || difference < -3 {
			safeReport = false
			continue
		}
	}

	return safeReport
}

func CheckWithDampeners(row []int) bool {
	for i := range row {
		slice := append([]int{}, row[:i]...)
		slice = append(slice, row[i+1:]...)

		if CheckIRowIsSafe(slice) {
			return true
		}
	}
	return false
}
