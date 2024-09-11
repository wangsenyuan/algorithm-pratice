package main

import "testing"

func runSample(t *testing.T, s string, x string, expect bool) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababcd"
	x := "abcba"
	expect := true
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "a"
	x := "b"
	expect := false
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "defi"
	x := "fed"
	expect := false
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "xyz"
	x := "x"
	expect := true
	runSample(t, s, x, expect)
}
