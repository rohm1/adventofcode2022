package main

import "math"

type pos struct {
	x int
	y int
}

func newPos(x int, y int) pos {
	return pos{x: x, y: y}
}

func (p1 pos) dist(p2 pos) int {
	return int(math.Abs(float64(p1.x-p2.x))) + int(math.Abs(float64(p1.y-p2.y)))
}
