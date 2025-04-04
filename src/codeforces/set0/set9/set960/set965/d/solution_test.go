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
	s := `10 5
0 0 1 0 2 0 0 1 0
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 3
1 1 1 1 2 1 1 1 1
`
	expect := 3
	runSample(t, s, expect)
}
