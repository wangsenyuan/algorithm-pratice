package main

import "testing"

func runSample(t *testing.T, s string, x int, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Errorf("Sample %s %d, expect %d, but got %d", s, x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "116285317"
	x := 8
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "314159265359"
	x := 1
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "13"
	x := 13
	expect := 0
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "3434343434"
	x := 7
	expect := 5
	runSample(t, s, x, expect)
}
