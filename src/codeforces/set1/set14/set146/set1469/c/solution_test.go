package main

import "testing"

func runSample(t *testing.T, k int, H []int, expect bool) {
	res := solve(k, H)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{0, 0, 2, 5, 1, 1}
	k := 3
	expect := true
	runSample(t, k, H, expect)
}

func TestSample2(t *testing.T) {
	H := []int{3, 0, 2}
	k := 2
	expect := false
	runSample(t, k, H, expect)
}
