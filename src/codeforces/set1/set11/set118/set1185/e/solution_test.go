package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	ok, res, grid := process(reader)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t %v", expect, ok, res)
	}

	if !expect {
		return
	}

	n := len(grid)
	m := len(grid[0])
	buf := make([][]byte, n)
	for i := range n {
		buf[i] = make([]byte, m)
		for j := range m {
			buf[i][j] = '.'
		}
	}

	draw := func(x byte, pos []int) {
		a, b, c, d := pos[0], pos[1], pos[2], pos[3]
		if a == c {
			for i := b - 1; i < d; i++ {
				buf[a-1][i] = x
			}
		} else {
			for i := a - 1; i < c; i++ {
				buf[i][b-1] = x
			}
		}
	}
	for i := 0; i < len(res); i++ {
		draw(byte('a'+i), res[i])
	}

	for i := range n {
		if string(buf[i]) != grid[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 6
...a..
..bbb.
...a..
.cccc.
...a..`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
...
...
...`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
..c.
adda
bbcb
....`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 5
..b..
aaaaa
..b..`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 3
...
.a.
...`
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2 2
bb
cc`
	expect := true
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `3 2
..
.a
aa`
	expect := false
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `3 2
.a
..
.a`
	expect := false
	runSample(t, s, expect)
}
