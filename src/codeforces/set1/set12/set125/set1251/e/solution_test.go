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
	s := `3
1 5
2 10
2 8`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
0 1
3 1
1 1
6 1
1 1
4 1
4 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
2 6
2 3
2 8
2 7
4 4
5 5`
	expect := 7
	runSample(t, s, expect)
}
