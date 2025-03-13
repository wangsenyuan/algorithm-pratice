package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if expect != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
10 10 10`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
4 4 7 7`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 4
1 2 3 4 5 6 7`
	expect := 35
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `20 7
1 4 5 8 47 777777777 1 5 4 8 5 9 5 4 7 4 5 7 7 44474`
	expect := 29172
	runSample(t, s, expect)
}
