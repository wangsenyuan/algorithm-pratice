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
	s := `2 1
1 2
1 2
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
3 1 2
1 2
3 1
3 2
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 2
3 1 5 4 2
5 2
5 4
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 11
5 1 3 4 2
5 1
5 2
1 5
2 1
1 2
1 4
2 5
1 3
5 4
5 3
3 1
`
	expect := 2
	runSample(t, s, expect)
}
