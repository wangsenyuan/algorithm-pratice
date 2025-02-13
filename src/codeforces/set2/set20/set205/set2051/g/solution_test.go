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
	s := `3 6
1 +
1 -
3 +
3 -
2 +
2 -
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 13
5 +
3 +
5 -
2 +
4 +
3 +
5 +
5 -
2 +
3 -
3 +
3 -
2 +
`
	expect := 11
	runSample(t, s, expect)
}
