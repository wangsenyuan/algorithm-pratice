package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
3 7 6
9 2 4`
	expect := 42
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3
6 4 1 5 9
8 6 5 1 7`
	expect := 60
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 6
61 95 61 57 69 49 46 47 14 43
39 79 48 92 90 76 30 16 30 94`
	expect := 14579
	runSample(t, s, expect)
}
