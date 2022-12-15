package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101"
	k := 3
	expect := 6
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "32"
	k := 1
	expect := 8
	runSample(t, s, k, expect)
}