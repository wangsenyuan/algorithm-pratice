package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "to be or not to be"
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "a ab a a b ab a a b c"
	expect := 13
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aa bb aa aa bb bb"
	expect := 11
	runSample(t, s, expect)
}