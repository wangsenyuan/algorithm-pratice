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
	s := `5 3
3 1 4
1 5 9
2 6 5
3 5 8
9 7 9
`
	expect := 56
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3
1 -2 3
-4 5 -6
7 -8 -9
-10 11 -12
13 -14 15
`
	expect := 54
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 5
10 -80 21
23 8 38
-94 28 11
-26 -2 18
-69 72 79
-26 -86 -54
-72 -50 59
21 65 -32
40 -94 87
-62 18 82
`
	expect := 638
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 2
2000000000 -9000000000 4000000000
7000000000 -5000000000 3000000000
6000000000 -1000000000 8000000000
`
	expect := 30000000000
	runSample(t, s, expect)
}
