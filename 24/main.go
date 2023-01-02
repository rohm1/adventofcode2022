package main

import (
	"aoc2022/utils"
	"fmt"
	"sort"
)

func goTo(states states, start pos, end pos, startTime int) int {
	visited := map[visit]int{}
	queue := make([]visit, 0)
	queue = append(queue, newVisit(start, startTime))

	for len(queue) > 0 {
		sort.Slice(queue, func(i int, j int) bool {
			return queue[i].t+queue[i].pos.dist(end) < queue[j].t+queue[j].pos.dist(end)
		})
		v := queue[0]
		queue = queue[1:]

		if v.pos == end {
			return v.t
		}

		nextPos := []pos{
			v.pos,
			newPos(v.x, v.y+1),
			newPos(v.x+1, v.y),
			newPos(v.x, v.y-1),
			newPos(v.x-1, v.y),
		}

		for i := 0; i < len(nextPos); i++ {
			if states.getAt(v.t+1).pos(nextPos[i]) == '.' {
				nv := v.createNext(nextPos[i])
				if _, f := visited[nv]; !f {
					visited[nv] = 1
					queue = append(queue, nv)
				}
			}
		}
	}

	panic("no way found")
}

func main() {
	lines, err := utils.ReadFile("24", false)
	if err != nil {
		return
	}

	firstState := make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		firstState[i] = []rune(lines[i])
	}

	states := newStates(map[int]state{0: newState(firstState, 0)})
	start := newPos(1, 0)
	end := newPos(len(lines[0])-2, len(lines)-1)

	part1Time := goTo(states, start, end, 0)
	fmt.Println("Part 1: ", part1Time)

	part2Time := goTo(states, end, start, part1Time)
	part2Time = goTo(states, start, end, part2Time)
	fmt.Println("Part 2: ", part2Time)
}
