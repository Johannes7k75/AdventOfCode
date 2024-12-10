package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
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
	pageOrdering, pageUpdates := ParseInput(input)

	for _, updates := range pageUpdates {
		orderings := GetValidOrderings(pageOrdering, updates)

		if IsSorted(orderings, updates) {
			res += updates[(len(updates)-1)/2]
		}

	}

	return res
}

func IsSorted(pageOrdering [][2]int, pageUpdates []int) bool {
	for _, ordering := range pageOrdering {
		first := ordering[0]
		second := ordering[1]

		if slices.Index(pageUpdates, first) > slices.Index(pageUpdates, second) {
			return false
		}
	}
	return true
}

func GetValidOrderings(orderings [][2]int, updates []int) [][2]int {
	var order [][2]int
	for _, v := range orderings {
		if slices.Contains(updates, v[0]) && slices.Contains(updates, v[1]) {
			order = append(order, v)
		}
	}
	return order
}

func ParseInput(input string) ([][2]int, [][]int) {
	var pageOrderings [][2]int
	var pageUpdates [][]int

	split := strings.Split(input, "\r\n\r\n")
	splitPageOrderings := strings.Split(split[0], "\r\n")
	splitPageUpdates := strings.Split(split[1], "\r\n")

	pageOrder := make([]int, 2)
	for _, v := range splitPageOrderings {
		splitOrder := strings.Split(v, "|")
		firstNumber, _ := strconv.Atoi(splitOrder[0])
		secondNumber, _ := strconv.Atoi(splitOrder[1])

		pageOrder[0] = firstNumber
		pageOrder[1] = secondNumber

		pageOrderings = append(pageOrderings, [2]int(pageOrder))
	}

	for _, v := range splitPageUpdates {
		splitUpdates := strings.Split(v, ",")
		updatedPages := make([]int, len(splitUpdates))
		for i, num := range splitUpdates {
			number, _ := strconv.Atoi(num)
			updatedPages[i] = number
		}
		pageUpdates = append(pageUpdates, updatedPages)
	}

	return pageOrderings, pageUpdates
}

func solve2(input string) int {
	res := 0
	pageOrdering, pageUpdates := ParseInput(input)

	for _, updates := range pageUpdates {
		orderings := GetValidOrderings(pageOrdering, updates)

		if IsSorted(orderings, updates) {
			continue
		}

		for {
			if IsSorted(orderings, updates) {
				break
			}

			for _, o := range orderings {
				aIndex := slices.Index(updates, o[0])
				bIndex := slices.Index(updates, o[1])
				if aIndex > bIndex {
					updates[aIndex], updates[bIndex] = updates[bIndex], updates[aIndex]
				}
			}

		}
		res += updates[(len(updates)-1)/2]
	}

	return res
}
