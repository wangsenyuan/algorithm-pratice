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
	s := `3 10
6 3
8 2
1 4`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 10
4 1
3 2
2 6
1 100`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 15
10 11
14 100`
	expect := -1
	runSample(t, s, expect)
}
