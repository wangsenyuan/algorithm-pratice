package main

import "testing"

func runSample(t *testing.T, dest []int, C []int, expect int64) {
	res := solve(dest, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	dest := []int{-3, 1}
	C := []int{1, 3, 5, 7, 9, 11}
	var expect int64 = 18
	runSample(t, dest, C, expect)
}

func TestSample2(t *testing.T) {
	dest := []int{1000000000, 1000000000}
	C := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	var expect int64 = 1000000000000000000
	runSample(t, dest, C, expect)
}
