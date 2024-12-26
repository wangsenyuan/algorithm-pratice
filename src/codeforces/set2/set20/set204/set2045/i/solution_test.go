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
	s := `5 4
3 2 1 3 2
`
	expect := 13
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 1 1
`
	expect := 2
	runSample(t, s, expect)
}
