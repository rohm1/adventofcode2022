package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

type node struct {
	originalIndex int
	number        int
	prev          *node
	next          *node
}

func parseInput(lines []string) *node {
	var first *node
	var prev *node
	var crt *node
	for i := 0; i < len(lines); i++ {
		number, _ := strconv.Atoi(lines[i])
		prev = crt
		crt = &node{prev: prev, number: number, originalIndex: i}
		if i == 0 {
			first = crt
		} else {
			prev.next = crt
		}
	}
	crt.next = first
	first.prev = crt
	return first
}

func run(first *node, cnt int, runs int) int {
	for run := 0; run < runs; run++ {
		for i := 0; i < cnt; i++ {
			node := first
			for node.originalIndex != i {
				node = node.next
			}

			if node.number == 0 {
				continue
			}

			moves := node.number % (cnt - 1)
			if moves < 0 {
				moves = cnt + moves - 1
			}

			newPos := node
			for j := 0; j < moves; j++ {
				newPos = newPos.next
			}

			node.prev.next = node.next
			node.next.prev = node.prev

			node.prev = newPos
			node.next = newPos.next

			newPos.next = node
			node.next.prev = node
		}
	}

	zero := first
	for zero.number != 0 {
		zero = zero.next
	}

	crt := zero
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			crt = crt.next
		}

		sum += crt.number
	}

	return sum
}

func main() {
	lines, err := utils.ReadFile("20", false)
	if err != nil {
		return
	}

	fmt.Println("Part 1: ", run(parseInput(lines), len(lines), 1))

	first := parseInput(lines)
	for i := 0; i < len(lines); i++ {
		first.number *= 811589153
		first = first.next
	}
	fmt.Println("Part 2: ", run(first, len(lines), 10))
}
