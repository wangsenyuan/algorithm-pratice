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
	s := `6
7 2 3 1 5 6
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
424238336 649760493 681692778 714636916 719885387 804289384 846930887 957747794 596516650 189641422
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
1 2 3 4 5 5 6 7 8 9
`
	expect := 6
	runSample(t, s, expect)
}
