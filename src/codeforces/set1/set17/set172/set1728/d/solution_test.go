package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "forces"
	expect := "Alice"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abba"
	expect := "Draw"
	runSample(t, s, expect)
}
