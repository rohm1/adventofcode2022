package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadFile("04", false)
	if err != nil {
		return
	}

	sum1 := 0
	sum2 := 0
	for i := 0; i < len(lines); i++ {
		sections := strings.Split(lines[i], ",")

		firstElveParts := strings.Split(sections[0], "-")
		firstElveStart, _ := strconv.Atoi(firstElveParts[0])
		firstElveEnd, _ := strconv.Atoi(firstElveParts[1])

		secondElveParts := strings.Split(sections[1], "-")
		secondElveStart, _ := strconv.Atoi(secondElveParts[0])
		secondElveEnd, _ := strconv.Atoi(secondElveParts[1])

		if (firstElveStart >= secondElveStart && firstElveEnd <= secondElveEnd) || (secondElveStart >= firstElveStart && secondElveEnd <= firstElveEnd) {
			sum1++
		}

		if (firstElveStart <= secondElveStart && firstElveEnd >= secondElveStart) || (secondElveStart <= firstElveStart && secondElveEnd >= firstElveStart) {
			sum2++
		}
	}

	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}
