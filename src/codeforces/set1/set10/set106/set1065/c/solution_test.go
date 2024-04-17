package main

import "testing"

func runSample(t *testing.T, k int, H []int, expect int) {
	res := solve(k, H)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{3, 1, 2, 2, 4}
	k := 5
	expect := 2
	runSample(t, k, H, expect)
}

func TestSample2(t *testing.T) {
	H := []int{2, 3, 4, 5}
	k := 5
	expect := 2
	runSample(t, k, H, expect)
}
