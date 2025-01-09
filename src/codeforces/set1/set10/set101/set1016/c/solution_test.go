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
	s := `3
1 2 3
6 5 4
`
	expect := 70
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 1000 10000
10 100 100000
`
	expect := 543210
	runSample(t, s, expect)
}
