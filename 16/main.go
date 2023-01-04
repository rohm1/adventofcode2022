package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type room struct {
	flowRate int
	leadsTo  []string
}

type visit struct {
	t                int
	room             string
	toOpen           []string
	pressureReleased int
}

func newRoom(flowRate int, leadsTo []string) room {
	return room{flowRate: flowRate, leadsTo: leadsTo}
}

func newVisit(t int, room string, toOpen []string, pressureReleased int) visit {
	return visit{t: t, room: room, toOpen: toOpen, pressureReleased: pressureReleased}
}

func (v visit) toString() string {
	toOpenCopy := make([]string, len(v.toOpen))
	copy(toOpenCopy, v.toOpen)
	sort.Strings(toOpenCopy)
	return fmt.Sprintf("%d-%d-%s-%s", v.t, v.pressureReleased, v.room, strings.Join(toOpenCopy, "_"))
}

func parseLines(lines []string) map[string]room {
	rooms := map[string]room{}
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")

		flowParts := strings.Split(parts[4], "=")
		flowRate, _ := strconv.Atoi(strings.ReplaceAll(flowParts[1], ";", ""))

		leadsTo := make([]string, 0)
		for j := 9; j < len(parts); j++ {
			leadsTo = append(leadsTo, strings.ReplaceAll(parts[j], ",", ""))
		}
		rooms[parts[1]] = newRoom(flowRate, leadsTo)
	}

	return rooms
}

func calcDistances(rooms map[string]room) map[string]map[string]int {
	distances := map[string]map[string]int{}
	for n, _ := range rooms {
		distances[n] = map[string]int{}
	}

	for n, r := range rooms {
		for nn, _ := range rooms {
			if utils.Includes(r.leadsTo, nn) {
				distances[n][nn] = 1
			} else {
				distances[n][nn] = 1_000_000_000
			}
		}
	}

	for n, _ := range rooms {
		for nn, _ := range rooms {
			for nnn, _ := range rooms {
				a := float64(distances[nn][nnn])
				b := float64(distances[nn][n])
				c := float64(distances[n][nnn])

				distances[nn][nnn] = int(math.Min(a, b+c))
			}
		}
	}

	return distances
}

func calToOpen(rooms map[string]room) []string {
	toOpen := make([]string, 0)
	for n, r := range rooms {
		if r.flowRate > 0 {
			toOpen = append(toOpen, n)
		}
	}

	return toOpen
}

func determineMaxReleasedPressure(rooms map[string]room, maxTime int, toOpen []string, distances map[string]map[string]int) int {
	visited := map[string]visit{}
	queue := make([]visit, 0)
	queue = append(queue, newVisit(0, "AA", toOpen, 0))

	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		if v.t >= maxTime {
			continue
		}

		if _, f := visited[v.toString()]; f {
			continue
		}

		visited[v.toString()] = v

		for i := 0; i < len(v.toOpen); i++ {
			tto := v.t + distances[v.room][v.toOpen[i]] + 1
			vv := newVisit(tto, v.toOpen[i], utils.Remove(v.toOpen, v.toOpen[i]), v.pressureReleased+(maxTime-tto)*rooms[v.toOpen[i]].flowRate)
			queue = append(queue, vv)
		}
	}

	max := 0
	for _, v := range visited {
		if max < v.pressureReleased {
			max = v.pressureReleased
		}
	}

	return max
}

// from https://github.com/encse/adventofcode/blob/master/2022/Day16/Solution.cs#L38
func generatePermutations(toOpen []string) [][][]string {
	permutations := make([][][]string, 0)
	maxMask := 1 << (len(toOpen) - 1)

	for mask := 0; mask < maxMask; mask++ {
		elephant := make([]string, 0)
		human := make([]string, 0)

		elephant = append(elephant, toOpen[0])

		for ivalve := 1; ivalve < len(toOpen); ivalve++ {
			if mask&(1<<ivalve) == 0 {
				human = append(human, toOpen[ivalve])
			} else {
				elephant = append(elephant, toOpen[ivalve])
			}
		}

		permutations = append(permutations, [][]string{human, elephant})
	}

	return permutations
}

func allocateRoomsAndDetermineMaxReleasedPressure(rooms map[string]room, maxTime int, toOpen []string, distances map[string]map[string]int) int {
	permutations := generatePermutations(toOpen)

	max := 0
	for i := 0; i < len(permutations); i++ {
		score := determineMaxReleasedPressure(rooms, maxTime, permutations[i][0], distances) + determineMaxReleasedPressure(rooms, maxTime, permutations[i][1], distances)
		if score > max {
			max = score
		}
	}

	return max
}

func main() {
	lines, err := utils.ReadFile("16", false)
	if err != nil {
		return
	}

	rooms := parseLines(lines)

	distances := calcDistances(rooms)
	toOpen := calToOpen(rooms)

	fmt.Println("Part 1: ", determineMaxReleasedPressure(rooms, 30, toOpen, distances))
	fmt.Println("Part 2: ", allocateRoomsAndDetermineMaxReleasedPressure(rooms, 26, toOpen, distances))
}
