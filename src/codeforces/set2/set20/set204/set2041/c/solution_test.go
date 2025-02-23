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
	s := `3
1 2 3
4 5 6
7 8 9
1 1 1
2 2 2
3 3 3
4 3 0
2 1 4
9 8 9
`
	expect := 5
	runSample(t, s, expect)
}
