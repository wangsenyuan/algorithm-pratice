package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
4 6
10 -2
8 -1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 5
4 -5
4 -2
1 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
5 2
5 -3
2 1
4 -2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 10
10 0
10 -10
30 0
`
	expect := false
	runSample(t, s, expect)
}
