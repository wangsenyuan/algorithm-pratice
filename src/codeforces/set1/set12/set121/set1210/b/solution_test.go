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
		t.Fatalf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
3 2 3 6
2 8 5 10
`
	expect := 15
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 3
1 2 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
3 3 5 5 6 6 1 2 4 7
1 1 1 1 1 1 1 1 1 1
`
	expect := 9
	runSample(t, s, expect)
}
