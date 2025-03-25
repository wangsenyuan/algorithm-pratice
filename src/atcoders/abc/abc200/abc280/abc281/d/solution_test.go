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
	s := `4 2 2
1 2 3 4
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1 2
1 3 5
`
	expect := -1
	runSample(t, s, expect)
}