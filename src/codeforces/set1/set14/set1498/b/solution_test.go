package main

import "testing"

func runSample(t *testing.T, W int, boxes []int, expect int) {
	res := solve(W, boxes)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	W := 16
	boxes := []int{1, 2, 8, 4, 8}
	expect := 2
	runSample(t, W, boxes, expect)
}

func TestSample2(t *testing.T) {
	W := 10
	boxes := []int{2, 8, 8, 2, 2, 8}
	expect := 3
	runSample(t, W, boxes, expect)
}
