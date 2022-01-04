package main

import "testing"

func runSample(t *testing.T, S1, S2, X string, expect int) {
	res := int(solve(S1, S2, X))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S1 := "ab"
	S2 := "bc"
	X := "abc"
	expect := 5
	runSample(t, S1, S2, X, expect)
}

func TestSample2(t *testing.T) {
	S1 := "aa"
	S2 := "bb"
	X := "ab"
	expect := 3
	runSample(t, S1, S2, X, expect)
}

func TestSample3(t *testing.T) {
	S1 := "aab"
	S2 := "acb"
	X := "bcaabacbc"
	expect := 10
	runSample(t, S1, S2, X, expect)
}
