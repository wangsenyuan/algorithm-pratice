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
	s := `4 4
S...
#..#
#...
..#T
4
1 1 3
1 3 5
3 2 1
2 3 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
S.
T.
1
1 2 4
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 5
..#..
.S##.
.##T.
.....
3
3 1 5
1 2 3
2 2 1
`
	expect := true
	runSample(t, s, expect)
}
