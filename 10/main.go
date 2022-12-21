package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func processCycle(cycle int, reg int) int {
	pixel := (cycle - 1) % 40
	if pixel == 0 {
		fmt.Println("")
	}

	if pixel >= reg-1 && pixel < reg+2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return cycle * reg
	}

	return 0
}

func main() {
	lines, err := utils.ReadFile("10", false)
	if err != nil {
		return
	}

	cycle := 0
	reg := 1

	signalStrengths := 0
	for i := 0; i < len(lines) && cycle < 241; i++ {
		cycle++
		signalStrengths += processCycle(cycle, reg)

		parts := strings.Split(lines[i], " ")
		if parts[0] == "addx" {
			cycle++
			signalStrengths += processCycle(cycle, reg)

			x, _ := strconv.Atoi(parts[1])
			reg += x
		}
	}

	fmt.Println("")
	fmt.Println("")

	fmt.Println("Part 1: ", signalStrengths)
}
