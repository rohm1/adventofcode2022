package main

import (
	"aoc2022/utils"
	"fmt"
)

func hasUniqueChars(chunk string) bool {
	charsLookup := make([]int, 130)
	for i := 0; i < len(chunk); i++ {
		if charsLookup[chunk[i]] != 0 {
			return false
		}

		charsLookup[chunk[i]]++
	}

	return true
}

func getStart(line string, markerLen int) int {
	for markerStart := markerLen; markerStart < len(line); markerStart++ {
		if hasUniqueChars(line[markerStart-markerLen : markerStart]) {
			return markerStart
		}
	}

	return 0
}

func main() {
	lines, err := utils.ReadFile("06", false)
	if err != nil {
		return
	}

	line := lines[0]

	fmt.Println("Part 1: ", getStart(line, 4))
	fmt.Println("Part 2: ", getStart(line, 14))
}
