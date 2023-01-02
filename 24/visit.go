package main

type visit struct {
	pos
	t int
	d int
}

func newVisit(p pos, t int) visit {
	return visit{pos: p, t: t}
}

func (v visit) createNext(nextPos pos) visit {
	return newVisit(nextPos, v.t+1)
}
