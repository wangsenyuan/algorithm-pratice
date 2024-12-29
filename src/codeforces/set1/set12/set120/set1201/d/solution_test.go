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
	s := `3 3 3 2
1 1
2 1
3 1
2 3
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 5 3 2
1 2
2 3
3 1
1 5
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 6 3 2
1 6
2 2
3 4
1 6
`
	expect := 15
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 11 7 6
1 1
1 5
1 2
1 4
1 11
1 8
1 3
3 6 5 10 2 7
`
	expect := 10
	runSample(t, s, expect)
}