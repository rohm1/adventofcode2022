package main

import (
	"aoc2022/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type monkey struct {
	ready   bool
	depends []string
	op      string
	number  float64
	expr    string
}

func parseInput(lines []string, part int) map[string]monkey {
	monkeys := map[string]monkey{}
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], ":")
		monkey := monkey{ready: false}

		depends := make([]string, 0)

		expr := strings.TrimSpace(parts[1])
		number, err := strconv.Atoi(expr)
		if err == nil {
			monkey.number = float64(number)
			monkey.expr = expr
			monkey.ready = true
		} else {
			opParts := strings.Split(expr, " ")
			depends = append(depends, opParts[0])
			depends = append(depends, opParts[2])
			monkey.op = opParts[1]
		}

		monkey.depends = depends

		if part == 2 {
			if parts[0] == "root" {
				monkey.op = "="
			} else if parts[0] == "humn" {
				monkey.expr = "x"
			}
		}

		monkeys[parts[0]] = monkey
	}
	return monkeys
}

func op(left float64, right float64, op string) float64 {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}

	panic("operator not handled: " + op)
}

func solve(expr string) float64 {
	exprRegex := regexp.MustCompile("(-?[0-9.]+) ([+\\-*/]) (-?[0-9.]+)")
	intRegex := regexp.MustCompile("\\((-?[0-9.]+)\\)")

	for {
		replace := false

		matched := exprRegex.FindStringSubmatch(expr)
		if len(matched) > 0 {
			l, _ := strconv.ParseFloat(matched[1], 2)
			r, _ := strconv.ParseFloat(matched[3], 2)
			number := op(l, r, matched[2])
			expr = strings.ReplaceAll(expr, matched[0], fmt.Sprintf("%f", number))
			replace = true
		}

		matched = intRegex.FindStringSubmatch(expr)
		if len(matched) > 0 {
			expr = strings.ReplaceAll(expr, matched[0], matched[1])
			replace = true
		}

		if !replace {
			break
		}
	}

	n, _ := strconv.ParseFloat(expr, 2)
	return n
}

func run(lines []string, part int) int {
	monkeys := parseInput(lines, part)

	for {
		for name, monkey := range monkeys {
			if monkey.ready {
				continue
			}

			ready := true
			for j := 0; j < len(monkey.depends); j++ {
				if !monkeys[monkey.depends[j]].ready {
					ready = false
					break
				}
			}

			if !ready {
				continue
			}

			if part == 1 {
				left := monkeys[monkey.depends[0]].number
				right := monkeys[monkey.depends[1]].number

				monkey.number = op(left, right, monkey.op)
				monkey.ready = true
				monkeys[name] = monkey

				if name == "root" {
					return int(monkey.number)
				}
			} else {
				leftExpr := monkeys[monkey.depends[0]].expr
				rightExpr := monkeys[monkey.depends[1]].expr

				l, le := strconv.ParseFloat(leftExpr, 2)
				r, re := strconv.ParseFloat(rightExpr, 2)

				expr := fmt.Sprintf("(%s %s %s)", leftExpr, monkey.op, rightExpr)
				if le == nil && re == nil {
					expr = fmt.Sprintf("%f", op(l, r, monkey.op))
				}

				monkey.expr = expr
				monkey.ready = true
				monkeys[name] = monkey

				if name == "root" {
					start := 0
					end := 1_000_000_000_000_000
					for {
						mid := (end + start) / 2

						l := solve(strings.ReplaceAll(leftExpr, "x", fmt.Sprintf("%d", mid)))
						r := solve(strings.ReplaceAll(rightExpr, "x", fmt.Sprintf("%d", mid)))

						if r > l {
							end = mid
						} else if r < l {
							start = mid
						} else {
							return mid
						}
					}
				}
			}
		}
	}

	panic(fmt.Sprintf("could not solve part %d", part))
}

func main() {
	lines, err := utils.ReadFile("21", false)
	if err != nil {
		return
	}

	fmt.Println("Part 1: ", run(lines, 1))
	fmt.Println("Part 2: ", run(lines, 2))
}
