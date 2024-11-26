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
	s := `7
2 1 1
3 2 0
4 2 1
5 2 0
6 7 1
7 2 1
`
	expect := 34
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
2 1 1
`
	expect := 2
	runSample(t, s, expect)
}
