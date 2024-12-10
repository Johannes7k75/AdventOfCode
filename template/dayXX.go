package main

import (
	_ "embed"
	"fmt"
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
	return res
}

func solve2(input string) int {
	res := 0
	return res
}
