package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "BGGGG"
	k := 2
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "RBRGR"
	k := 3
	expect := 0
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "BBBRR"
	k := 5
	expect := 3
	runSample(t, s, k, expect)
}
