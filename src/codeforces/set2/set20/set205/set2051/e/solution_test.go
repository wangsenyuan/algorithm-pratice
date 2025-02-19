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
	s := `2 0
2 1
3 4`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1
2
5`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
1 5 2
3 6 4`
	expect := 9
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 3
2 3 2 8
3 7 3 9`
	expect := 14
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 1
2 9 5
12 14 9`
	expect := 15
	runSample(t, s, expect)
}
