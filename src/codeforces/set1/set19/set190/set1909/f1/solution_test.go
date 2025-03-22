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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2 3 4 5`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
0 2 2 2 4 6`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `15
0 0 1 1 1 2 3 4 5 6 7 9 11 13 15`
	expect := 532305727
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
0 1 3 4 5 5`
	expect := 0
	runSample(t, s, expect)
}