package main

import (
	_ "embed"
	"fmt"
	"regexp"
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
	var validMult = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	mults := validMult.FindAllString(input, -1)
	res := 0

	for _, v := range mults {
		v = strings.Replace(v, "mul(", "", 1)
		v = strings.Replace(v, ")", "", 1)

		nums := strings.Split(v, ",")

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Could not parse num")
			continue
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("Could not parse num")
			continue
		}

		res += num1 * num2
	}

	return res
}

func solve2(input string) int {
	var validMult = regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(don't\(\))|(do\(\))`)

	mults := validMult.FindAllString(input, -1)
	res := 0
	fmt.Println(mults)

	shouldDo := true

	for _, v := range mults {
		if v == "do()" {
			shouldDo = true
			continue
		} else if v == "don't()" {
			shouldDo = false
			continue
		}

		if !shouldDo {
			continue
		}

		v = strings.Replace(v, "mul(", "", 1)
		v = strings.Replace(v, ")", "", 1)

		nums := strings.Split(v, ",")

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Could not parse num")
			continue
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("Could not parse num")
			continue
		}

		res += num1 * num2
	}

	return res
}
