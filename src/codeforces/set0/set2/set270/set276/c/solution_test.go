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
	s := `3 3
5 3 2
1 2
2 3
1 3
`
	expect := 25
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3
5 2 4 1 3
1 5
2 3
2 3
`
	expect := 33
	runSample(t, s, expect)
}
