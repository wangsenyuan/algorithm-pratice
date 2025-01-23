package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 3 4
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
2 7 1 8 2 8
`
	expect := 53
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `12
3 31 314 3141 31415 314159 2 27 271 2718 27182 271828
`
	expect := 592622
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
1 1 1 1
`
	expect := 6
	runSample(t, s, expect)
}
