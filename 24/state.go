package main

import "fmt"

type state struct {
	s [][]rune
	t int
}

func newState(s [][]rune, t int) state {
	return state{s: s, t: t}
}

func (state state) pos(p pos) rune {
	if p.x < 0 || p.y < 0 || p.x >= len(state.s[0]) || p.y >= len(state.s) {
		return '#'
	}

	return state.s[p.y][p.x]
}

func (state state) print() {
	fmt.Println(fmt.Sprintf("t: %d", state.t))
	for y := 0; y < len(state.s); y++ {
		fmt.Println(string(state.s[y]))
	}
	fmt.Println("==================")
}
