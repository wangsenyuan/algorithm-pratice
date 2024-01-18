package main

import "testing"

func runSample(t *testing.T, s int, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s, k := 5, 1000000000
	expect := 9999999990
	runSample(t, s, k, expect)
}
