package main

import (
	"testing"
)

var TestCases = map[snafu]dec{
	snafu("2=-1=0"): dec(4890),
	snafu("1=-0-2"): dec(1747),
	snafu("12111"):  dec(906),
	snafu("2=0="):   dec(198),
	snafu("21"):     dec(11),
	snafu("2=01"):   dec(201),
	snafu("111"):    dec(31),
	snafu("20012"):  dec(1257),
	snafu("112"):    dec(32),
	snafu("1=-1="):  dec(353),
	snafu("1-12"):   dec(107),
	snafu("12"):     dec(7),
	snafu("1="):     dec(3),
	snafu("122"):    dec(37),
}

func TestSnafuToInt(t *testing.T) {
	for sn, i := range TestCases {
		if sn.toDec() != i {
			t.Errorf("failed converting %s to %d, got: %d", sn, i, sn.toDec())
		}
	}
}

func TestIntToSnafu(t *testing.T) {
	for sn, i := range TestCases {
		if sn != i.toSnafu() {
			t.Errorf("failed converting %d to %s, got: %s", i, sn, i.toSnafu())
		}
	}
}
