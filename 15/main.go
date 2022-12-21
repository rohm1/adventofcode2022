package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var EMPTY = 0
var SENSOR = 1
var BEACON = 2
var NotPossible = 3

var TargetRow = 2_000_000
var TargetSize = 4_000_000

type input struct {
	sx int
	sy int
	bx int
	by int
	d  int
}

type rect struct {
	x int
	y int
	w int
	h int
}

func parseInput(lines []string) []input {
	inputs := make([]input, len(lines))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		line = strings.ReplaceAll(line, "=", " ")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ":", "")
		parts := strings.Split(line, " ")

		sx, _ := strconv.Atoi(parts[3])
		sy, _ := strconv.Atoi(parts[5])
		bx, _ := strconv.Atoi(parts[11])
		by, _ := strconv.Atoi(parts[13])

		inputs[i] = input{sx: sx, sy: sy, bx: bx, by: by, d: int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))}
	}

	return inputs
}

func createMap(inputs []input) map[int]map[int]int {
	caves := make(map[int]map[int]int, 0)
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]

		if _, found := caves[input.sy]; !found {
			caves[input.sy] = make(map[int]int, 0)
		}
		caves[input.sy][input.sx] = SENSOR

		if _, found := caves[input.by]; !found {
			caves[input.by] = make(map[int]int, 0)
		}
		caves[input.by][input.bx] = BEACON
	}

	return caves
}

func part1(inputs []input) int {
	caves := createMap(inputs)
	if _, found := caves[TargetRow]; !found {
		caves[TargetRow] = make(map[int]int, 0)
	}

	count := 0
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		if input.sy+input.d < TargetRow || input.sy-input.d > TargetRow {
			continue
		}

		dd := input.d - int(math.Abs(float64(input.sy-TargetRow)))
		for x := input.sx - dd; x < input.sx+dd+1; x++ {
			d1 := int(math.Abs(float64(input.sx-x))) + int(math.Abs(float64(input.sy-TargetRow)))

			if d1 <= input.d && caves[TargetRow][x] == EMPTY {
				count++
				caves[TargetRow][x] = NotPossible
			}
		}
	}

	return count
}

func scanRect(inputs []input, searchArea rect) *rect {
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]

		// fully included in a sensor range: not possible
		if int(math.Abs(float64(input.sx-searchArea.x))+math.Abs(float64(input.sy-searchArea.y))) <= input.d &&
			int(math.Abs(float64(input.sx-(searchArea.x+searchArea.w-1)))+math.Abs(float64(input.sy-(searchArea.y+searchArea.h-1)))) <= input.d &&
			int(math.Abs(float64(input.sx-(searchArea.x+searchArea.w-1)))+math.Abs(float64(input.sy-searchArea.y))) <= input.d &&
			int(math.Abs(float64(input.sx-searchArea.x))+math.Abs(float64(input.sy-(searchArea.y+searchArea.h-1)))) <= input.d {
			return nil
		}
	}

	// 1x1 rect, not covered: that's the one
	if searchArea.w == 1 && searchArea.h == 1 {
		return &searchArea
	}

	w0 := searchArea.w / 2
	w1 := searchArea.w - w0
	h0 := searchArea.h / 2
	h1 := searchArea.h - h0
	childs := []rect{
		{x: searchArea.x, y: searchArea.y, w: w0, h: h0},
		{x: searchArea.x + w0, y: searchArea.y, w: w1, h: h0},
		{x: searchArea.x, y: searchArea.y + h0, w: w0, h: h1},
		{x: searchArea.x + w0, y: searchArea.y + h0, w: w1, h: h1},
	}

	for i := 0; i < len(childs); i++ {
		if childs[i].w == 0 || childs[i].h == 0 {
			continue
		}

		rect := scanRect(inputs, childs[i])
		if rect != nil {
			return rect
		}
	}

	return nil
}

func part2(inputs []input, searchArea rect) int {
	c := scanRect(inputs, searchArea)
	return c.x*TargetSize + c.y
}

func main() {
	lines, err := utils.ReadFile("15", false)
	if err != nil {
		return
	}

	inputs := parseInput(lines)
	fmt.Println("Part 1: ", part1(inputs))
	fmt.Println("Part 2: ", part2(inputs, rect{x: 0, y: 0, w: TargetSize, h: TargetSize}))
}
