package main

import (
	"aoc2022/utils"
	"fmt"
)

type pos struct {
	x int
	y int
}

type rock struct {
	width  int
	height int
	blocks []pos
}

type seenKey struct {
	rockId int
	jet    int
	tops   string
}

type seenValue struct {
	step   int
	height int
}

var GridHeight = 50_000
var GridSize = 7

func getNextRock(i int) rock {
	if i%5 == 0 {
		// ####
		return rock{
			width:  4,
			height: 1,
			blocks: []pos{
				{x: 0, y: 0},
				{x: 1, y: 0},
				{x: 2, y: 0},
				{x: 3, y: 0},
			},
		}
	}

	if i%5 == 1 {
		//  #
		// ###
		//  #
		return rock{
			width:  3,
			height: 3,
			blocks: []pos{
				{x: 1, y: 0},
				{x: 0, y: 1},
				{x: 1, y: 1},
				{x: 2, y: 1},
				{x: 1, y: 2},
			},
		}
	}

	if i%5 == 2 {
		//   #
		//   #
		// ###
		return rock{
			width:  3,
			height: 3,
			blocks: []pos{
				{x: 0, y: 0},
				{x: 1, y: 0},
				{x: 2, y: 0},
				{x: 2, y: 1},
				{x: 2, y: 2},
			},
		}
	}

	if i%5 == 3 {
		// #
		// #
		// #
		// #
		return rock{
			width:  1,
			height: 4,
			blocks: []pos{
				{x: 0, y: 0},
				{x: 0, y: 1},
				{x: 0, y: 2},
				{x: 0, y: 3},
			},
		}
	}

	// ##
	// ##
	return rock{
		width:  2,
		height: 2,
		blocks: []pos{
			{x: 0, y: 0},
			{x: 1, y: 0},
			{x: 0, y: 1},
			{x: 1, y: 1},
		},
	}
}

func isFreeSpace(top int, left int, grid [][]rune, rock rock) bool {
	if left == -1 || left+rock.width > GridSize || top == -1 {
		return false
	}

	for j := 0; j < len(rock.blocks); j++ {
		if grid[top+rock.blocks[j].y][left+rock.blocks[j].x] != ' ' {
			return false
		}
	}

	return true
}

func getChar(i int) rune {
	switch i % 5 {
	case 0:
		return 'X'
	case 1:
		return 'O'
	case 2:
		return '#'
	case 3:
		return '+'
	default:
		return '~'
	}
}

func getTops(grid [][]rune, top int) string {
	tops := ""
	for i := 0; i < GridSize; i++ {
		for j := top; j > -1; j-- {
			if grid[j][i] != ' ' {
				tops += fmt.Sprintf("%d_", top-j)
				break
			}
		}
	}
	return tops
}

func run(line string, steps int) int {
	jet := 0
	highest := 0

	grid := make([][]rune, GridHeight)
	for i := 0; i < GridHeight; i++ {
		grid[i] = make([]rune, GridSize)
		for j := 0; j < GridSize; j++ {
			grid[i][j] = ' '
		}
	}

	seen := map[seenKey]seenValue{}

	for step := 0; step < steps; step++ {
		top := highest + 3
		left := 2
		rock := getNextRock(step)

		for {
			/*
				for k := 200; k > -1; k-- {
					for kk := 0; kk < GridSize; kk++ {
						if grid[k][kk] == '*' {
							grid[k][kk] = ' '
						}
					}
				}
				for j := 0; j < len(rock.blocks); j++ {
					grid[top+rock.blocks[j].y][left+rock.blocks[j].x] = '*'
				}
				for k := 20; k > -1; k-- {
					fmt.Printf("|%s|\n", string(grid[k]))
				}
				fmt.Println("|=======|", string(line[jet]))
			//*/
			if line[jet] == '>' {
				if isFreeSpace(top, left+1, grid, rock) {
					left++
				}
			} else {
				if isFreeSpace(top, left-1, grid, rock) {
					left--
				}
			}

			jet++
			if jet == len(line) {
				jet = 0
			}

			top--

			if isFreeSpace(top, left, grid, rock) {
				continue
			}

			top++
			for j := 0; j < len(rock.blocks); j++ {
				grid[top+rock.blocks[j].y][left+rock.blocks[j].x] = getChar(step)
			}
			if highest < top+rock.height {
				highest = top + rock.height
			}
			/*
				for k := 200; k > -1; k-- {
					for kk := 0; kk < GridSize; kk++ {
						if grid[k][kk] == '*' {
							grid[k][kk] = ' '
						}
					}
				}
				for k := 20; k > -1; k-- {
					fmt.Printf("|%s|\n", string(grid[k]))
				}
				fmt.Println("|=======|", highest)
				//*/

			cacheKey := seenKey{rockId: step % 5, jet: jet, tops: getTops(grid, highest)}
			if s, found := seen[cacheKey]; step > 2022 && found {
				period := step - s.step

				m := (steps - step) / period
				ascend := highest - s.height
				highest += (m) * ascend

				restSteps := steps - m*period - step
				for _, ss := range seen {
					if ss.step == s.step+restSteps {
						return highest + (ss.height - s.height) - 1
					}
				}
			}
			seen[cacheKey] = seenValue{step: step, height: highest}
			break
		}
	}

	return highest
}

func main() {
	lines, err := utils.ReadFile("17", false)
	if err != nil {
		return
	}

	fmt.Println("Part 1: ", run(lines[0], 2022))
	fmt.Println("Part 2: ", run(lines[0], 1000000000000))
}
