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
	s := `5 1
1 2 3 4 5
1 2
2 3
3 4
3 5
`
	expect := 11
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 2
2 1 2 1 2 1 1
6 4
1 5
3 1
2 3
7 5
7 4
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 3
1 5 2 4 3
5 2
3 2
1 4
2 1
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `14 2
3 2 3 1 2 3 4 2 3 1 4 2 1 3
12 10
2 7
13 12
14 11
10 5
1 2
11 4
4 9
14 8
12 4
7 9
13 6
10 3
`
	expect := 15
	runSample(t, s, expect)
}
