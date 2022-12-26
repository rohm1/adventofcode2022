package main

import (
	"aoc2022/utils"
	"fmt"
)

type coord struct {
	x int
	y int
}

type elve struct {
	coord
	willMove      bool
	proposedCoord coord
}

type board struct {
	elves          []elve
	coords         map[coord]int
	proposedCoords map[coord]int
}

func (c1 coord) add(c2 coord) coord {
	return coord{x: c1.x + c2.x, y: c1.y + c2.y}
}

func (elve *elve) proposeMove(board board, turn int) {
	hasNeighbors := false
	for y := -1; y < 2 && !hasNeighbors; y++ {
		for x := -1; x < 2 && !hasNeighbors; x++ {
			if x == 0 && y == 0 {
				continue
			}

			if board.isOccupied(elve.coord.add(coord{x: x, y: y})) {
				hasNeighbors = true
			}
		}
	}

	if !hasNeighbors {
		return
	}

	// first 3: conditions, 4th: target
	order := [][]coord{
		// north
		[]coord{
			{-1, -1}, {0, -1}, {1, -1}, {0, -1},
		},
		// south
		[]coord{
			{-1, 1}, {0, 1}, {1, 1}, {0, 1},
		},
		// west
		[]coord{
			{-1, -1}, {-1, 0}, {-1, 1}, {-1, 0},
		},
		// east
		[]coord{
			{1, -1}, {1, 0}, {1, 1}, {1, 0},
		},
	}

	for i := turn; i < turn+4; i++ {
		direction := order[i%4]
		if !board.isOccupied(elve.coord.add(direction[0])) &&
			!board.isOccupied(elve.coord.add(direction[1])) &&
			!board.isOccupied(elve.coord.add(direction[2])) {
			elve.proposedCoord = elve.coord.add(direction[3])
			elve.willMove = true
			return
		}
	}

	return
}

func (elve *elve) reset() {
	elve.willMove = false
}

func (board *board) addElve(elve elve) {
	board.elves = append(board.elves, elve)
	board.coords[elve.coord] = 1
}

func (board *board) isOccupied(coord coord) bool {
	_, found := board.coords[coord]
	return found
}

func (board *board) proposeMoves(turn int) {
	for i := 0; i < len(board.elves); i++ {
		elve := &board.elves[i]
		elve.proposeMove(*board, turn)
		if elve.willMove {
			if _, found := board.proposedCoords[elve.proposedCoord]; !found {
				board.proposedCoords[elve.proposedCoord] = 0
			}

			board.proposedCoords[elve.proposedCoord]++
		}
	}
}

func (board *board) move() bool {
	board.coords = map[coord]int{}
	elves := board.elves
	board.elves = make([]elve, 0)
	moved := false

	for i := 0; i < len(elves); i++ {
		elve := &elves[i]
		willMove := elve.willMove
		elve.reset()

		if !willMove {
			board.addElve(*elve)
			continue
		}

		cnt, _ := board.proposedCoords[elve.proposedCoord]
		if cnt != 1 {
			board.addElve(*elve)
			continue
		}

		moved = true
		elve.coord = elve.proposedCoord
		board.addElve(*elve)
	}

	board.proposedCoords = map[coord]int{}

	return moved
}

func (board *board) countEmptyCoords() int {
	minX := 1_000_000_000
	maxX := -1_000_000_000
	minY := 1_000_000_000
	maxY := -1_000_000_000

	for i := 0; i < len(board.elves); i++ {
		elve := board.elves[i]

		if elve.x < minX {
			minX = elve.x
		}

		if elve.x > maxX {
			maxX = elve.x
		}

		if elve.y < minY {
			minY = elve.y
		}

		if elve.y > maxY {
			maxY = elve.y
		}
	}

	count := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if !board.isOccupied(coord{x: x, y: y}) {
				count++
			}
		}
	}

	return count
}

func parseLines(lines []string) board {
	board := board{elves: make([]elve, 0), coords: map[coord]int{}, proposedCoords: map[coord]int{}}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '#' {
				board.addElve(elve{coord: coord{x: x, y: y}, proposedCoord: coord{x: x, y: y}})
			}
		}
	}
	return board
}

func run(lines []string) (int, int) {
	board := parseLines(lines)

	turn := 0
	part1 := 0
	moved := true

	for ; moved; turn++ {
		board.proposeMoves(turn)
		moved = board.move()

		if turn == 9 {
			part1 = board.countEmptyCoords()
		}
	}

	return part1, turn
}

func main() {
	lines, err := utils.ReadFile("23", false)
	if err != nil {
		return
	}

	part1, part2 := run(lines)
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
