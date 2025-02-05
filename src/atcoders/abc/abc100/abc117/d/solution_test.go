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
	s := `3 7
1 6 3
`
	expect := 14
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 9
7 4 0 3
`
	expect := 46
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 0
1000000000000
`
	expect := 1000000000000
	runSample(t, s, expect)
}
