package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

type pos struct {
	left int
	top  int
}

type nextCubeFace struct {
	coord     pos
	dirChange int
}

var Right = 0
var Down = 1
var Left = 2
var Up = 3

var Empty = uint8(' ')
var Open = uint8('.')
var Wall = uint8('#')

var CubeSize = 50

// Cube sides:
//
//	16
//	4
// 35
// 2

var Face1 = pos{left: 1, top: 0}
var Face2 = pos{left: 0, top: 3}
var Face3 = pos{left: 0, top: 2}
var Face4 = pos{left: 1, top: 1}
var Face5 = pos{left: 1, top: 2}
var Face6 = pos{left: 2, top: 0}

var Cube = map[pos]map[int]nextCubeFace{
	Face1: map[int]nextCubeFace{
		Right: nextCubeFace{coord: Face6, dirChange: 0},
		Down:  nextCubeFace{coord: Face4, dirChange: 0},
		Left:  nextCubeFace{coord: Face3, dirChange: 2},
		Up:    nextCubeFace{coord: Face2, dirChange: 1},
	},
	Face6: map[int]nextCubeFace{
		Right: nextCubeFace{coord: Face5, dirChange: 2},
		Down:  nextCubeFace{coord: Face4, dirChange: 1},
		Left:  nextCubeFace{coord: Face1, dirChange: 0},
		Up:    nextCubeFace{coord: Face2, dirChange: 0},
	},
	Face4: map[int]nextCubeFace{
		Right: nextCubeFace{coord: Face6, dirChange: 3},
		Down:  nextCubeFace{coord: Face5, dirChange: 0},
		Left:  nextCubeFace{coord: Face3, dirChange: 3},
		Up:    nextCubeFace{coord: Face1, dirChange: 0},
	},
	Face3: map[int]nextCubeFace{
		Right: nextCubeFace{coord: Face5, dirChange: 0},
		Down:  nextCubeFace{coord: Face2, dirChange: 0},
		Left:  nextCubeFace{coord: Face1, dirChange: 2},
		Up:    nextCubeFace{coord: Face4, dirChange: 1},
	},
	Face5: map[int]nextCubeFace{
		Right: nextCubeFace{coord: Face6, dirChange: 2},
		Down:  nextCubeFace{coord: Face2, dirChange: 1},
		Left:  nextCubeFace{coord: Face3, dirChange: 0},
		Up:    nextCubeFace{coord: Face4, dirChange: 0},
	},
	Face2: map[int]nextCubeFace{
		Right: nextCubeFace{coord: Face5, dirChange: 3},
		Down:  nextCubeFace{coord: Face6, dirChange: 0},
		Left:  nextCubeFace{coord: Face1, dirChange: 3},
		Up:    nextCubeFace{coord: Face3, dirChange: 0},
	},
}

func parseMoves(instructions string, instructionsIndex int) (int, int) {
	originalIndex := instructionsIndex
	for {
		instructionsIndex++

		if instructionsIndex == len(instructions) || instructions[instructionsIndex] < '0' || instructions[instructionsIndex] > '9' {
			break
		}
	}

	moves, _ := strconv.Atoi(instructions[originalIndex:instructionsIndex])
	return moves, instructionsIndex
}

func parseDirection(instructions string, instructionsIndex int, currentDirection int) (int, int) {
	if instructions[instructionsIndex] == 'R' {
		return (currentDirection + 1) % 4, instructionsIndex + 1
	} else if instructions[instructionsIndex] == 'L' {
		currentDirection--
		if currentDirection == -1 {
			currentDirection = 3
		}
		return currentDirection, instructionsIndex + 1
	}

	panic(fmt.Sprintf("unkown direction %d", instructions[instructionsIndex]))
}

