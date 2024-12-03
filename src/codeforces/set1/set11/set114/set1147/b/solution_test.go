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
	s := `12 6
1 3
3 7
5 7
7 11
9 11
11 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 3
1 2
3 2
7 2
`
	expect := false
	runSample(t, s, expect)
}
func TestSample3(t *testing.T) {
	s := `10 2
1 6
2 7
`
	expect := true
	runSample(t, s, expect)
}
