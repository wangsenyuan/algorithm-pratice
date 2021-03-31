package main

import "testing"

func runSample(t *testing.T, W []int, expect int64) {
	res := solve(len(W), W)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	W := []int{1, 1, 1}
	var expect int64 = 3
	runSample(t, W, expect)
}

func TestSample2(t *testing.T) {
	W := []int{5, 3, 2}
	var expect int64 = 7
	runSample(t, W, expect)
}

func TestSample3(t *testing.T) {
	W := []int{4, 2, 5}
	var expect int64 = 8
	runSample(t, W, expect)
}