func determineNextPosPart1(board []string, direction int, currentPos pos) pos {
	nextPosition := currentPos

	if direction == Right {
		for {
			nextPosition.left++
			if nextPosition.left == len(board[nextPosition.top]) {
				nextPosition.left = 0
			}

			if board[nextPosition.top][nextPosition.left] != Empty {
				return nextPosition
			}
		}
	}

	if direction == Left {
		for {
			nextPosition.left--
			if nextPosition.left == -1 {
				nextPosition.left = len(board[nextPosition.top]) - 1
			}

			if board[nextPosition.top][nextPosition.left] != Empty {
				return nextPosition
			}
		}
	}

	if direction == Down {
		for {
			nextPosition.top++
			if nextPosition.top == len(board) {
				nextPosition.top = 0
			}

			if nextPosition.left >= len(board[nextPosition.top]) {
				continue
			}

			if board[nextPosition.top][nextPosition.left] != Empty {
				return nextPosition
			}
		}
	}

	if direction == Up {
		for {
			nextPosition.top--
			if nextPosition.top == -1 {
				nextPosition.top = len(board) - 1
			}

			if nextPosition.left >= len(board[nextPosition.top]) {
				continue
			}

			if board[nextPosition.top][nextPosition.left] != Empty {
				return nextPosition
			}
		}
	}

	panic(fmt.Sprintf("unkown direction %d", direction))
}

func determineNextPos2(direction int, currentPos pos) (pos, int) {
	// first translate position to face relative, then do the move and adapt if needed
	nextPosition := pos{
		left: currentPos.left % CubeSize,
		top:  currentPos.top % CubeSize,
	}

	if direction == Right {
		nextPosition.left++
	} else if direction == Left {
		nextPosition.left--
	} else if direction == Down {
		nextPosition.top++
	} else if direction == Up {
		nextPosition.top--
	} else {
		panic(fmt.Sprintf("unkown direction %d", direction))
	}

	// stayed on the same face of the cube
	if nextPosition.left < CubeSize && nextPosition.top < CubeSize && nextPosition.left >= 0 && nextPosition.top >= 0 {
		// adapt relative position to grid
		nextPosition.left = (currentPos.left/CubeSize)*CubeSize + nextPosition.left
		nextPosition.top = (currentPos.top/CubeSize)*CubeSize + nextPosition.top

		return nextPosition, direction
	}

	currentFaceMappings, _ := Cube[pos{left: currentPos.left / CubeSize, top: currentPos.top / CubeSize}]
	nextFaceDefinition := currentFaceMappings[direction]

	// new direction after face change
	nextDirection := (direction + nextFaceDefinition.dirChange) % 4

	// save border that switches
	var s int
	if direction == Right {
		s = nextPosition.top
	} else if direction == Down {
		s = CubeSize - nextPosition.left - 1
	} else if direction == Left {
		s = CubeSize - nextPosition.top - 1
	} else {
		s = nextPosition.left
	}

	// adapt position to new face in 50 (CubeSize) mod
	if nextDirection == Right {
		nextPosition = pos{left: 0, top: s}
	} else if nextDirection == Down {
		nextPosition = pos{left: CubeSize - s - 1, top: 0}
	} else if nextDirection == Left {
		nextPosition = pos{left: CubeSize - 1, top: CubeSize - s - 1}
	} else {
		nextPosition = pos{left: s, top: CubeSize - 1}
	}

	// adapt relative position to grid
	nextPosition.left = nextFaceDefinition.coord.left*CubeSize + nextPosition.left
	nextPosition.top = nextFaceDefinition.coord.top*CubeSize + nextPosition.top

	return nextPosition, nextDirection
}

func run(board []string, instructions string, part int) int {
	direction := Right
	position := pos{top: 0, left: 0}
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == Open {
			position.left = i
			break
		}
	}

	for instructionsIndex := 0; instructionsIndex != len(instructions); {
		if instructions[instructionsIndex] < '0' || instructions[instructionsIndex] > '9' {
			direction, instructionsIndex = parseDirection(instructions, instructionsIndex, direction)
		}

		moves, nid := parseMoves(instructions, instructionsIndex)
		instructionsIndex = nid

		for i := 0; i < moves; i++ {
			var nextPosition pos
			nextDirection := direction
			if part == 1 {
				nextPosition = determineNextPosPart1(board, direction, position)
			} else {
				nextPosition, nextDirection = determineNextPos2(direction, position)
			}

			if board[nextPosition.top][nextPosition.left] == Wall {
				break
			}

			position = nextPosition
			direction = nextDirection
		}
	}

	return 1000*(position.top+1) + 4*(position.left+1) + direction
}

func main() {
	lines, err := utils.ReadFile("22", false)
	if err != nil {
		return
	}

	fmt.Println("Part 1: ", run(lines[0:len(lines)-2], lines[len(lines)-1], 1))
	fmt.Println("Part 2: ", run(lines[0:len(lines)-2], lines[len(lines)-1], 2))
}
