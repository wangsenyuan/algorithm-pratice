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
	s := `2 3
3 3 7 2
3 4 1 5
`
	expect := 15
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 3
4 4 3 1 2
`
	expect := 9
	runSample(t, s, expect)
}
