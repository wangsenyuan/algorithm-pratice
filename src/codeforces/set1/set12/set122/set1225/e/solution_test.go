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
	s := `1 1
.
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
...R
.RR.
.RR.
R...
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2
.R
R.
`
	expect := 0
	runSample(t, s, expect)
}
