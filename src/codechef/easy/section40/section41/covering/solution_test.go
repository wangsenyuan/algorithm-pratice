package main

import "testing"

func runSample(t *testing.T, n int, F, G, H []int, expect int) {
	res := solve(n, F, G, H)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	F := []int{1, 3, 9, 12}
	G := []int{0, 5, 1, 2}
	H := []int{2, 3, 4, 1}
	expect := 7680
	runSample(t, n, F, G, H, expect)
}
