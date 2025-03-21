package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
3 2 2 4
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
4 5 3 6 7
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
10 1 2 3 4
`
	expect := 5
	runSample(t, s, expect)
}
