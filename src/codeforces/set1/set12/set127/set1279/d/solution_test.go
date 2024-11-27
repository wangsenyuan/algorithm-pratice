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
	s := `2
2 2 1
1 1
`
	expect := 124780545

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
2 1 2
2 3 1
3 2 4 3
2 1 4
3 4 3 2
`
	expect := 798595483

	runSample(t, s, expect)
}
