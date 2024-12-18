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
	s := `6 10 5
2 5 1
5 3 9
6 2 2
1 3 1
5 1 8
6 5 10
1 6 5
6 4 6
3 6 2
3 4 5
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 15 18
2 6 3
5 7 4
6 5 4
3 6 9
6 7 7
1 6 4
7 1 6
7 2 1
4 3 2
3 2 8
5 3 6
2 5 5
3 7 9
4 1 8
2 1 1
`
	expect := 9
	runSample(t, s, expect)
}
