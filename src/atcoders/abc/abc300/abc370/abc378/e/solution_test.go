package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	expect := bruteForce(bufio.NewReader(strings.NewReader(s)))
	res := process(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
2 5 0
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10 100
320 578 244 604 145 839 156 857 556 400
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 4
1 2 3
`
	runSample(t, s)
}
