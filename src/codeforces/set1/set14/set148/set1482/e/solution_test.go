package main

import "testing"

func runSample(t *testing.T, H []int, V []int, expect int64) {
	res := solve(H, V)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{1, 2, 3, 5, 4}
	V := []int{1, 5, 3, 2, 4}
	var expect int64 = 15
	runSample(t, H, V, expect)
}

func TestSample2(t *testing.T) {
	H := []int{1, 4, 3, 2, 5}
	V := []int{-3, 4, -10, 2, 7}
	var expect int64 = 10
	runSample(t, H, V, expect)
}

func TestSample3(t *testing.T) {
	H := []int{2, 1}
	V := []int{-2, -3}
	var expect int64 = -3
	runSample(t, H, V, expect)
}

func TestSample4(t *testing.T) {
	H := []int{4, 7, 3, 2, 5, 1, 9, 10, 6, 8}
	V := []int{-4, 40, -46, -8, -16, 4, -10, 41, 12, 3}
	var expect int64 = 96
	runSample(t, H, V, expect)
}
