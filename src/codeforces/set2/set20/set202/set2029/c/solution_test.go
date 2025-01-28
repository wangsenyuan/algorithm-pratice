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
	s := `6
1 2 3 4 5 6`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
1 2 1 1 1 3 4`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9
9 9 8 2 4 4 3 5 3`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
1 2 3 4 1 3 2 1 1 10`
	expect := 5
	runSample(t, s, expect)
}
