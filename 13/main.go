package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

type data struct {
	isList bool
	list   []data
	number int
}

type pair struct {
	first  data
	second data
}

func parseAt(line string, i int) (data, int) {
	d := data{isList: true, list: make([]data, 0)}

	for ; i < len(line); i++ {
		if line[i] == ',' {
			i++
		}

		if line[i] == ']' {
			return d, i
		}

		if line[i] == '[' {
			cd, j := parseAt(line, i+1)
			d.list = append(d.list, cd)
			i = j
			continue
		}

		k := 1
		for ; line[i+k] >= 48 && line[i+k] <= 57; k++ {
		}

		number, _ := strconv.Atoi(line[i : i+k])
		d.list = append(d.list, data{isList: false, number: number})
		i += k - 1
	}

	return d, i
}

func parseLine(line string) data {
	d, _ := parseAt(line, 1)
	return d
}

func areDataOrdered(d1 data, d2 data, shortSecondListIsOk bool) bool {
	if !d1.isList && !d2.isList {
		return d1.number <= d2.number
	}

	if d1.isList && d2.isList {
		for i := 0; i < len(d1.list); i++ {
			if i >= len(d2.list) {
				return shortSecondListIsOk
			}

			if !areDataOrdered(d1.list[i], d2.list[i], false) {
				return false
			}
		}

		return true
	}

	d1L := d1
	if !d1.isList {
		d1L = data{isList: true, list: make([]data, 1)}
		d1L.list[0] = data{isList: false, number: d1.number}
	}

	d2L := d2
	if !d2.isList {
		d2L = data{isList: true, list: make([]data, 1)}
		d2L.list[0] = data{isList: false, number: d2.number}
	}

	return areDataOrdered(d1L, d2L, true)
}

func countOrderedPairs(pairs []pair) int {
	o := 0
	for i := 0; i < len(pairs); i++ {
		if areDataOrdered(pairs[i].first, pairs[i].second, false) {
			o += i + 1
		}
	}

	return o
}

func main() {
	lines, err := utils.ReadFile("13", true)
	if err != nil {
		return
	}

	pairs := make([]pair, len(lines)/3)
	for i := 0; i < len(lines)/3; i++ {
		pairs[i] = pair{first: parseLine(lines[3*i]), second: parseLine(lines[3*i+1])}
	}

	// 399 too low
	fmt.Println("Part 1: ", countOrderedPairs(pairs))
}
