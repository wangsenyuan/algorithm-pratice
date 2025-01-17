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
	s := `3 2
1 3 2
1 2 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1
1 1 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 5
3 8 4
3 10 4
1 7 2
1 8 3
3 8 7
`
	expect := 1598400
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `15 17
2 11 9
2 15 13
1 14 2
5 11 5
3 15 11
1 6 2
4 15 12
3 11 6
9 13 10
2 14 6
10 15 11
1 8 6
6 14 8
2 10 2
6 12 6
3 14 12
2 6 2
`
	expect := 921467228
	runSample(t, s, expect)
}
