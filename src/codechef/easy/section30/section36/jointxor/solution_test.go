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
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "11111"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "011"
	expect := 2
	runSample(t, s, expect)
}
