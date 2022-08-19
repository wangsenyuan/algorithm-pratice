package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1010"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "101"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "100100"
	expect := 1
	runSample(t, s, expect)
}
