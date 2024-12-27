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
	s := `3
1 1
1 2
1 3
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 1
2 1
3 1
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
2 1
2 2
3 1
3 2
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
1 1
2 1
3 1
1 2
2 2
3 2
`
	expect := 12
	runSample(t, s, expect)
}
