package main

import "testing"

func runSample(t *testing.T, bonus []int, k int, expect int64) {
	res := solve(k, bonus)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 0
	bonus := []int{1, 1, 1}
	var expect int64 = 3
	runSample(t, bonus, k, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	bonus := []int{-1, -2, -3, -4, 5}
	var expect int64 = 11
	runSample(t, bonus, k, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	bonus := []int{3, 1, 4, 1, 5, -9, -2, -6, -5, -3, -5, -8, -9}
	var expect int64 = 71
	runSample(t, bonus, k, expect)
}

func TestSample4(t *testing.T) {
	k := 1
	bonus := []int{-1, -2, -3, -4, 5}
	var expect int64 = 11
	runSample(t, bonus, k, expect)
}
