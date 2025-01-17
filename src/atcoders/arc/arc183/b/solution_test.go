package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 1
1 1 2
1 2 2`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4
2 4 5 1 3
2 1 3 2 2`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `13 1
3 1 3 3 5 3 3 4 2 2 2 5 1
5 3 3 3 4 2 2 2 2 5 5 1 3`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `20 14
10 6 6 19 13 16 15 15 2 10 2 16 9 12 2 6 13 5 5 9
5 9 6 2 10 19 16 15 13 12 10 2 9 6 5 16 19 12 15 13`
	expect := true
	runSample(t, s, expect)
}
