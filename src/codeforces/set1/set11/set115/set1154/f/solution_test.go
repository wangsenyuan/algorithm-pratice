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
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := `7 4 5
2 5 4 2 6 3 1
2 1
6 5
2 1
3 1
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 4 8
6 8 5 1 8 1 1 2 1
9 2
8 4
5 3
9 7
`
	expect := 17
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 1 4
2 5 7 4 6
5 4
`
	expect := 17
	runSample(t, s, expect)
}
