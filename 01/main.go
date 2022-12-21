package main

import (
	"aoc2022/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	lines, err := utils.ReadFile("01", true)
	if err != nil {
		return
	}

	var calsByElves []int
	cals := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if line == "" {
			calsByElves = append(calsByElves, cals)
			cals = 0
			continue
		}

		_cals, _ := strconv.Atoi(line)
		cals += _cals
	}

	arrSlice := calsByElves[:]
	sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))

	fmt.Println("Part 1: ", arrSlice[0])
	fmt.Println("Part 2: ", arrSlice[0]+arrSlice[1]+arrSlice[2])
}
