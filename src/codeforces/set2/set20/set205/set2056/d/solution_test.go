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
	s := `4
1 1 1 1`
	expect := 10
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 10 2 3 3`
	expect := 11
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
6 3 2 3 5 3 4 2 3 5`
	expect := 42
	runSample(t, s, expect)
}
