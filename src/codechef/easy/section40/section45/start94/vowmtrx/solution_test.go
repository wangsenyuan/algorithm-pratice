package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "neo"
	k := 1
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "babylonian"
	k := 2
	expect := 2
	runSample(t, s, k, expect)
}
