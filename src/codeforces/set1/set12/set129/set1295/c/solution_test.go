package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aabce"
	x := "ace"
	expect := 1
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "abacaba"
	x := "aax"
	expect := -1
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "ty"
	x := "yyt"
	expect := 3
	runSample(t, s, x, expect)
}
