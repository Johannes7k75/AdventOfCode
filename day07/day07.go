package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Equation struct {
	result  int
	numbers []int
}

func main() {
	part1 := solve(input)
	fmt.Println("Result of part1", part1)

	part2 := solve2(input)
	fmt.Println("Result of part2", part2)
}

func solve(input string) int {
	res := 0

	parsed := ParseInput(input)

	for _, v := range parsed {

		if UseOperator(v.numbers, 0, v.result) {
			res += v.result
		}
	}

	return res
}

func UseOperator(numbers []int, result int, expected int) bool {
	if len(numbers) == 0 {
		return result == expected
	}
	if result > expected {
		return false
	}

	mul := UseOperator(numbers[1:], result*numbers[0], expected)
	sum := UseOperator(numbers[1:], result+numbers[0], expected)

	return mul || sum
}

func ParseInput(input string) []Equation {
	var equations []Equation

	split := strings.Split(input, "\r\n")
	for _, v := range split {
		numbers := make([]int, 0)
		splitNumbers := strings.Split(v, " ")
		result, err := strconv.Atoi(strings.TrimRight(splitNumbers[0], ":"))
		if err != nil {
			fmt.Println("Couldnt parse number")
			continue
		}

		for _, num := range splitNumbers[1:] {
			number, err := strconv.Atoi(strings.TrimRight(num, ":"))
			if err != nil {
				fmt.Println("Couldnt parse number")
				continue
			}
			numbers = append(numbers, number)
		}

		equations = append(equations, Equation{result, numbers})
	}

	return equations
}

func solve2(input string) int {
	res := 0

	parsed := ParseInput(input)

	for _, v := range parsed {

		if UseOperator2(v.numbers, 0, v.result) {
			res += v.result
		}
	}

	return res
}

func UseOperator2(numbers []int, result int, expected int) bool {
	if len(numbers) == 0 {
		return result == expected
	}
	if result > expected {
		return false
	}

	mul := UseOperator2(numbers[1:], result*numbers[0], expected)
	sum := UseOperator2(numbers[1:], result+numbers[0], expected)
	concat := UseOperator2(numbers[1:], ConcatIntegers(result, numbers[0]), expected)

	return mul || sum || concat
}

func ConcatIntegers(a, b int) int {
	temp := b
	multiplier := 1
	for temp > 0 {
		multiplier *= 10
		temp /= 10
	}

	return a*multiplier + b
}
