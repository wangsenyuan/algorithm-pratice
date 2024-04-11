package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "09"
	x := "90"
	expect := 1
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "11"
	x := "55"
	expect := 0
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "?????"
	x := "?????"
	expect := 993531194
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "??"
	x := "??"
	expect := 45 * 45 * 2
	runSample(t, s, x, expect)
}
