package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type cube struct {
	z int
	y int
	x int
}

func newCube(z int, y int, x int) cube {
	return cube{z: z, y: y, x: x}
}

func (c cube) getNeighbors() []cube {
	return []cube{
		newCube(c.z-1, c.y, c.x),
		newCube(c.z+1, c.y, c.x),
		newCube(c.z, c.y+1, c.x),
		newCube(c.z, c.y-1, c.x),
		newCube(c.z, c.y, c.x-1),
		newCube(c.z, c.y, c.x+1),
	}
}

func parseLines(lines []string) map[cube]cube {
	grid := map[cube]cube{}
	for i := 0; i < len(lines); i++ {
		coords := strings.Split(lines[i], ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		c := newCube(z, y, x)
		grid[c] = c
	}

	return grid
}

func countExposedSides(grid map[cube]cube) int {
	exposedSides := 0

	for _, c := range grid {
		ns := c.getNeighbors()
		for i := 0; i < len(ns); i++ {
			if _, f := grid[ns[i]]; !f {
				exposedSides++
			}
		}
	}

	return exposedSides
}

func countExteriorExposedSides(grid map[cube]cube) int {
	// determine a cube that encompasses all the cubes, gather all the cubes that are between
	// the structure and the big cube, then count the adjacent sides

	min := 1_000_000_000
	max := -1_000_000_000
	for _, c := range grid {
		min = int(math.Min(float64(min), math.Min(float64(c.z), math.Min(float64(c.y), float64(c.x)))))
		max = int(math.Max(float64(max), math.Max(float64(c.z), math.Max(float64(c.y), float64(c.x)))))
	}
	min--
	max++

	exteriors := map[cube]int{}
	queue := []cube{newCube(min, min, min)}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if _, f := exteriors[c]; f {
			continue
		}

		exteriors[c] = 1

		ns := c.getNeighbors()
		for i := 0; i < len(ns); i++ {
			n := ns[i]
			if _, f := grid[n]; f {
				continue
			}

			if n.z < min || n.y < min || n.x < min || n.z > max || n.y > max || n.x > max {
				continue
			}

			queue = append(queue, n)
		}
	}

	exteriorExposedSides := 0
	for c, _ := range exteriors {
		ns := c.getNeighbors()
		for i := 0; i < len(ns); i++ {
			if _, f := grid[ns[i]]; f {
				exteriorExposedSides++
			}
		}
	}

	return exteriorExposedSides
}

func main() {
	lines, err := utils.ReadFile("18", false)
	if err != nil {
		return
	}

	grid := parseLines(lines)
	fmt.Println("Part 1: ", countExposedSides(grid))
	fmt.Println("Part 2: ", countExteriorExposedSides(grid))
}
