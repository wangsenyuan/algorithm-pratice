package main

import "testing"

func runSample(t *testing.T, n int, P []int, have []int, limit []int, expect int64) {
	res := solve(n, P, have, limit)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	P := []int{1, 2}
	have := []int{10, 100}
	limit := []int{10, 1000}
	var expect int64 = 10
	runSample(t, n, P, have, limit, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	P := []int{1, 1, 2, 3, 2}
	have := []int{10, 1, 10, 1, 10}
	limit := []int{10, 2, 10, 2, 10}
	var expect int64 = 12
	runSample(t, n, P, have, limit, expect)
}
