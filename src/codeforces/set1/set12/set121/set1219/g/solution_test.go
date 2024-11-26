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
	s := `2 2
1 2
3 4
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
0 9 2 7 0
9 0 3 0 5
0 8 0 3 1
6 7 4 3 9
3 6 4 1 0
`
	expect := 80
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 5
1 1 1 1 1
1 0 0 0 1
1 0 0 0 1
1 0 0 0 1
1 1 1 1 1
`
	expect := 16
	runSample(t, s, expect)
}
