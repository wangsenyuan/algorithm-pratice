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
	s := `2
1 2
2 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
10 10
10 10`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
2 3 3
4 4 1
2 1 1`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
4 2 2 17
1 9 3 1
5 5 5 11
1 2 1 1`
	expect := 3
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1
1`
	expect := 1
	runSample(t, s, expect)
}

