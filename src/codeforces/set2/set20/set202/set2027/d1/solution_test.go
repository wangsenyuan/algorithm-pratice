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
	s := `4 2
9 3 4 3
11 7`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2
20
19 18`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 2
2 5 2 1 10 3 2 9 9 6
17 9`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 11
2 2 2 2 2 2 2 2 2 2
20 18 16 14 12 10 8 6 4 2 1`
	expect := 10
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1 6
10
32 16 8 4 2 1`
	expect := 4
	runSample(t, s, expect)
}
