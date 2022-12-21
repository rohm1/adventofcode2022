package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

var AIR = 0
var ROCK = 1
var REST = 2

var CaveSize = 10_000

func parseInput(lines []string) ([][]int, int) {
	caves := make([][]int, CaveSize)
	for i := 0; i < CaveSize; i++ {
		caves[i] = make([]int, CaveSize)
	}

	floor := 0

	for i := 0; i < len(lines); i++ {
		line := strings.ReplaceAll(lines[i], " ", "")
		parts := strings.Split(line, "->")
		for j := 0; j < len(parts)-1; j++ {
			coords1 := strings.Split(parts[j], ",")
			x1, _ := strconv.Atoi(coords1[0])
			y1, _ := strconv.Atoi(coords1[1])
			if y1 > floor {
				floor = y1
			}

			coords2 := strings.Split(parts[j+1], ",")
			x2, _ := strconv.Atoi(coords2[0])
			y2, _ := strconv.Atoi(coords2[1])
			if y2 > floor {
				floor = y2
			}

			if x1 != x2 {
				xStep := 1
				if x2 < x1 {
					xStep = -1
				}

				for xx := x1; ; {
					caves[y1][xx] = ROCK

					if xx == x2 {
						break
					}

					xx += xStep
				}
			} else {
				yStep := 1
				if y2 < y1 {
					yStep = -1
				}

				for yy := y1; ; {
					caves[yy][x1] = ROCK

					if yy == y2 {
						break
					}

					yy += yStep
				}
			}
		}
	}

	return caves, floor
}

func run(caves [][]int, floor int) int {
	units := 0
	for {
		x := 500
		y := 0
		for y < floor+2 {
			if caves[y+1][x] == AIR {
				y++
				continue
			}

			if caves[y+1][x-1] == AIR {
				y++
				x--
				continue
			}

			if caves[y+1][x+1] == AIR {
				y++
				x++
				continue
			}

			units++
			caves[y][x] = REST
			break
		}

		// part 1
		if y >= floor+2 {
			break
		}

		// part 2
		if x == 500 && y == 0 {
			break
		}
	}

	return units
}

func main() {
	lines, err := utils.ReadFile("14", false)
	if err != nil {
		return
	}

	fmt.Println("Part 1: ", run(parseInput(lines)))

	caves, floor := parseInput(lines)
	for x := 0; x < CaveSize; x++ {
		caves[floor+2][x] = ROCK
	}
	fmt.Println("Part 2: ", run(caves, floor))
}
