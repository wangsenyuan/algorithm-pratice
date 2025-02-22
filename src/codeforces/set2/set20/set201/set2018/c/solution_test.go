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
	s := `7
1 2
1 3
2 4
2 5
4 6
4 7`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
1 2
1 3
1 4
2 5
3 6
5 7`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `15
12 9
1 6
6 14
9 11
8 7
3 5
13 5
6 10
13 15
13 6
14 12
7 2
8 1
1 4`
	expect := 5
	runSample(t, s, expect)
}
