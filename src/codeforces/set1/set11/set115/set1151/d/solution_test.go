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
	s := `3
4 2
2 3
6 1
`
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
2 4
3 3
7 1
2 3
`
	expect := 25
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
5 10
12 4
31 45
20 55
30 17
29 30
41 32
7 1
5 5
3 15
`
	expect := 1423
	runSample(t, s, expect)
}
