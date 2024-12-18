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
	s := `7 3 10
2 -4 15 -3 4 8 3
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2 1000
-13 -4 -9 -20 -11
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 1 1
10 -20 3 0 -100
`
	expect := 9
	runSample(t, s, expect)
}
