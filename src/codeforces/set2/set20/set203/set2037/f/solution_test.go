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
	s := `5 5 3
7 7 7 7 7
1 2 3 4 5`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 10 2
1 1
1 20`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 10 1
69696969 420420420
1 20`
	expect := 6969697
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 10 2
10 15
1 19`
	expect := 15
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2 2 2
1000000000 1
1 3`
	expect := 1000000000
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `10 26 1
13 9 6 3 3 20 7 5 13 9
5 8 9 13 15 20 24 26 30 34`
	expect := 1
	runSample(t, s, expect)
}
