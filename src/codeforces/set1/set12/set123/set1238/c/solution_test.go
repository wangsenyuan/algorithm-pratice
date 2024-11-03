package main

import "testing"

func runSample(t *testing.T, h int, p []int, expect int) {
	res := solve(h, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := 3
	a := []int{3, 1}
	expect := 0
	runSample(t, h, a, expect)
}

func TestSample2(t *testing.T) {
	h := 8
	a := []int{8, 7, 6, 5, 3, 2}
	// 8 -> 6 -> 4 -> 2 -> 0 (只需要把4放进来)
	//
	expect := 1
	runSample(t, h, a, expect)
}
