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
	s := `5 3
5 1 4
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 8
1 2 3 4 4 3 2 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100000 1
42
`
	expect := 299997
	runSample(t, s, expect)
}
