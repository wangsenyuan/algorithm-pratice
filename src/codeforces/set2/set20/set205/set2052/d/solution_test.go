package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res, cmds, arcs := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	n := len(res)
	pos := make([]int, n)
	var value bool

	for i, x := range res {
		pos[x-1] = i
		s := cmds[x-1]
		if s == "set true" {
			if value == true {
				t.Fatalf("Sample result %v, not correct", res)
			}
			value = true
		} else if s == "set false" {
			if value == false {
				t.Fatalf("Sample result %v, not correct", res)
			}
			// value no change
		} else if s == "unset true" {
			if value == false {
				t.Fatalf("Sample result %v, not correct", res)
			}
			value = false
		} else {
			// unset false
			if value == true {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}

	for _, cur := range arcs {
		u, v := cur[0], cur[1]
		u--
		v--
		if pos[u] > pos[v] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
set true
unset true
set false
unset false
unset false
2
1 4
5 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
unset true
unset false
set true
0
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
unset false
set true
1
2 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
unset false
set false
0
`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2
unset false
set true
0
`
	expect := true
	runSample(t, s, expect)
}
