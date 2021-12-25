package main

import "testing"

func runSample(t *testing.T, s1, s2, x string, expect int) {
	res := solve(s1, s2, x)

	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "ab"
	s2 := "bc"
	x := "abc"
	expect := 5
	runSample(t, s1, s2, x, expect)
}

func TestSample2(t *testing.T) {
	s1 := "aa"
	s2 := "bb"
	x := "ab"
	expect := 3
	runSample(t, s1, s2, x, expect)
}

func TestSample3(t *testing.T) {
	s1 := "aab"
	s2 := "acb"
	x := "bcaabacbc"
	expect := 10
	runSample(t, s1, s2, x, expect)
}
