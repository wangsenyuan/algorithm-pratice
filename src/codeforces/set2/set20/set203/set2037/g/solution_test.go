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
	s := `5
2 6 3 4 6
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
4 196 2662 2197 121
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
3 6 8 9 11 12 20
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
2 3
`
	expect := 0
	runSample(t, s, expect)
}
