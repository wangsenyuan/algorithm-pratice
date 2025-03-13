package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res, roads := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	for _, road := range roads {
		u, v, l, m := road[0], road[1], road[2], road[3]
		u--
		v--
		if res[u]*res[v] != l*m {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `1 0`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1
1 2 1 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 2
3 2 1 2
3 1 1 10
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 1
1 2 3 7
`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 2
1 2 8 8
1 3 1 48
`
	expect := false
	runSample(t, s, expect)
}
