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
	s := `3 4
3 1 6 -2
2 3 3
2 -5 1
2 3 1 3
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 1
4 0 8 -3 -10
8 3 -2 -5 10 8 -9 -5 -4
1 0
1 -3
3 -8 5 6
2 9 6
1
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1
2 -4 -6
5 6 8 3 5 -2
1
`
	expect := -4
	runSample(t, s, expect)
}
