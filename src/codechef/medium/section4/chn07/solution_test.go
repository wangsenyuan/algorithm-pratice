package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "R"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "BB"
	expect := false
	runSample(t, s, expect)
}
