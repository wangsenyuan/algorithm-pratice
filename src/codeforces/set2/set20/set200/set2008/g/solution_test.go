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
	s := `1 3
3`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 10
1 1`
	expect := 11
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1
1 2 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 2
1 2 4`
	expect := 4
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4 5
2 2 2 16`
	expect := 8
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4 5
2 2 2 3`
	expect := 8
	runSample(t, s, expect)
}