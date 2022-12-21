package main

import (
	"aoc2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	business  int64
	items     []int64
	operation []string
	test      int64
	testTrue  int
	testFalse int
}

type worryLevelAdjuster func(int64) int64

func parseInput(lines []string) []monkey {
	monkeyCount := len(lines) / 7
	monkeys := make([]monkey, monkeyCount)
	for i := 0; i < monkeyCount; i++ {
		itemsStrArray := strings.Split(lines[i*7+1], " ")
		var items []int64
		for j := 4; j < len(itemsStrArray); j++ {
			worry, _ := strconv.Atoi(strings.ReplaceAll(itemsStrArray[j], ",", ""))
			items = append(items, int64(worry))
		}

		test, _ := strconv.Atoi(strings.Split(lines[i*7+3], " ")[5])
		testTrue, _ := strconv.Atoi(strings.Split(lines[i*7+4], " ")[9])
		testFalse, _ := strconv.Atoi(strings.Split(lines[i*7+5], " ")[9])

		monkeys[i] = monkey{
			business:  0,
			items:     items,
			operation: strings.Split(lines[i*7+2], " "),
			test:      int64(test),
			testTrue:  testTrue,
			testFalse: testFalse,
		}
	}

	return monkeys
}

func op(worry int64, operation []string) int64 {
	left := worry
	if operation[5] != "old" {
		l, _ := strconv.Atoi(operation[5])
		left = int64(l)
	}

	right := worry
	if operation[7] != "old" {
		r, _ := strconv.Atoi(operation[7])
		right = int64(r)
	}

	if operation[6] == "*" {
		worry = left * right
	} else if operation[6] == "+" {
		worry = left + right
	}

	return worry
}

func run(monkeys []monkey, adjust worryLevelAdjuster, turns int) int64 {
	for i := 0; i < turns; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := &monkeys[m]
			items := len(monkey.items)
			monkey.business += int64(items)
			for j := 0; j < items; j++ {
				item := adjust(op(monkey.items[0], monkey.operation))
				monkey.items = monkey.items[1:]

				targetMonkey := monkey.testFalse
				if item%monkey.test == 0 {
					targetMonkey = monkey.testTrue
				}
				monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, item)
			}
		}
	}

	businessLevels := make([]int64, len(monkeys))
	for i := 0; i < len(monkeys); i++ {
		businessLevels[i] = monkeys[i].business
	}

	sort.Slice(businessLevels, func(i, j int) bool { return businessLevels[i] > businessLevels[j] })
	return businessLevels[0] * businessLevels[1]
}

func main() {
	lines, err := utils.ReadFile("11", true)
	if err != nil {
		return
	}

	fmt.Println("Part 1: ", run(parseInput(lines), func(worry int64) int64 { return worry / 3 }, 20))

	monkeys := parseInput(lines)
	mod := int64(1)
	for i := 0; i < len(monkeys); i++ {
		mod *= monkeys[i].test
	}
	fmt.Println("Part 2: ", run(monkeys, func(worry int64) int64 { return worry % mod }, 10_000))
}
