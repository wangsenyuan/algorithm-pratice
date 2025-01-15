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
	s := `4 4
1 2 1
2 3 1
3 4 1
`
	expect := 252
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 6
1 2 0
1 3 0
1 4 0
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 5
1 2 1
2 3 0
`
	expect := 210
	runSample(t, s, expect)
}
