package main

import (
	"aoc2022/utils"
	"fmt"
)

func charToPriority(c byte) int {
	asciiCode := int(c)
	if asciiCode > 96 {
		return asciiCode - 96
	}

	return asciiCode - 38
}

func part1(lines []string) int {
	sum := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineLen := len(line)
		compartmentLen := lineLen / 2
		itemsInFirstCompartment := make([]int, 130)

		for c := 0; c < compartmentLen; c++ {
			itemsInFirstCompartment[line[c]] = 1
		}

		for c := compartmentLen; c < lineLen; c++ {
			if itemsInFirstCompartment[line[c]] == 1 {
				itemsInFirstCompartment[line[c]]++
				sum += charToPriority(line[c])
			}
		}
	}

	return sum
}

func calGroupBadge(groupSize int, elvesCharsMap [][]int) int {
	for char := 65; char < 130; char++ {
		if elvesCharsMap[0][char] == 1 {
			sameCnt := 1
			for elve := 1; elve < groupSize; elve++ {
				if elvesCharsMap[elve][char] == 1 {
					sameCnt++
				}
			}

			if sameCnt == groupSize {
				return charToPriority(byte(char))
			}
		}
	}

	return 0
}

func part2(lines []string) int {
	groupSize := 3
	sum := 0
	elvesCharsMap := make([][]int, groupSize)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		elveNrInGroup := i % groupSize
		if elveNrInGroup == 0 {
			if i != 0 {
				sum += calGroupBadge(groupSize, elvesCharsMap)
			}

			elvesCharsMap = make([][]int, groupSize)
			for j := range elvesCharsMap {
				elvesCharsMap[j] = make([]int, 130)
			}
		}

		for l := 0; l < len(line); l++ {
			elvesCharsMap[elveNrInGroup][line[l]] = 1
		}
	}

	sum += calGroupBadge(groupSize, elvesCharsMap)

	return sum
}

func main() {
	lines, err := utils.ReadFile("03", false)
	if err != nil {
		return
	}

	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}
