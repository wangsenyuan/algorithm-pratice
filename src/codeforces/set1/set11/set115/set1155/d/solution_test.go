package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 -2
-3 8 -2 1 -6
`
	expect := 22
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `12 -3
1 3 3 7 1 3 3 7 1 3 3 7
`
	expect := 42
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 10
-1 -2 -3 -4 -5
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 -2
-1 -2 -3 8 -5
`
	expect := 20
	runSample(t, s, expect)
}
