package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0100"
	x := "0010"
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "01001"
	x := "00011"
	expect := -1
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "000101"
	x := "010011"
	expect := 5
	runSample(t, s, x, expect)
}
