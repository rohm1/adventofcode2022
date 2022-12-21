package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func runSimulation(lines []string, ropeLength int) int {
	knotPositions := make([]pos, ropeLength)
	for i := 0; i < ropeLength; i++ {
		knotPositions[i] = pos{x: 0, y: 0}
	}

	tVisitedPositions := map[string]int{"0x0": 1}
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")

		steps, _ := strconv.Atoi(parts[1])
		for j := 0; j < steps; j++ {
			// move head
			switch parts[0] {
			case "R":
				knotPositions[0].x++
				break
			case "L":
				knotPositions[0].x--
				break
			case "U":
				knotPositions[0].y++
				break
			case "D":
				knotPositions[0].y--
				break
			}

			// move knots
			for p := 1; p < ropeLength; p++ {
				hPos := &knotPositions[p-1]
				tPos := &knotPositions[p]

				if (tPos.x == hPos.x && tPos.y == hPos.y) ||
					(tPos.x == hPos.x-1 && tPos.y == hPos.y) ||
					(tPos.x == hPos.x-1 && tPos.y == hPos.y-1) ||
					(tPos.x == hPos.x && tPos.y == hPos.y-1) ||
					(tPos.x == hPos.x+1 && tPos.y == hPos.y-1) ||
					(tPos.x == hPos.x+1 && tPos.y == hPos.y) ||
					(tPos.x == hPos.x+1 && tPos.y == hPos.y+1) ||
					(tPos.x == hPos.x && tPos.y == hPos.y+1) ||
					(tPos.x == hPos.x-1 && tPos.y == hPos.y+1) {
					continue
				}

				if tPos.x != hPos.x && tPos.y != hPos.y {
					if math.Abs(float64(hPos.x-tPos.x)) == 2 && math.Abs(float64(hPos.y-tPos.y)) == 2 {
						tPos.x += (hPos.x - tPos.x) / 2
						tPos.y += (hPos.y - tPos.y) / 2
						continue
					}

					if math.Abs(float64(hPos.x-tPos.x)) == 1 {
						tPos.x = hPos.x
					} else {
						tPos.y = hPos.y
					}
				}

				if tPos.x == hPos.x {
					if tPos.y < hPos.y {
						tPos.y++
					} else {
						tPos.y--
					}
				} else {
					if tPos.x < hPos.x {
						tPos.x++
					} else {
						tPos.x--
					}
				}
			}

			/*
				output := [][]rune{
					[]rune("====================================="),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
					[]rune("                                     "),
				}

				yO := 1
				xO := 0
				for q := 0; q < ropeLength; q++ {
					output[knotPositions[q].y+yO][knotPositions[q].x+xO] = 'X'
				}

				fmt.Println(lines[i], i, j)
				for q := 0; q < len(output); q++ {
					fmt.Println(string(output[q]))
				}
			*/

			// mark tail position
			tVisitedPositions[fmt.Sprintf("%dx%d", knotPositions[ropeLength-1].x, knotPositions[ropeLength-1].y)] = 1
		}
	}

	return len(tVisitedPositions)
}

func main() {
	lines, err := utils.ReadFile("09", false)
	if err != nil {
		return
	}

	fmt.Println("Part 1: ", runSimulation(lines, 2))
	fmt.Println("Part 2: ", runSimulation(lines, 10))
}
