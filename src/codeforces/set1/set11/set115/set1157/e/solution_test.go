package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
0 1 2 1
3 2 1 1
`
	expect := "1 0 0 2"

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
2 5 1 5 3 4 3
2 4 3 5 6 5 1
`
	expect := "0 0 0 1 0 2 4"

	runSample(t, s, expect)
}
