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
	s := `5 0 6
5 4 3 4 9
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2 0
1 2 3 4
`
	expect := 6
	runSample(t, s, expect)
}
