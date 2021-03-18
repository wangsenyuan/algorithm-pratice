package main

import "testing"

func runSample(t *testing.T, H []int, expect int) {
	res := solve(len(H), H)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{3, 2, 5, 3, 2, 4, 3}
	expect := 3
	runSample(t, H, expect)
}

func TestSample2(t *testing.T) {
	H := []int{1, 2, 4, 3, 5, 4, 2, 1}
	expect := 2
	runSample(t, H, expect)
}
