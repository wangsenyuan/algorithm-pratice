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
	s := `10
1 2 3 4 5 6 7 8
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
0 4 1 0 0 9 8 3
`
	expect := 3
	runSample(t, s, expect)
}
