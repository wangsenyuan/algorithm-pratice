package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "BBWBW"
	k := 3
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "BBWBW"
	k := 5
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "W"
	k := 1
	expect := 1
	runSample(t, s, k, expect)
}
