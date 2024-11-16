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
	s := `4
0 0
1 1
0 3
1 2
`
	expect := 14
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
0 0
0 2
0 4
2 0
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
-1 -1
1 0
3 1
`
	expect := 0
	runSample(t, s, expect)
}
