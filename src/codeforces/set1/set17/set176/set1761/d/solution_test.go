package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 15)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 0, 27)
}

func TestSample3(t *testing.T) {
	runSample(t, 998, 244, 573035660)
}
