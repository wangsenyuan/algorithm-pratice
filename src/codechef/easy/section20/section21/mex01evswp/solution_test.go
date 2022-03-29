package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "00100"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "01010"
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "10001"
	expect := 2
	runSample(t, s, expect)
}
