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
	s := `6
6 3 3 3 5 5`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
5 6 4 1 4 5`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9
8 6 4 2 1 3 5 7 9`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
4 3 3 2`
	expect := 2
	runSample(t, s, expect)
}
