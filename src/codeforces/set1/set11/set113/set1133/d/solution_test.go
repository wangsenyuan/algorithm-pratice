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
	s := `5
1 2 3 4 5
2 4 7 11 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
13 37 39
1 2 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
0 0 0 0
1 2 3 4
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
1 2 -1
-6 -12 6
`
	expect := 3
	runSample(t, s, expect)
}
