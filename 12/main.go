package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"sort"
)

type point struct {
	x         int
	y         int
	elevation int
}

type routePoint struct {
	p        point
	distance int
}

type filterPoint func(point) bool

func createMap(lines []string) ([][]point, point, point) {
	heightmap := make([][]point, len(lines))
	start := point{elevation: 0}
	end := point{elevation: 25}
	for i := 0; i < len(lines); i++ {
		heightmap[i] = make([]point, len(lines[i]))

		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'S' {
				start.x = j
				start.y = i
				heightmap[i][j] = start
			} else if lines[i][j] == 'E' {
				end.x = j
				end.y = i
				heightmap[i][j] = end
			} else {
				heightmap[i][j] = point{x: j, y: i, elevation: int(lines[i][j]) - 97}
			}
		}
	}

	return heightmap, start, end
}

func navigate(heightmap [][]point, end point) map[point]routePoint {
	routes := make([]routePoint, 0)
	visited := map[point]routePoint{}

	endPoint := routePoint{p: end, distance: 0}
	routes = append(routes, endPoint)
	visited[end] = endPoint
	for len(routes) != 0 {
		crtRoutePoint := routes[0]
		routes = routes[1:]

		jRange := []int{0, -1, 1}
		for i := 1; i > -2; i-- {
			for _, j := range jRange {
				if math.Abs(float64(i)) == 1 && math.Abs(float64(j)) == 1 ||
					i == 0 && j == 0 ||
					crtRoutePoint.p.y+i < 0 || crtRoutePoint.p.y+i >= len(heightmap) ||
					crtRoutePoint.p.x+j < 0 || crtRoutePoint.p.x+j >= len(heightmap[0]) {
					continue
				}

				nextPoint := heightmap[crtRoutePoint.p.y+i][crtRoutePoint.p.x+j]
				if crtRoutePoint.p.elevation-nextPoint.elevation > 1 {
					continue
				}

				if _, found := visited[nextPoint]; found {
					continue
				}

				nextRoutePoint := routePoint{p: nextPoint, distance: crtRoutePoint.distance + 1}
				routes = append(routes, nextRoutePoint)
				visited[nextPoint] = nextRoutePoint
			}
		}
	}

	return visited
}

func filterRoutes(routes map[point]routePoint, filter filterPoint) []routePoint {
	routesFiltered := make([]routePoint, 0)
	for p, routePoint := range routes {
		if filter(p) {
			routesFiltered = append(routesFiltered, routePoint)
		}
	}

	return routesFiltered
}

func run(part int, routes map[point]routePoint, filter filterPoint) {
	routesFiltered := filterRoutes(routes, filter)
	sort.Slice(routesFiltered, func(i, j int) bool { return routesFiltered[i].distance < routesFiltered[j].distance })
	fmt.Printf("Part %d: %d\n", part, routesFiltered[0].distance)
}

func main() {
	lines, err := utils.ReadFile("12", false)
	if err != nil {
		return
	}

	heightmap, start, end := createMap(lines)
	routes := navigate(heightmap, end)

	run(1, routes, func(p point) bool { return p == start })
	run(2, routes, func(p point) bool { return p.elevation == 0 })
}
