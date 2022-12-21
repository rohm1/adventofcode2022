package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	count int
	from  int
	to    int
}

func parseStacks(lines []string) []string {
	stackCount := 0
	maxStackHeight := 0
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			maxStackHeight = i - 2
			break
		}

		line = strings.ReplaceAll(line, " ", "")
		crtLineStackCount := len(line) / 3
		if crtLineStackCount > stackCount {
			stackCount = crtLineStackCount
		}
	}

	stacks := make([]string, stackCount)
	for i := 0; i < stackCount; i++ {
		column := i*4 + 1
		for j := maxStackHeight; j >= 0; j-- {
			if len(lines[j]) < column || string(lines[j][column]) == " " {
				break
			}

			stacks[i] += string(lines[j][column])
		}
	}

	return stacks
}

func parseInstructions(lines []string) []instruction {
	instructionsCnt := 0
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "move ") {
			instructionsCnt++
		}
	}

	instructions := make([]instruction, instructionsCnt)
	instructionNr := 0
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		if len(parts) == 6 {
			count, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])
			instructions[instructionNr] = instruction{count: count, from: from - 1, to: to - 1}
			instructionNr++
		}
	}

	return instructions
}

func calcTops(stacks []string) string {
	tops := ""
	for i := 0; i < len(stacks); i++ {
		stack := stacks[i]
		tops += string(stack[len(stack)-1])
	}

	return tops
}

func part1(lines []string, instructions []instruction) string {
	stacks := parseStacks(lines)

	for i := 0; i < len(instructions); i++ {
		for j := 0; j < instructions[i].count; j++ {
			fromStack := stacks[instructions[i].from]
			crate := fromStack[len(fromStack)-1]
			stacks[instructions[i].from] = fromStack[0 : len(fromStack)-1]
			stacks[instructions[i].to] += string(crate)
		}
	}

	return calcTops(stacks)
}

func part2(lines []string, instructions []instruction) string {
	stacks := parseStacks(lines)

	for i := 0; i < len(instructions); i++ {
		fromStack := stacks[instructions[i].from]
		crates := fromStack[len(fromStack)-instructions[i].count:]
		stacks[instructions[i].from] = fromStack[0 : len(fromStack)-instructions[i].count]
		stacks[instructions[i].to] += crates
	}

	return calcTops(stacks)
}

func main() {
	lines, err := utils.ReadFile("05", false)
	if err != nil {
		return
	}

	instructions := parseInstructions(lines)

	fmt.Println("Part 1: ", part1(lines, instructions))
	fmt.Println("Part 2: ", part2(lines, instructions))
}
