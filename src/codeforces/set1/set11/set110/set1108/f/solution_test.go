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
	s := `8 10
1 2 1
2 3 2
2 4 5
1 4 2
6 3 3
6 1 3
3 5 2
3 7 1
4 8 1
6 2 4
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3
2 1 3
4 3 4
2 4 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
1 2 1
2 3 2
1 3 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 3
1 2 1
2 3 3
1 3 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 6
1 2 2
2 3 1
4 5 3
2 4 2
1 4 2
1 5 3
`
	expect := 2
	runSample(t, s, expect)
}
