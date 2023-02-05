package main

import "testing"

func runSample(t *testing.T, s []string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"abab",
		"ab",
		"abc",
		"abacb",
		"c",
	}
	expect := "10100"
	runSample(t, s, expect)
}
