package main

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"strconv"
)

type snafu string
type dec int

func (snafu snafu) toDec() dec {
	nbr := dec(0)
	for i := 0; i < len(snafu); i++ {
		n, e := strconv.Atoi(string(snafu[i]))
		if e != nil {
			if snafu[i] == '=' {
				n = -2
			} else if snafu[i] == '-' {
				n = -1
			} else {
				panic("unknown snafu symbol " + string(snafu[i]))
			}
		}

		nbr += dec(n * int(math.Pow(5, float64(len(snafu)-i-1))))
	}

	return nbr
}

func (dec dec) toSnafu() snafu {
	sn := snafu("")
	n := int(dec)
	for p := 0; n != 0; p++ {
		q := n / 5
		r := n - q*5

		symbol := fmt.Sprintf("%d", r)

		if r == 3 {
			n += 5
			symbol = "="
		} else if r == 4 {
			n += 5
			symbol = "-"
		}

		sn = snafu(symbol) + sn
		n = n / 5
	}

	return sn
}

func main() {
	lines, err := utils.ReadFile("25", false)
	if err != nil {
		return
	}

	sum := dec(0)
	for i := 0; i < len(lines); i++ {
		sum += snafu(lines[i]).toDec()
	}

	fmt.Println("Part 1: ", sum.toSnafu())
	//fmt.Println("Part 2: ", part2)
}
