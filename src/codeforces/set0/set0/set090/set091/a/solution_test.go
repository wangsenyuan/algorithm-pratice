package main

import "testing"

func runSample(t *testing.T, s1, s2 string, expect int) {
	res := solve(s1, s2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abc"
	x := "xyz"
	expect := -1
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "abcd"
	x := "dabc"
	expect := 2
	runSample(t, s, x, expect)
}
