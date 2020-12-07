package main

import "testing"

func runSample(t *testing.T, n int, F []int, C []int, expect int64) {
	res := solve(n, F, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	F := []int{1, 1, 1}
	C := []int{1, 1, 1}
	var expect int64 = 3
	runSample(t, n, F, C, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	F := []int{3, 0, 0}
	C := []int{10, 0, 0}
	var expect int64 = 30
	runSample(t, n, F, C, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	F := []int{3, 3, 3}
	C := []int{3, 2, 1}
	var expect int64 = 3
	runSample(t, n, F, C, expect)
}
