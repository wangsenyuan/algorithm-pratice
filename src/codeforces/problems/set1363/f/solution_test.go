package main

import "testing"

func runSample(tc *testing.T, s, t string, expect int) {
	res := solve(s, t)
	if res != expect {
		tc.Errorf("Sample %s %s, expect %d, but got %d", s, t, expect, res)
	}
}

func TestSample1(tc *testing.T) {
	s := "abab"
	t := "baba"
	expect := 1
	runSample(tc, s, t, expect)
}