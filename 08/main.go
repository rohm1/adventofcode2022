package main

import (
	"aoc2022/utils"
	"fmt"
)

func isVisible(lines []string, x int, y int) bool {
	if x == 0 || x == len(lines[0])-1 || y == 0 || y == len(lines)-1 {
		return true
	}

	// from top
	visible := true
	for yy := 0; yy < y; yy++ {
		if lines[yy][x] >= lines[y][x] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	// from bottom
	visible = true
	for yy := len(lines) - 1; yy > y; yy-- {
		if lines[yy][x] >= lines[y][x] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	// from left
	visible = true
	for xx := 0; xx < x; xx++ {
		if lines[y][xx] >= lines[y][x] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	// from right
	visible = true
	for xx := len(lines[0]) - 1; xx > x; xx-- {
		if lines[y][xx] >= lines[y][x] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	return false
}

func calcScenicScore(lines []string, x int, y int) int {
	score := 1

	// to top
	yLen := y
	for yy := y - 1; yy >= 0; yy-- {
		if lines[yy][x] >= lines[y][x] {
			yLen = y - yy
			break
		}
	}
	score *= yLen

	// to bottom
	yLen = len(lines) - 1 - y
	for yy := y + 1; yy < len(lines); yy++ {
		if lines[yy][x] >= lines[y][x] {
			yLen = yy - y
			break
		}
	}
	score *= yLen

	// to left
	xLen := x
	for xx := x - 1; xx >= 0; xx-- {
		if lines[y][xx] >= lines[y][x] {
			xLen = x - xx
			break
		}
	}
	score *= xLen

	// to right
	xLen = len(lines[0]) - 1 - x
	for xx := x + 1; xx < len(lines[0]); xx++ {
		if lines[y][xx] >= lines[y][x] {
			xLen = xx - x
			break
		}
	}
	score *= xLen

	return score
}

func main() {
	lines, err := utils.ReadFile("08", false)
	if err != nil {
		return
	}

	sum := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if isVisible(lines, x, y) {
				sum++
			}
		}
	}

	maxScenicScore := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			scenicScore := calcScenicScore(lines, x, y)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", maxScenicScore)
}
