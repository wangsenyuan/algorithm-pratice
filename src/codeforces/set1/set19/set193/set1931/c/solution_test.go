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
	s := `7
4 3 2 7 1 1 3`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9
9 9 2 9 2 5 5 5 3`
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 2 3`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9
1 1 1 1 2 1 1 1 1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6
6 6 6 5 6 6`
	expect := 1
	runSample(t, s, expect)
}