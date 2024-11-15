package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	r := bufio.NewReader(strings.NewReader(s))

	res := process(r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
..
..
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
....
#.#.
....
.#..
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 4
....
.##.
....
`
	expect := 2
	runSample(t, s, expect)
}
