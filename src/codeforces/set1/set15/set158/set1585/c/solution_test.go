package main

import "testing"

func runSample(t *testing.T, X []int, k int, expect int64) {
	res := solve(X, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	X := []int{1, 2, 3, 4, 5}
	var expect int64 = 25
	runSample(t, X, k, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	X := []int{-5, -10, -15, 6, 5, 8, 3, 7, 4}
	var expect int64 = 41
	runSample(t, X, k, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	X := []int{2, 2, 3, 3, 3}
	var expect int64 = 7
	runSample(t, X, k, expect)
}

func TestSample4(t *testing.T) {
	k := 2
	X := []int{1000000000, 1000000000, 1000000000, 1000000000}
	var expect int64 = 3000000000
	runSample(t, X, k, expect)
}
