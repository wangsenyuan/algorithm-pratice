package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abab"
	x := "ba"
	expect := 12
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "cacdcdbbbb"
	x := "bdcaccdbbb"
	expect := 24
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "rotator"
	x := "rotator"
	expect := 4
	runSample(t, s, x, expect)
}
