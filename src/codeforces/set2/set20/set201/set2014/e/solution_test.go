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
	s := `2 1 1
1
1 2 10`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1 2
2 3
1 2 10`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3 1
2
1 2 4
1 3 10
2 3 6`
	expect := 6
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 3 2
2 3
1 2 10
2 3 18
3 4 16`
	expect := 19
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 2 1
2
1 2 4
1 3 16`
	expect := 14
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `7 7 1
3
1 5 2
2 6 12
1 2 12
6 4 8
7 3 4
6 3 4
7 6 4`
	expect := 12
	runSample(t, s, expect)
}
