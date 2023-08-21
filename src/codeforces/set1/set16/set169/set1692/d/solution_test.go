package main

import "testing"

func runSample(t *testing.T, s string, x int, expect int) {
	res := solve(s, x)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "00:00"
	x := 1
	expect := 16
	runSample(t, s, x, expect)
}
