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
	s := `4
1 2
1 3
2 4
1 1 1 2
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
2 1
1 1000000000
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
7 3
2 5
2 4
3 1
3 6
2 1
2 7 6 9 3 4 6
`
	expect := 56
	runSample(t, s, expect)
}
