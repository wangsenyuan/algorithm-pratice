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
	s := `3 5 1 3
0 0 1
2 1 4`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 5 2 3
1 0
3 2`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 20 2 4
0
16`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 20 2 5
1 1 0 1 0 0
0 8 2 9 11 6`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4 16 3 6
1 0 1 1
8 3 5 6`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6 20 3 6
0 1 0 0 1 0
20 11 3 20 16 17`
	expect := 4
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `7 17 1 6
1 1 0 1 0 0 0
1 7 0 11 10 15 10`
	expect := 0
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `6 17 2 6
0 0 1 0 0 1
7 6 3 7 10 12`
	expect := 1
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `5 17 2 5
1 1 1 1 0
17 11 10 6 4`
	expect := 2
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := `1 1 1 2
0
1`
	expect := 1
	runSample(t, s, expect)
}
