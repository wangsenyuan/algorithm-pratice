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
	s := `6 2 1
1 1 10 1 1 10`
	expect := 22
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 2 2
1 1 10 1 1 10`
	expect := 12
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 2 3
1 1 10 1 1 10`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 2 10
1 1 10 1 1 10`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6 2 11
1 1 10 1 1 10`
	expect := 2
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6 2 12
1 1 10 1 1 10`
	expect := 0
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `6 2 12
1 1 1 1 10 10`
	expect := -1
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `1 1 1
1`
	expect := 0
	runSample(t, s, expect)
}
