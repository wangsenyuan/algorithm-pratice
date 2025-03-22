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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
-1 -2 -3
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
-4 2 0 5 0
`
	expect := 11
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
-1 10 -5 10 -2
`
	expect := 18
	runSample(t, s, expect)
}
