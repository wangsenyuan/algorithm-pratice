package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2
2 3
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 2
1 3
1 4
2 5
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
1 2
1 3
1 4
2 5
2 6
`
	expect := true
	runSample(t, s, expect)
}
