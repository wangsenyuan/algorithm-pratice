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
	s := `0001100111
0000011111`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `010101
111000`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `0101
0110`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `0101
1010`
	expect := -1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `011001
001110`
	expect := -1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `0
1`
	expect := -1
	runSample(t, s, expect)
}
