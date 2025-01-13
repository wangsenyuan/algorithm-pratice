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
	s := `6
1 0 1 1 0 1
1 2 2 2 2
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 0 1 0 1
1 1 1 1
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8
1 0 0 1 0 1 1 0
1 1 2 2 3 3 3
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9
1 1 0 0 1 0 1 0 1
1 1 2 2 3 3 4 4
`
	expect := 5
	runSample(t, s, expect)
}
