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
	s := `3 5
3 3
4 4
2 5
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 6
3 3
4 4
2 5
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 14
10 5
7 4
11 4
9 8
3 6
6 2
8 9
`
	expect := 32
	runSample(t, s, expect)
}
