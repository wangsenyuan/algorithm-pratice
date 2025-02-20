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
	s := `1 3 5
0`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2 5
2 5 5`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 10 10
10`
	expect := 20
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 1 10
0 1 2 3 4 5 6 7 8 9`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2 1 2
0 0`
	expect := 2
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2 1 2
0 2`
	expect := 1
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `2 1 3
0 2`
	expect := 2
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `2 2 4
1 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `9 12 54
3 3 8 24 25 27 29 34 53`
	expect := 7
	runSample(t, s, expect)
}
