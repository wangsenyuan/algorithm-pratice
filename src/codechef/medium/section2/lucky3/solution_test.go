package main

import "testing"

func runSample(t *testing.T, W []int, expect int64) {
	res := solve(W)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	W := []int{4, 7}
	var expect int64 = 3
	runSample(t, W, expect)
}

func TestSample2(t *testing.T) {
	W := []int{43, 87, 44}
	var expect int64 = 2
	runSample(t, W, expect)
}

func TestSample3(t *testing.T) {
	W := []int{2, 44, 774, 3331, 7542, 45, 132110, 74, 77792, 6}
	var expect int64 = 30
	runSample(t, W, expect)
}
