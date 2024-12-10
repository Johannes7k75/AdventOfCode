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

	diskMap, size := ParseInput(input)
	mappedDisk := MapDisk(diskMap, size)
	sortedMappedDisk := SortMappedDisk(mappedDisk)

	res = Checksum(sortedMappedDisk)

	return res
}

func solve2(input string) int {
	res := 0

	diskMap, size := ParseInput(input)
	mappedDisk := MapDisk(diskMap, size)
	sortedMappedDisk := SortBlockMappedDisk(mappedDisk)

	res = Checksum(sortedMappedDisk)

	return res
}

func MapDisk(diskMap []int, size int) []int {
	mappedDisk := make([]int, size)

	index := 0
	for i, v := range diskMap {
		isFree := i%2 == 1

		for f := 0; f < v; f++ {
			if isFree {
				mappedDisk[index] = -1
			} else {
				mappedDisk[index] = i / 2
			}
			index++
		}
	}

	return mappedDisk
}

func SortMappedDisk(mappedDisk []int) []int {
	var cleanedMappedDisk []int

	for _, v := range mappedDisk {
		if v >= 0 {
			cleanedMappedDisk = append(cleanedMappedDisk, v)
		}
	}

	clonedMappedDisk := slices.Clone(mappedDisk)
	sorted := make([]int, len(cleanedMappedDisk))

	for i := 0; i < len(clonedMappedDisk); i++ {
		if i > len(clonedMappedDisk) || i >= len(sorted) {
			continue
		}

		if clonedMappedDisk[i] >= 0 {
			sorted[i] = clonedMappedDisk[i]
		} else if clonedMappedDisk[i] == -1 {
			sorted[i] = cleanedMappedDisk[len(cleanedMappedDisk)-1]
			cleanedMappedDisk = cleanedMappedDisk[:len(cleanedMappedDisk)-1]
			clonedMappedDisk = clonedMappedDisk[:len(clonedMappedDisk)-1]
		}
	}

	return sorted
}

func SortBlockMappedDisk(mappedDisk []int) []int {
	type block struct {
		value int
		len   int
		idx   int
	}

	var diskMap []block
	blockLen := 0
	startIdx := -1
	var blockVal int = -2

	for i, v := range mappedDisk {
		if blockVal == v {
			blockLen++
		}
		if (blockVal != v || i == (len(mappedDisk)-1)) && blockLen > 0 {
			diskMap = append(diskMap, block{value: mappedDisk[startIdx], len: blockLen, idx: startIdx})
			blockLen = 0
		}
		if blockVal != v && v != -1 {
			startIdx = i
			blockVal = v
			blockLen = 1
		}
	}

	findFreeSpace := func(validBlockSize int) int {
		idx := -1

		blockSize := 0
		for i, v := range mappedDisk {
			if blockSize >= validBlockSize {
				return idx
			}

			if v == -1 && blockSize == 0 {
				idx = i
			}

			if v != -1 {
				idx = -1
				blockSize = 0
			}

			if idx != -1 {
				blockSize++
			}
		}

		return -1
	}

	slices.Reverse(diskMap)
	for _, fileBlock := range diskMap {
		spaceIdx := findFreeSpace(fileBlock.len)
		if spaceIdx == -1 || spaceIdx > fileBlock.idx {
			continue
		}
		mappedDisk = slices.Delete(mappedDisk, fileBlock.idx, fileBlock.idx+fileBlock.len)
		mappedDisk = slices.Insert(mappedDisk, fileBlock.idx, createArr(-1, fileBlock.len)...)
		mappedDisk = slices.Delete(mappedDisk, spaceIdx, spaceIdx+fileBlock.len)
		mappedDisk = slices.Insert(mappedDisk, spaceIdx, createArr(fileBlock.value, fileBlock.len)...)
	}

	return mappedDisk
}

func createArr(val, size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = val
	}
	return arr
}

func ParseInput(input string) ([]int, int) {
	split := strings.Split(input, "")
	diskMap := make([]int, len(split))
	size := 0

	for i, v := range split {
		numnber, _ := strconv.Atoi(v)
		diskMap[i] = numnber
		size += numnber
	}

	return diskMap, size
}

func Checksum(fileMap []int) int {
	res := 0
	for i, v := range fileMap {
		if v == -1 {
			continue
		}
		res += i * v
	}
	return res
}
