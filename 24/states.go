package main

type states struct {
	states map[int]state
}

func newStates(s map[int]state) states {
	return states{states: s}
}

func (states *states) build(t int) state {
	firstState := states.states[0].s

	maxY := len(firstState)
	maxX := len(firstState[0])

	newStateData := make([][]rune, maxY)
	for y := 0; y < maxY; y++ {
		newStateData[y] = make([]rune, maxX)
		for x := 0; x < maxX; x++ {
			newStateData[y][x] = '.'

			if x == 0 || y == 0 || x == maxX-1 || y == maxY-1 {
				if y == 0 && x == 1 || y == maxY-1 && x == maxX-2 {
					continue
				}

				newStateData[y][x] = '#'
			}
		}
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			c := firstState[y][x]
			if c == '>' {
				t := t % (maxX - 2)
				newStateData[y][(x-1+t)%(maxX-2)+1] = c
			} else if c == '<' {
				t := t % (maxX - 2)
				newStateData[y][(x-1-t+maxX-2)%(maxX-2)+1] = c
			} else if c == 'v' {
				t := t % (maxY - 2)
				newStateData[(y-1+t)%(maxY-2)+1][x] = c
			} else if c == '^' {
				t := t % (maxY - 2)
				newStateData[(y-1-t+maxY-2)%(maxY-2)+1][x] = c
			}
		}
	}

	return newState(newStateData, t)
}

func (states *states) getAt(t int) state {
	if s, f := states.states[t]; f {
		return s
	}

	s := states.build(t)
	//s.print()
	states.states[t] = s
	return s
}
