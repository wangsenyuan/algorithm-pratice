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
	s := `100 0 24`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 10
1 2 1`
	expect := 55
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1 2
2 1 1`
	expect := 777777784
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 2 4
1 2 25
3 2 24`
	expect := 40000020
	runSample(t, s, expect)
}
