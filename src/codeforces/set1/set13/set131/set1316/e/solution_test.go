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
	s := `4 1 2
1 16 10 3
18
19
13
15
`
	expect := 44
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 2 3
78 93 9 17 13 78
80 97
30 52
26 17
56 68
60 36
84 55
`
	expect := 377
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 2 1
500 498 564
100002 3
422332 2
232323 1
`
	expect := 422899
	runSample(t, s, expect)
}
