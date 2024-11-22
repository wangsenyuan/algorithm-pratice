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
	s := `6
9 12
2 11
1 3
6 10
5 7
4 8`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 3
2 4
5 9
6 8
7 10
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
5 8
3 6
2 9
7 10
1 4
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
1 4
2 5
3 6
7 9
8 10
`
	expect := false
	runSample(t, s, expect)
}

